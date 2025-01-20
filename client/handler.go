package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/olekukonko/tablewriter"
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

	if res.Data == nil {
		fmt.Println("remind empty")
		return nil
	}

	sl, ok := res.Data.([]any)
	if !ok {
		return errors.New("error processing data")
	}

	var remindDatas [][]string

	for _, v := range sl {
		remindData, ok := v.(map[string]any)
		if !ok {
			return errors.New("error processing data")
		}

		var rowTable []string

		if v, ok := remindData["id"].(float64); ok {
			rowTable = append(rowTable, strconv.Itoa(int(v)))
		} else {
			rowTable = append(rowTable, "")
		}

		if v, ok := remindData["title"].(string); ok {
			rowTable = append(rowTable, v)
		} else {
			rowTable = append(rowTable, "")
		}

		if v, ok := remindData["name"].(string); ok {
			rowTable = append(rowTable, v)
		} else {
			rowTable = append(rowTable, "")
		}

		if v, ok := remindData["date"].(string); ok {
			rowTable = append(rowTable, v)
		} else {
			rowTable = append(rowTable, "")
		}

		if v, ok := remindData["time"].(string); ok {
			rowTable = append(rowTable, v)
		} else {
			rowTable = append(rowTable, "")
		}

		if v, ok := remindData["checked_at"].(string); ok {
			if v != "" {
				v = "  X  "
			}
			rowTable = append(rowTable, v)
		} else {
			rowTable = append(rowTable, "")
		}

		remindDatas = append(remindDatas, rowTable)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"id", "title", "name", "date", "time", "check"})
	table.AppendBulk(remindDatas)
	table.Render()

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
