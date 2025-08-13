package dto

// ListTasksQuery описывает параметры фильтрации и пагинации, которые клиент передаёт в query-строке запроса
type ListTasksQuery struct {
	Status *string `json:"status"`
	Limit  int     `json:"limit"`
	Offset int     `json:"offset"`
}
