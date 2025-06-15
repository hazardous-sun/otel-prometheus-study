package main

import (
	"log"
	http2 "otel-prometheus-study/internal/handler/http"
)

func main() {
	// Initialize the controllers
	customerCtrl := http2.NewCustomerController()
	productCtrl := http2.NewProductController()
	stockCtrl := http2.NewStockController()
	storeCtrl := http2.NewStoreController()
	storeProductCtrl := http2.NewStoreProductController()

	// Initialize the server with the controllers
	server := http2.NewRouter(
		customerCtrl,
		productCtrl,
		stockCtrl,
		storeCtrl,
		storeProductCtrl,
	)

	if err := server.Run(":8000"); err != nil {
		log.Fatal(err)
		return
	}
}
