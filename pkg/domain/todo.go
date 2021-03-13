package domain

type Todo struct {
	ID        string `json:"id" docstore:"id"`
	Content   string `json:"content" docstore:"content"`
	Completed bool   `json:"completed" docstore:"completed"`
}
