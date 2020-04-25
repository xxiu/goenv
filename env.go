package goenv

import (
	log "github.com/sirupsen/logrus"
)

var (
	EnvMap map[string]string
)

func Load(envfile ...string) {
	for _, fi := range envfile {
		loadFile(fi)
	}
}

func loadFile(file string) {
	m, err := Read(file)
	if err != nil {
		log.WithFields(log.Fields{
			"configfile": file,
			"err":        err,
		}).Error("loadFile")
		return
	}
	for k, v := range m {
		EnvMap[k] = v
	}
}

func Getenv(s string) string {
	return EnvMap[s]
}
