package main

import (
	"net/http"

	"github.com/akwanmaroso/monitoring/distributed/web/controller"
)

func main() {
	controller.Initialize()

	http.ListenAndServe(":3000", nil)
}
