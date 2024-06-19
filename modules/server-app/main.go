package serverapp

import (
	"fmt"

	client "github.com/limpid-kzonix/golang-sandbox/modules/client-app"
)

func Server() {
	fmt.Println("Server")
	client.Client()
}
