package main

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type SetRequest struct {
	Name string `json:"name"`
	Time string `json:"time"`
	Date string `json:"date"`
}
