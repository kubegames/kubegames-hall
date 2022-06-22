package game

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	modelgame "github.com/kubegames/kubegames-hall/app/model/game"
	gameservice "github.com/kubegames/kubegames-hall/app/service/game"
	"github.com/kubegames/kubegames-hall/app/service/game/types"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	mysql "github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	gamev1 "github.com/kubegames/kubegames-operator/pkg/apis/game/v1"
	gamesclientset "github.com/kubegames/kubegames-operator/pkg/client/game/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"xorm.io/xorm"
)

const (
	Namespace = "games"
)

type gameImpl struct {
	mysql          *mysql.Mysql
	kubeclientset  kubernetes.Interface
	gamesclientset gamesclientset.Interface
}

//new game
func NewGame(
	mysql *mysql.Mysql,
	kubeclientset kubernetes.Interface,
	gamesclientset gamesclientset.Interface) *gameImpl {
	//sync table
	if err := mysql.GetWriteEngine().Sync2(
		new(modelgame.Game),
		new(modelgame.GameConfig),
	); err != nil {
		panic(err.Error())
	}

	return &gameImpl{
		mysql:          mysql,
		kubeclientset:  kubeclientset,
		gamesclientset: gamesclientset,
	}
}

func (impl *gameImpl) AddGameRequest(ctx context.Context, request *types.AddGameRequest) (response *types.AddGameResponse, err error) {
	if len(request.Config.Version) <= 0 {
		request.Config.Version = "v1"
	}
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//game insert
	if _, err := modelgame.InsertGame(session, modelgame.Game{
		GameID: request.GameID,
		Name:   request.Name,
	}); err != nil {
		log.Errorf("game insert error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//game config insert
	if _, err := modelgame.InsertGameConfig(session, modelgame.GameConfig{
		Name:      request.Name,
		GameID:    request.GameID,
		Platform:  request.Config.Platform,
		MaxPeople: request.Config.MaxPeople,
		Runmode:   request.Config.Runmode,
	}); err != nil {
		log.Errorf("game config insert error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//commit
	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit insert error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	return &types.AddGameResponse{Success: true}, nil
}

func (impl *gameImpl) StartGame(ctx context.Context, request *types.StartGameRequest) (response *types.StartGameResponse, err error) {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//select game
	_, ok, err := modelgame.FindGame(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("game_id = ?", request.GameID)
	})
	if err != nil {
		log.Errorf("mysql get game error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		return nil, gameservice.NewGameServiceError(http.StatusNotFound, "game not found")
	}

	//select game config
	config, ok, err := modelgame.FindGameConfig(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("game_id = ?", request.GameID)
	})
	if err != nil {
		log.Errorf("mysql get game error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		return nil, gameservice.NewGameServiceError(http.StatusNotFound, "game config not found")
	}

	//update game
	updates := map[string]interface{}{
		"status":   true,
		"image":    request.Image,
		"replicas": request.Replicas,
		"port":     request.Port,
		"cpu":      request.Cpu,
		"memory":   request.Memory,
	}
	if len(request.Commonds) > 0 {
		value, _ := json.Marshal(request.Commonds)
		updates["commonds"] = value
	}
	if _, err := modelgame.UpdateGame(session, updates, func(s *xorm.Session) *xorm.Session {
		return s.Where("game_id = ?", request.GameID)
	}); err != nil {
		log.Errorf("update games error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//deploy game
	deployGame := gamev1.Game{
		Spec: gamev1.GameSpec{
			GameID:   fmt.Sprintf("%d", request.GameID),
			Image:    request.Image,
			Commonds: request.Commonds,
			Port:     request.Port,
			Replicas: request.Replicas,
			Cpu:      request.Cpu,
			Memory:   request.Memory,
		},
	}

	//set name
	deployGame.Name = fmt.Sprintf("%d", request.GameID)
	deployGame.Namespace = Namespace

	//set config
	cfg, err := json.Marshal(config)
	if err != nil {
		log.Errorf("json marshal err %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusBadRequest, err.Error())
	}
	deployGame.Spec.Config = string(cfg)

	//chek game namespace
	if _, err := impl.kubeclientset.CoreV1().Namespaces().Get(ctx, Namespace, v1.GetOptions{}); err != nil {
		if errors.IsNotFound(err) == false {
			log.Errorf("get namespace error %s", err.Error())
			return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
		}

		//create namespace
		ns := &corev1.Namespace{}
		ns.Name = Namespace
		if _, err := impl.kubeclientset.CoreV1().Namespaces().Create(ctx, ns, v1.CreateOptions{}); err != nil {
			log.Errorf("create namespace error %s", err.Error())
			return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
		}
	}

	if _, err := impl.gamesclientset.KubegamesV1().Games(Namespace).Create(ctx, &deployGame, v1.CreateOptions{}); err != nil {
		log.Errorf("deploy games err %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//commit
	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit insert error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	return &types.StartGameResponse{Success: true}, nil
}

func (impl *gameImpl) CloseGame(ctx context.Context, request *types.CloseGameRequest) (response *types.CloseGameResponse, err error) {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//select game
	_, ok, err := modelgame.FindGame(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("game_id = ?", request.GameID)
	})
	if err != nil {
		log.Errorf("mysql get game error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}
	if ok {
		//update game
		if _, err := modelgame.UpdateGame(session, map[string]interface{}{
			"status": false,
		}, func(s *xorm.Session) *xorm.Session {
			return s.Where("game_id = ?", request.GameID)
		}); err != nil {
			log.Errorf("update games error %s", err.Error())
			return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
		}
	}

	//found games from k8s
	if err := impl.gamesclientset.KubegamesV1().Games(Namespace).Delete(ctx, fmt.Sprintf("%d", request.GameID), v1.DeleteOptions{}); err != nil {
		log.Errorf("delete games from k8s error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//commit
	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit insert error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.CloseGameResponse{Success: true}, nil
}

func (impl *gameImpl) FindGame(ctx context.Context, request *types.FindGameRequest) (response *types.FindGameResponse, err error) {
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	gameinfo, ok, err := modelgame.FindGame(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("game_id = ?", request.GameID)
	})
	if err != nil {
		log.Errorf("find gams error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("find game %s not found", request.GameID)
		return nil, gameservice.NewGameServiceError(http.StatusNotFound, fmt.Sprintf("find game %d not found", request.GameID))
	}

	//get config
	config, ok, err := modelgame.FindGameConfig(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("game_id = ?", gameinfo.GameID)
	})
	if err != nil {
		log.Errorf("find gams error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}
	if ok {
		gameinfo.Config = &config
	}

	return &types.FindGameResponse{Game: &gameinfo}, nil
}

func (impl *gameImpl) FindGames(ctx context.Context, request *types.FindGamesRequest) (response *types.FindGamesResponse, err error) {
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	games, err := modelgame.FindGames(session)
	if err != nil {
		log.Errorf("find gams error %s", err.Error())
		return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
	}

	//range
	for _, game := range games {
		config, ok, err := modelgame.FindGameConfig(session, func(s *xorm.Session) *xorm.Session {
			return s.Where("game_id = ?", game.GameID)
		})
		if err != nil {
			log.Errorf("find games error %s", err.Error())
			return nil, gameservice.NewGameServiceError(http.StatusInternalServerError, err.Error())
		}
		if ok {
			game.Config = &config
		}
	}
	return &types.FindGamesResponse{Games: games}, nil
}
