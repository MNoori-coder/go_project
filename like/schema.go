package like

type SchemaLike struct {
	ID       int `json:"id" validate:"required"`
	UserID   int `json:"user_id" validate:"required"`
	TicketID int `json:"ticket_id" validate:"required"`
	IsLike   int `json:"is_like" validate:"required"`
}
