package main

import (
	"os"
	"fmt"
	"bufio"
	"errors"
	"strconv"
)

type argMap = map[string]interface{}

type actFunc = func(chan string, argMap)

func ensureDir(path string) error {
	if path[0:1] != "/" {
		return errors.New("must be full path")
	}
	return os.MkdirAll(path, 0755)
}

func getDataDir() string {
	return os.Getenv("HOME") + "/" + DATAPATH
}

func getDataFile(listname string) string {
	datadir := getDataDir()
	return fmt.Sprintf("%s/%s.todo.md", datadir, listname)
}

func getOnOff(pos bool, neg bool) bool {
	if pos && !neg {
		return true
	}
	if !pos && neg {
		return false
	}
	return true // ?
}

func fnShowById(logs chan string, args argMap, id int) {
	fn := getDataFile(args["<name>"].(string))
	f, err := os.OpenFile(fn, os.O_RDWR, 0666)
	if err != nil {
		logs <- err.Error()
		return
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		id--
		if id == 0 {
			fmt.Println(scanner.Text())
			break
		}
	}
	err = scanner.Err()
	if err != nil {
		logs <- err.Error()
		return
	}
}

func fnNew(logs chan string, args argMap) {
	ensureDir(getDataDir())
	fn := getDataFile(args["<name>"].(string))
	err := Task{}.New(
		fn,
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	logs <- "New todo.md file: " + fn
}

func fnList(logs chan string, args argMap) {
	err := Task{}.List(
		getDataFile(args["<name>"].(string)),
	)
	if err != nil {
		logs <- err.Error()
		return
	}
}

func fnAdd(logs chan string, args argMap) {
	err := Task{}.Add(
		getDataFile(args["<name>"].(string)),
		args["<task>"].([]string),
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	fnList(logs, args)
}

func fnDel(logs chan string, args argMap) {
	id, err := strconv.Atoi(args["<line>"].(string))
	if err != nil {
		logs <- err.Error()
		return
	}
	err = Task{}.Del(
		getDataFile(args["<name>"].(string)),
		id,
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	fnList(logs, args)
}

func fnSetPriority(logs chan string, args argMap) {
	id, err := strconv.Atoi(args["<line>"].(string))
	if err != nil {
		logs <- err.Error()
		return
	}
	err = Task{}.SetPriority(
		getDataFile(args["<name>"].(string)),
		id,
		args["<priority>"].(string),
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	fnShowById(logs, args, id)
}

func fnSetProject(logs chan string, args argMap) {
	id, err := strconv.Atoi(args["<line>"].(string))
	if err != nil {
		logs <- err.Error()
		return
	}
	err = Task{}.SetProject(
		getDataFile(args["<name>"].(string)),
		id,
		getOnOff(args["+"].(bool), args["-"].(bool)),
		args["<key>"].(string),
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	fnShowById(logs, args, id)
}

func fnSetContent(logs chan string, args argMap) {
	id, err := strconv.Atoi(args["<line>"].(string))
	if err != nil {
		logs <- err.Error()
		return
	}
	err = Task{}.SetContext(
		getDataFile(args["<name>"].(string)),
		id,
		getOnOff(args["+"].(bool), args["-"].(bool)),
		args["<key>"].(string),
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	fnShowById(logs, args, id)
}

func fnSetTag(logs chan string, args argMap) {
	id, err := strconv.Atoi(args["<line>"].(string))
	if err != nil {
		logs <- err.Error()
		return
	}
	err = Task{}.SetTag(
		getDataFile(args["<name>"].(string)),
		id,
		getOnOff(args["+"].(bool), args["-"].(bool)),
		args["<key>"].(string),
		args["<value>"].(string),
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	fnShowById(logs, args, id)
}

func fnComplete(logs chan string, args argMap) {
	id, err := strconv.Atoi(args["<line>"].(string))
	if err != nil {
		logs <- err.Error()
		return
	}
	err = Task{}.SetComplete(
		getDataFile(args["<name>"].(string)),
		id,
		true,
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	fnShowById(logs, args, id)
}

func fnReopen(logs chan string, args argMap) {
	id, err := strconv.Atoi(args["<line>"].(string))
	if err != nil {
		logs <- err.Error()
		return
	}
	err = Task{}.SetComplete(
		getDataFile(args["<name>"].(string)),
		id,
		false,
	)
	if err != nil {
		logs <- err.Error()
		return
	}
	fnShowById(logs, args, id)
}

func fnSort(logs chan string, args argMap) {
	var err error = nil
	if args["context"].(bool) {
		err = Task{}.SortByContext(
			getDataFile(args["<name>"].(string)),
		)
	}
	if args["priority"].(bool) {
		err = Task{}.SortByPriority(
			getDataFile(args["<name>"].(string)),
		)
	}
	if args["project"].(bool) {
		err = Task{}.SortByProject(
			getDataFile(args["<name>"].(string)),
		)
	}
	if err != nil {
		logs <- err.Error()
		return
	}
	fnList(logs, args)
}
