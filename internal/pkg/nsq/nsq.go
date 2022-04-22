package nsq

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"github.com/nsqio/go-nsq"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type (
	NsqCfg struct {
		//nsq nsqlookupd address
		Nsqlookupd string `ini:"Nsqlookupd"`
	}

	Nodes struct {
		Producers []Producer `json:"producers"`
	}

	Producer struct {
		RemoteAddress    string        `json:"remote_address"`
		Hostname         string        `json:"hostname"`
		BroadcastAddress string        `json:"broadcast_address"`
		TCPPort          int           `json:"tcp_port"`
		HTTPPort         int           `json:"http_port"`
		Version          string        `json:"version"`
		Tombstones       []interface{} `json:"tombstones"`
		Topics           []interface{} `json:"topics"`
	}

	Nsq interface {
		//publish message
		Publish(topic string, msg []byte) error

		//publish deferred message
		DeferredPublish(topic string, delay time.Duration, msg []byte) error

		//nsqd consumer
		Consumer(topic, channel string, f func(msg []byte) error)

		//clsoe
		Close()
	}

	nsqImpl struct {
		cfg      *NsqCfg
		nodes    *Nodes
		producer *nsq.Producer
		consumer []*nsq.Consumer
		lock     sync.Mutex
	}
)

//NewNsq 新建一个Nsq
func NewNsq(config *NsqCfg) Nsq {
	url, err := buildLookupAddr(config.Nsqlookupd)
	if err != nil {
		panic(err)
	}
	nodes, err := httpGet(url)
	if err != nil {
		panic(err)
	}

	nsq := &nsqImpl{
		nodes: nodes,
		cfg:   config,
	}
	if err := nsq.connectNsqd(nodes); err != nil {
		panic(err)
	}
	return nsq
}

//publish message
func (n *nsqImpl) Publish(topic string, msg []byte) error {
	topic = strings.Replace(topic, "/", "_", -1)
	if err := n.producer.Ping(); err != nil {
		n.producer.Stop()
		if err := n.connectNsqd(n.nodes); err == nil {
			return err
		}
	}
	return n.producer.Publish(topic, msg)
}

//publish deferred message
func (n *nsqImpl) DeferredPublish(topic string, delay time.Duration, msg []byte) error {
	topic = strings.Replace(topic, "/", "_", -1)
	if err := n.producer.Ping(); err != nil {
		n.producer.Stop()
		if err := n.connectNsqd(n.nodes); err == nil {
			return err
		}
	}
	return n.producer.DeferredPublish(topic, delay, msg)
}

//nsqd consumer
func (n *nsqImpl) Consumer(topic, channel string, f func(msg []byte) error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = time.Second
	config.MaxInFlight = len(n.nodes.Producers)
	topic = strings.Replace(topic, "/", "_", -1)
	channel = strings.Replace(channel, "/", "_", -1)
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		panic(err)
	}
	consumer.SetLogger(new(Logger), nsq.LogLevelInfo)
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		if err := f(message.Body); err != nil {
			if !message.IsAutoResponseDisabled() {
				message.RequeueWithoutBackoff(time.Second * 5)
			}
		} else {
			message.Finish()
		}
		return nil
	}))
	if err := consumer.ConnectToNSQLookupd(n.cfg.Nsqlookupd); err != nil {
		panic(err)
	}

	n.consumer = append(n.consumer, consumer)
	<-consumer.StopChan

	//end
	stats := consumer.Stats()
	log.Tracef("message received %d, finished %d, requeued:%s, connections:%s", stats.MessagesReceived, stats.MessagesFinished, stats.MessagesRequeued, stats.Connections)
}

//clsoe
func (n *nsqImpl) Close() {
	for _, consumer := range n.consumer {
		consumer.Stop()
	}
	n.producer.Stop()
}

func (n *nsqImpl) connectNsqd(nodes *Nodes) error {
	n.lock.Lock()
	defer n.lock.Unlock()

	if len(nodes.Producers) <= 0 {
		return errors.New("not find nsq node")
	}
	node := nodes.Producers[rand.Intn(len(nodes.Producers))]
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(fmt.Sprintf("%s:%d", node.BroadcastAddress, node.TCPPort), config)
	if err != nil {
		return err
	}
	producer.SetLogger(new(Logger), nsq.LogLevelInfo)
	err = producer.Ping()
	if err != nil {
		producer.Stop()
		return err
	}

	n.producer = producer
	return nil
}

func buildLookupAddr(addr string) (string, error) {
	urlString := addr
	if !strings.Contains(urlString, "://") {
		urlString = "http://" + addr
	}

	u, err := url.Parse(urlString)
	if err != nil {
		return "", err
	}

	if u.Port() == "" {
		return "", errors.New("missing port")
	}

	if u.Path == "/" || u.Path == "" {
		u.Path = "/nodes"
	}
	return u.String(), nil
}

func httpGet(url string) (*Nodes, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		nodes := new(Nodes)
		if err := json.Unmarshal(body, nodes); err != nil {
			return nil, err
		}
		return nodes, nil
	}
	return nil, fmt.Errorf("http get status code %d != 200 ", resp.StatusCode)
}

//Logger 日志
type Logger struct {
}

//Output 输出
func (l *Logger) Output(calldepth int, s string) error {
	//log.Traceln(s)
	return nil
}
