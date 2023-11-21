package ticket

import (
	"net/http"
	_ "net/http"
)

func TicketRouters(mux *http.ServeMux) {

	// GET
	mux.HandleFunc("/ticket/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListTicket(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// POST
	mux.HandleFunc("/ticket/insert/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateTicket(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// PUT
	mux.HandleFunc("/ticket/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			UpdateTicket(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// DELETE
	mux.HandleFunc("/ticket/delete/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			DeleteTicket(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// GET
	mux.HandleFunc("/ticket/id/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			RetrieveTicket(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
