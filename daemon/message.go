package main

import "errors"

var (
	ErrCreateTmpDir  = errors.New("failed create temporary directory")
	ErrWriteErrorLog = errors.New("failed write error log")
	SuccCreateTmpDir = "success create tmp dir"
)
