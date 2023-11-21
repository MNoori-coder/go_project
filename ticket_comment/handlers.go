package ticket_comment

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ListTicketComments(w http.ResponseWriter, r *http.Request) {
	var ticket_comments []SchemaTicketComment

	// Open SQLite database connection
	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	// Query all records from the camera table
	rows, err := db.Query("SELECT * FROM ticket_comment")
	if err != nil {
		http.Error(w, "Can not execute query\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ticket_comment SchemaTicketComment
		err := rows.Scan(&ticket_comment.ID, &ticket_comment.UserID, &ticket_comment.SupporterID, &ticket_comment.Message, &ticket_comment.CreatedAt, &ticket_comment.UpdatedAt)
		if err != nil {
			http.Error(w, "Can not scan row data\n"+err.Error(), http.StatusBadRequest)
			return
		}

		ticket_comments = append(ticket_comments, ticket_comment)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error during iteration\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    ticket_comments,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func CreateTicketComment(w http.ResponseWriter, r *http.Request) {
	var ticket_comment SchemaTicketComment
	err := json.NewDecoder(r.Body).Decode(&ticket_comment)
	if err != nil {
		http.Error(w, "Error decoding request body\n"+err.Error(), http.StatusBadRequest)
		return
	}

	// Open SQLite database connection
	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	// Insert the new ticket into the ticket table
	_, err = db.Exec("INSERT INTO ticket_comment (user_id, supporter_id, message, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", ticket_comment.UserID, ticket_comment.SupporterID, ticket_comment.Message, ticket_comment.CreatedAt, ticket_comment.UpdatedAt)
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

func UpdateTicketComment(w http.ResponseWriter, r *http.Request) {
	// Extract the ticket ID from the URL path
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	ticket_comment_IDStr := vars[len(vars)-1]

	ticket_comment_ID, err := strconv.Atoi(ticket_comment_IDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	var ticket_comment SchemaTicketComment
	err = json.NewDecoder(r.Body).Decode(&ticket_comment)
	if err != nil {
		http.Error(w, "Error decoding request body\n"+err.Error(), http.StatusBadRequest)
		return
	}

	// Open SQLite database connection
	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	// Update the ticket in the ticket table
	_, err = db.Exec("UPDATE ticket_comment SET user_id=?, supporter_id=?, message=?, created_at=?, updated_at=? WHERE id=?", ticket_comment.UserID, ticket_comment.SupporterID, ticket_comment.Message, ticket_comment.CreatedAt, ticket_comment.UpdatedAt, ticket_comment_ID)
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
