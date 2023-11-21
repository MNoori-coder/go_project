package ticket_category

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ListTicketCategory(w http.ResponseWriter, r *http.Request) {
	var ticket_categories []SchemaTicketCategory

	db, err := sql.Open("sqlite3", "my_db.db")
	if err != nil {
		http.Error(w, "Can not connect to db\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM ticket_category")
	if err != nil {
		http.Error(w, "Can not execute query\n"+err.Error(), http.StatusBadRequest)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var ticket_category SchemaTicketCategory
		err := rows.Scan(&ticket_category.ID, &ticket_category.Name, &ticket_category.Content)
		if err != nil {
			http.Error(w, "Can not scan row data\n"+err.Error(), http.StatusBadRequest)
			return
		}

		ticket_categories = append(ticket_categories, ticket_category)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error during iteration\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    ticket_categories,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}

func CreateTicketCategory(w http.ResponseWriter, r *http.Request) {
	var ticket_category SchemaTicketCategory
	err := json.NewDecoder(r.Body).Decode(&ticket_category)
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

	_, err = db.Exec("INSERT INTO ticket_category (name, content) VALUES (?, ?)", ticket_category.Name, ticket_category.Content)
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

func UpdateTicketCategory(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	ticket_category_IDStr := vars[len(vars)-1]

	ticket_category_ID, err := strconv.Atoi(ticket_category_IDStr)
	if err != nil {
		http.Error(w, "Invalid ticket ID\n"+err.Error(), http.StatusBadRequest)
		return
	}

	var ticket_category SchemaTicketCategory
	err = json.NewDecoder(r.Body).Decode(&ticket_category)
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

	_, err = db.Exec("UPDATE ticket_category SET name=?, content=? WHERE id=?", ticket_category.Name, ticket_category.Content, ticket_category_ID)
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

func DeleteTicketCategory(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	ticket_category_IDStr := vars[len(vars)-1]

	ticket_category_ID, err := strconv.Atoi(ticket_category_IDStr)
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

	_, err = db.Exec("DELETE FROM ticket_category WHERE id=?", ticket_category_ID)
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

func RetrieveTicketCategory(w http.ResponseWriter, r *http.Request) {
	vars := strings.Split(strings.TrimSuffix(r.URL.Path, "/"), "/")
	ticket_category_IDStr := vars[len(vars)-1]

	ticket_category_ID, err := strconv.Atoi(ticket_category_IDStr)
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

	row := db.QueryRow("SELECT * FROM ticket_category WHERE id=?", ticket_category_ID)

	var ticket_category SchemaTicketCategory
	err = row.Scan(&ticket_category.ID, &ticket_category.Name, &ticket_category.Content)
	if err != nil {
		http.Error(w, "Error retrieving ticket\n"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"data":    ticket_category,
		"success": true,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response to JSON\n", http.StatusInternalServerError)
		return
	}
}
