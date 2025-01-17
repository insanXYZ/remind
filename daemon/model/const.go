package model

import (
	"os"
	"path/filepath"
)

const (
	EVERY_DAY_DATE = "every-day"
)

var (
	APP_TEMP_DIRNAME       = ".remind-tmp"
	APP_DATA_FILENAME      = "remind-data.json"
	APP_ERROR_LOG_FILENAME = "remind-error.log"
	ROOT_DIR               = filepath.Join(os.Getenv("HOME"), APP_TEMP_DIRNAME)
)

type (
	RemindDatas = map[int]*RemindData
)
