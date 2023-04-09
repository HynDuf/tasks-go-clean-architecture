package entity

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title" form:"title" binding:"required"`
	UserID    int    `json:"userID"`
	Completed bool   `json:"completed`
}
