package ticket

import (
	"database/sql"
	"encoding/json"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
	"strconv"
	"strings"
)

func ListTicket(w http.ResponseWriter, r *http.Request) {
	var tickets []SchemaTicket

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ticket")
	if err != nil {
		http.Error(w, "Can not execute query\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ticket SchemaTicket
		err := rows.Scan(&ticket.ID, &ticket.UserID, &ticket.Title, &ticket.Description, &ticket.CategoryID, &ticket.Status, &ticket.CreatedAt, &ticket.UpdatedAt)
		if err != nil {
			http.Error(w, "Can not scan row data\n"+err.Error(), http.StatusBadRequest)
			return
		}

		tickets = append(tickets, ticket)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error during iteration\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    tickets,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func CreateTicket(w http.ResponseWriter, r *http.Request) {
	var ticket SchemaTicket
	err := json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		http.Error(w, "Error decoding request body\n"+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO ticket (user_id, title, description, category_id, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)", ticket.UserID, ticket.Title, ticket.Description, ticket.CategoryID, ticket.Status, ticket.CreatedAt, ticket.UpdatedAt)
	if err != nil {
		http.Error(w, "Error executing insert query\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func UpdateTicket(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	ticketIDStr := vars[len(vars)-1]

	ticketID, err := strconv.Atoi(ticketIDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	var ticket SchemaTicket
	err = json.NewDecoder(r.Body).Decode(&ticket)
	if err != nil {
		http.Error(w, "Error decoding request body\n"+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	_, err = db.Exec("UPDATE ticket SET user_id=?, title=?, description=?, category_id=?, status=?, created_at=?, updated_at=? WHERE id=?", ticket.UserID, ticket.Title, ticket.Description, ticket.CategoryID, ticket.Status, ticket.CreatedAt, ticket.UpdatedAt, ticketID)
	if err != nil {
		http.Error(w, "Error executing update query\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func DeleteTicket(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	ticketIDStr := vars[len(vars)-1]

	ticketID, err := strconv.Atoi(ticketIDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM ticket WHERE id=?", ticketID)
	if err != nil {
		http.Error(w, "Error executing delete query\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func RetrieveTicket(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	ticketIDStr := vars[len(vars)-1]

	ticketID, err := strconv.Atoi(ticketIDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM ticket WHERE id=?", ticketID)

	var ticket SchemaTicket
	err = row.Scan(&ticket.ID, &ticket.UserID, &ticket.Title, &ticket.Description, &ticket.CategoryID, &ticket.Status, &ticket.CreatedAt, &ticket.UpdatedAt)
	if err != nil {
		http.Error(w, "Error retrieving ticket\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    ticket,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}
