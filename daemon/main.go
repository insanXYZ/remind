package main

type RemindType string

const (
	ALARM RemindType = "alarm"
	TASK  RemindType = "task"
)

const (
	APP_TEMP_DIRNAME  = ".remind-tmp"
	APP_DATA_FILENAME = "remind-data.json"
	APP_TASK_FILENAME = "remind-tasks.json"
)

type RemindData struct {
	Id    int        `json:"id"`
	Name  string     `json:"name"`
	Time  string     `json:"time"`
	Type  RemindType `json:"type"`
	Extra string     `json:"extra"`
}

func main() {

	loaded := loadRemindData()

	server := NewServer(loaded)
	server.Run()
}

func loadRemindData() map[int]RemindData {
	return nil
}
