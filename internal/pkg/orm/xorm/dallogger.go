package xorm

import (
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	xorm "xorm.io/xorm/log"
)

// logger is a logger interface
type Logger struct {
	level   xorm.LogLevel
	showSQL bool
}

//debug
func (l *Logger) Debug(v ...interface{}) {
	log.Infoln(v...)
}

//debugf
func (l *Logger) Debugf(format string, v ...interface{}) {
	log.Infof(format, v...)
}

//error
func (l *Logger) Error(v ...interface{}) {
	log.Errorln(v...)
}

//errorf
func (l *Logger) Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}

//info
func (l *Logger) Info(v ...interface{}) {
	log.Infoln(v...)
}

//infof
func (l *Logger) Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}

//warn
func (l *Logger) Warn(v ...interface{}) {
	log.Warnln(v...)
}

//warnf
func (l *Logger) Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

//level
func (l *Logger) Level() xorm.LogLevel {
	return l.level
}

//set level
func (l *Logger) SetLevel(level xorm.LogLevel) {
	l.level = level
}

//show sql
func (l *Logger) ShowSQL(show ...bool) {
	l.showSQL = true
	for _, s := range show {
		if s == false {
			l.showSQL = false
		}
	}
}

//is show sql
func (l *Logger) IsShowSQL() bool {
	return l.showSQL
}
