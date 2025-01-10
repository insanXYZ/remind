package main

func main() {

	err := CreateTempDir()
	if err != nil {
		WriteLog(err.Error())
		return
	}

	remindDatas, err := LoadData()
	if err != nil {
		WriteLog(err.Error())
		return
	}

	server := NewServer(remindDatas)
	err = server.Run()
	if err != nil {
		WriteLog(err.Error())
	}
}
