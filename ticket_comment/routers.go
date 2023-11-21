package ticket_comment

import "net/http"

func TicketCommentRouters(mux *http.ServeMux) {

	// GET
	mux.HandleFunc("/ticket_comment/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			ListTicketComments(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// POST
	mux.HandleFunc("/ticket_comment/insert/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			CreateTicketComment(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// PUT
	mux.HandleFunc("/ticket_comment/update/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPut {
			UpdateTicketComment(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

}
