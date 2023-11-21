package ticket_comments

type SchemaTicketComments struct {
	ID          int    `json:"id" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
	SupporterID int    `json:"supporter_id" validate:"required"`
	Message     string `json:"message" validate:"required"`
	CreatedAt   string `json:"created_at" validate:"required"`
	UpdatedAt   string `json:"updated_at" validate:"required"`
}
