package model

type RemindData struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Time      string `json:"time"`
	Date      string `json:"date"`
	CheckedAt string `json:"checked_at"`
}

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type CreateRequest struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Date string `json:"date"`
}
