package main

import (
	"fmt"
	ticket "go_project/ticket"
	"go_project/utils"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fmt.Println("Start server")

	// Database
	utils.Connection()

	// Routers
	ticket.TicketRouters(mux)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
