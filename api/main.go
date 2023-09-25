package main

import (
	"DevBook/api/src/config"
	"DevBook/api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Carregar()
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
