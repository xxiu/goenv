package goenv

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

var (
	EnvMap map[string]string
)

func init() {
	EnvMap = make(map[string]string)
}

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

func Getenv(value string) string {
	return EnvMap[value]
}

func String(value string) string {
	return Getenv(value)
}

func Int(value string) int {
	v := Getenv(value)
	if s, err := strconv.Atoi(v); err == nil {
		return s
	}
	return 0
}

func Boolean(value string) bool {
	v := Getenv(value)
	if s, err := strconv.ParseBool(v); err == nil {
		return s
	}
	return false
}

func Float(value string) float64 {
	v := Getenv(value)
	if s, err := strconv.ParseFloat(v, 64); err == nil {
		return s
	}
	return 0
}
