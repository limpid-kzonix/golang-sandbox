package clientapp

import (
	"fmt"

	server "github.com/limpid-kzonix/golang-sandbox/modules/server-app"
)

func Client() {
	fmt.Println("Client")
	server.Server()
}
