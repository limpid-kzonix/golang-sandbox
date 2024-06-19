package clientapp

import (
	"fmt"

	server "github.com/limpid-kzonix/golang-sandbox/modules/server-app"
)

func Cli() {
	fmt.Println("CLI")
	server.Server()
}
