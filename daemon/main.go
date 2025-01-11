package main

func main() {

	err := CreateTempDir()
	if err != nil {
		WriteFatalLog(err)
		return
	}

	remindDatas, err := LoadData()
	if err != nil {
		WriteFatalLog(err)
		return
	}

	server := NewServer(remindDatas)
	err = server.Run()
	if err != nil {
		WriteFatalLog(err)
	}
}
