package ticket_category

type SchemaTicketCategory struct {
	ID      int    `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required"`
}
