package clientapp

import (
	"fmt"

	cli "github.com/limpid-kzonix/golang-sandbox/modules/cli-app"
)

func Cli() {
	fmt.Println("CLI")
	cli.Server()
}
