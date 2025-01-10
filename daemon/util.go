package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
)

func ReadFile(filepath string) ([]byte, error) {
	return os.ReadFile(filepath)
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

func WriteFile(filename, message string) error {
	file, err := os.OpenFile(JoinPath(ROOT_DIR, filename), os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}

	_, err = file.Write([]byte(message))
	return errors.Join(ErrWriteErrorLog, err)
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

func LoadData() (RemindDatas, error) {
	res := make(RemindDatas)

	b, err := ReadFile(filepath.Join(ROOT_DIR, APP_DATA_FILENAME))
	if err != nil {
		return nil, err
	}

	err = json.NewDecoder(bytes.NewReader(b)).Decode(&res)

	return res, err

}
