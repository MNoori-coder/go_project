package ticket

type SchemaTicket struct {
	ID          int    `json:"id" validate:"required"`
	UserID      int    `json:"user_id" validate:"required"`
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CategoryID  int    `json:"category_id" validate:"required"`
	Status      string `json:"status" validate:"required"`
	CreatedAt   string `json:"created_at" validate:"required"`
	UpdatedAt   string `json:"updated_at" validate:"required"`
}
