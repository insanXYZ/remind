package main

import (
	"os"
)

const (
	ALARM RemindType = "alarm"
	TASK  RemindType = "task"
)

var (
	APP_TEMP_DIRNAME       = ".remind-tmp"
	APP_DATA_FILENAME      = "remind-data.json"
	APP_ERROR_LOG_FILENAME = "remind-error.log"
	ROOT_DIR               = JoinPath(os.Getenv("HOME"), APP_TEMP_DIRNAME)
)

type (
	RemindType  string
	RemindDatas = map[int]RemindData
	Id          = int
)

type RemindData struct {
	Id        int        `json:"id"`
	Title     string     `json:"title"`
	Name      string     `json:"name"`
	Time      string     `json:"time"`
	Type      RemindType `json:"type"`
	Extra     string     `json:"extra"`
	CheckedAt string     `json:"checked_at"`
}
