package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

func ReadFile(filename string) ([]byte, error) {
	file, err := OpenFile(filename)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(file)
}

func ReadDir(path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}

func JoinPath(strings ...string) string {
	return filepath.Join(strings...)
}

func WriteLog(message string) {
	log.Println(message)
}

func WriteFatalLog(err error) {
	log.Fatal(err)
}

func OpenFile(filename string) (*os.File, error) {
	return os.OpenFile(JoinPath(ROOT_DIR, filename), os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
}

func WriteFile(filename, message string) error {

	message = message + "\n"

	file, err := OpenFile(filename)
	if err != nil {
		return err
	}

	if _, err = file.Write([]byte(message)); err != nil {
		return errors.Join(ErrWriteErrorLog, err)
	}
	return nil
}

func DeleteFile(filename string) error {
	return os.Remove(JoinPath(ROOT_DIR, filename))
}

func CreateTempDir() error {
	if _, err := ReadDir(ROOT_DIR); err != nil && errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(ROOT_DIR, os.ModePerm); err != nil {
			return errors.Join(ErrCreateTmpDir, err)
		}
	}
	WriteLog(SuccCreateTmpDir)
	return nil
}

func GetUnixTime(t time.Time) int64 {
	date := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return date.Unix()
}

func LoadData() (RemindDatas, error) {
	res := make(RemindDatas)
	now := GetUnixTime(time.Now())

	b, err := ReadFile(APP_DATA_FILENAME)
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(b)).Decode(&res)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return res, nil
		}
		return res, err
	}

	for _, v := range res {
		if v.Type == TASK {
			if v.CheckedAt != "" {
				t, err := time.Parse(time.DateOnly, v.CheckedAt)
				if err != nil {
					return nil, err
				}
				checkAt := GetUnixTime(t)

				if now > checkAt {
					this := res[v.Id]
					this.CheckedAt = ""
				}
			}

		}
	}

	err = DeleteFile(APP_DATA_FILENAME)
	if err != nil {
		return nil, err
	}

	b, err = json.Marshal(res)
	if err != nil {
		return nil, err
	}

	err = WriteFile(APP_DATA_FILENAME, string(b))
	if err != nil {
		return nil, err
	}

	return res, err

}
