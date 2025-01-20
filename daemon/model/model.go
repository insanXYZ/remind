package model

type RemindData struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Name      string `json:"name"`
	Time      string `json:"time"` //time.Timeonly
	Date      string `json:"date"` //time.Dateonly
	CheckedAt string `json:"checked_at"`
}

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type SetRequest struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Date string `json:"date"`
}
