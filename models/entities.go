package models

type ToDo struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type APIGateway struct {
	Id   int64
	Type string
}
