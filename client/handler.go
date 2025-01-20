package main

import (
	"fmt"
	"net/http"
)

func ExecuteHandler(flagUsed Flag, attr *FlagAttr) error {
	handlerMap := make(map[Flag]func(*FlagAttr) error)
	handlerMap[HELP] = HelpHandler
	handlerMap[SET] = SetHandler
	handlerMap[DELETE] = DeleteHandler
	handlerMap[LS] = LsHandler

	f, ok := handlerMap[flagUsed]

	if !ok {
		return ErrWrongArgument
	}

	return f(attr)

}

func LsHandler(attr *FlagAttr) error {
	res, err := CreateRequest(http.MethodGet, "/", nil)
	if err != nil {
		return err
	}

	return nil
}

func SetHandler(attr *FlagAttr) error {
	return nil
}

func DeleteHandler(attr *FlagAttr) error {
	return nil
}

func HelpHandler(_ *FlagAttr) error {
	fmt.Println(`
Usage: remind COMMAND [OPTION]

Commands:
set		Create a new remind .
        	Format: remind set --name "NAME" [--date "YYYY-MM-DD"] [--time "HH:MM"]
		Options:
		--name		Specify the name of the reminder (required).
		--date    	Specify the date for the reminder (optional). Default: today.
		--time    	Specify the time for the reminder (optional).

check		Mark a remind as completed.
		Format: remind check --id ID
		Options:
		--u		Uncheck mark (optional).	
		--id      	Specify the ID of the remind (required).

delete  	Remove a remind.
		Format: remind delete --id ID
		Options:
		--id      	Specify the ID of the remind (required).

ls      	List all reminders.
		Format: remind ls

Other options:
--help  	Show this usage information.`)
	return nil
}
