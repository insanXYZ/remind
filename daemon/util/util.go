package util

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"remind-daemon/model"
	"strings"
	"time"
)

func JoinPath(strings ...string) string {
	return filepath.Join(strings...)
}

func WriteLog(message string) {
	log.Println(message)
}

func WriteFile(filename string, message string, appendFlag bool) error {
	file, err := OpenFile(filename, appendFlag)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer file.Close()

	if _, err = file.WriteString(message); err != nil {
		log.Println(err.Error())
		return errors.Join(model.ErrWriteErrorLog, err)
	}
	return nil
}

func CreateTempDir() error {
	if _, err := ReadDir(model.ROOT_DIR); err != nil && errors.Is(err, os.ErrNotExist) {
		if err = os.Mkdir(model.ROOT_DIR, os.ModePerm); err != nil {
			log.Println(err.Error())
			return errors.Join(model.ErrCreateTmpDir, err)
		}
	}
	WriteLog(model.SuccCreateTmpDir)
	return nil
}

func OpenFile(filename string, appendFlag bool) (*os.File, error) {
	flag := os.O_RDWR | os.O_CREATE
	if appendFlag {
		flag |= os.O_APPEND
	} else {
		flag |= os.O_TRUNC
	}

	return os.OpenFile(JoinPath(model.ROOT_DIR, filename), flag, os.ModePerm)
}

func ReadFile(filename string) ([]byte, error) {
	file, err := OpenFile(filename, true)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer file.Close()
	return io.ReadAll(file)
}

func ReadDir(path string) ([]os.DirEntry, error) {
	return os.ReadDir(path)
}

func SendNotif(title, name string) error {
	cmd := exec.Command("bash", "-c", title, name)
	return cmd.Run()
}

func StructToJsonString(data any) (string, error) {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	return string(b), err
}

func TrimSpace(s string) string {
	return strings.Trim(s, " ")
}

func GetStartOfDay() time.Time {
	t, _ := time.Parse(time.DateOnly, time.Now().Format(time.DateOnly))
	return t
}
