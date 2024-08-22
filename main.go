package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ShebnSp/endecrypror/internal/routes"
)

var webPort = 8080

func main() {
	
	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", webPort),
		Handler: routes.RegisterRoutes(),
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	
	log.Println("Server started on port :",8080)
	log.Fatal(srv.ListenAndServe())
}
