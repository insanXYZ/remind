package main

import "errors"

var (
	ErrCreateTmpDir       = errors.New("failed create temporary directory")
	ErrWriteErrorLog      = errors.New("failed write error log")
	ErrGetCacheRemindData = errors.New("failed get remind data")
	SuccCreateTmpDir      = "success create tmp dir"
	SuccRunServer         = "success running server"
)
