package autoload

import (
	env "github.com/xxiu/goenv"
)

const (
	LOCAL_CONFIG_FILE = ".env"
)

func init() {
	env.Load(LOCAL_CONFIG_FILE)
}
