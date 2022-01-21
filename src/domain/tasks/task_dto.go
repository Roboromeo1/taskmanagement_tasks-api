package tasks

type Task struct {
	Id            string      `json:"id"`
	DateCreated   string      `json:"date_created"`
	DueDate       string      `json:"dueDate"`
	AllocatedUser int64       `json:"allocated_user"`
	Status        Status      `json:"status"`
	Description   Description `json:"description"`
}

type Description struct {
	PlainText string `json:"plain_text"`
	Html      string `json:"html"`
}
type Status struct {
	NotUrgent string `json:"not_urgent"`
	DueSoon   string `json:"due_soon"`
	Overdue   string `json:"overdue"`
}
