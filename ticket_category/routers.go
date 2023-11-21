package ticket_category

import (
	"net/http"
	_ "net/http"
)

func TicketCategoryRouters(mux *http.ServeMux) {

	// GET
	mux.HandleFunc("/ticket_category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListTicketCategory(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// POST
	mux.HandleFunc("/ticket_category/insert/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateTicketCategory(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
