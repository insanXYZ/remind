package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"
)

var (
	ErrInformation     = errors.New("use --help for usage information")
	ErrMissingArgument = errors.New("missing argument")
	ErrWrongArgument   = errors.New("wrong argument")
)

type Flag = string

const (
	SET    Flag = "set"
	DELETE      = "delete"
	CHECK       = "check"
	LS          = "ls"
	HELP        = "help"
)

type FlagAttr struct {
	setName       string
	setTime       string
	setDate       string
	deleteId      string
	checkId       string
	checkIsRemove bool
}

func main() {
	flagAttr := new(FlagAttr)
	flagUsed, err := initFlag(flagAttr)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	err = ExecuteHandler(flagUsed, flagAttr)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func initFlag(flagAttr *FlagAttr) (Flag, error) {

	flagMap := make(map[string]*flag.FlagSet)

	setFlag := flag.NewFlagSet(SET, flag.ExitOnError)
	flagMap[SET] = setFlag
	setFlag.StringVar(&flagAttr.setName, "name", "", "assign title and name, example john doe")
	setFlag.StringVar(&flagAttr.setDate, "date", "", "assign date ,example 2005-01-16")
	setFlag.StringVar(&flagAttr.setTime, "time", "", fmt.Sprintf("assign clock, example %s", time.Now().Format(time.TimeOnly)))

	deleteFlag := flag.NewFlagSet(DELETE, flag.ExitOnError)
	flagMap[DELETE] = deleteFlag
	deleteFlag.StringVar(&flagAttr.deleteId, "id", "", "assign id for delete")

	checkFlag := flag.NewFlagSet(CHECK, flag.ExitOnError)
	flagMap[CHECK] = checkFlag
	checkFlag.StringVar(&flagAttr.checkId, "id", "", "assign id for checked")
	checkFlag.BoolVar(&flagAttr.checkIsRemove, "rm", false, "optional flag for remove check on remind")

	lsFlag := flag.NewFlagSet(LS, flag.ExitOnError)
	flagMap[LS] = lsFlag

	helpFlag := flag.NewFlagSet(HELP, flag.ExitOnError)
	flagMap[HELP] = helpFlag

	if len(os.Args) < 2 {
		return "", errors.Join(ErrMissingArgument, ErrInformation)
	}

	flagUsed := os.Args[1]
	f, ok := flagMap[flagUsed]

	if !ok {
		return "", errors.Join(ErrWrongArgument, ErrInformation)
	}

	return flagUsed, f.Parse(os.Args[2:])

}
