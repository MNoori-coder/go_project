package question

type SchemaQuestion struct {
	ID      int    `json:"id" validate:"required"`
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}
