package main

func main() {
	err := CreateTempDir()
	if err != nil {
		WriteFatalLog(err)
		return
	}

	remindDatas, lastId, err := LoadData()
	if err != nil {
		WriteFatalLog(err)
		return
	}

	server := NewServer(remindDatas, lastId)
	err = server.Run()
	if err != nil {
		WriteFatalLog(err)
	}

}
