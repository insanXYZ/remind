package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func JoinPath(strings ...string) string {
	return filepath.Join(strings...)
}

func WriteLog(message string) {
	log.Println(message)
}

func WriteFatalLog(err error) {
	log.Fatal(err)
}

func WriteFile(filename string, message any, appendFlag bool) error {

	var b []byte
	var err error

	if filepath.Ext(filename) == ".json" {
		b, err = json.Marshal(message)
		if err != nil {
			return err
		}
	} else {
		b = []byte(message.(string))
	}

	file, err := OpenFile(filename, appendFlag)
	if err != nil {
		return err
	}

	if _, err = file.Write(b); err != nil {
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

func LoadData() (RemindDatas, int, error) {
	changes := false
	res := make(RemindDatas)
	now := GetUnixTime(time.Now())
	lastId := 0

	b, err := ReadFile(APP_DATA_FILENAME)
	if err != nil {
		return nil, lastId, err
	}

	err = json.NewDecoder(bytes.NewReader(b)).Decode(&res)
	if err != nil {
		if errors.Is(err, io.EOF) {
			return res, lastId, nil
		}
		return res, lastId, err
	}

	for _, v := range res {
		if v.Id > lastId {
			lastId = v.Id
		}
		if v.Type == TASK {
			if v.CheckedAt != "" {
				t, err := time.Parse(time.DateOnly, v.CheckedAt)
				if err != nil {
					return nil, lastId, err
				}
				checkAt := GetUnixTime(t)

				if now > checkAt {
					changes = true
					res[v.Id].CheckedAt = ""
				}
			}

		}
	}

	if changes {

		err = WriteFile(APP_DATA_FILENAME, res, false)
		if err != nil {
			return nil, lastId, err
		}
	}

	return res, lastId, err

}

func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}
