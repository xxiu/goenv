package main

import (
	"fmt"

	env "github.com/xxiu/goenv"
	_ "github.com/xxiu/goenv/autoload"
)

func main() {

	db := env.Getenv("DB")
	fmt.Printf("db:%s \n", db)

}
