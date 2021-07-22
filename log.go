package main

import (
	zc "log"
	"go.uber.org/zap"
)

var log *zap.SugaredLogger

func init() {
	LogOn()
}

func LogOn() {
	log = zc.GetDevelopmentLogger("", "").Sugar()
}

func LogOff() {
	log = zc.GetNoopLogger().Sugar()
}

func LogSet(l *zap.SugaredLogger) {
	log = l
}
