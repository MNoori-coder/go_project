package main

import (
	"fmt"
	ticket "go_project/ticket"
	ticket_category "go_project/ticket_category"
	ticket_comment "go_project/ticket_comment"
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
	ticket_comment.TicketCommentRouters(mux)
	ticket_category.TicketCategoryRouters(mux)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
}
