package server

import (
	"errors"
)

var (
	ErrCreateTmpDir         = errors.New("failed create temporary directory")
	ErrWriteErrorLog        = errors.New("failed write error log")
	ErrGetCacheRemindData   = errors.New("failed get remind data")
	ErrSetRemind            = errors.New("failed set remind")
	ErrWrongType            = errors.New("wrong type")
	ErrWrongDate            = errors.New("wrong date")
	ErrValidateNameRequired = errors.New("name required")
	SuccCreateTmpDir        = "success create tmp dir"
	SuccRunServer           = "success running server"
	SuccSetRemind           = "success set remind"
)
