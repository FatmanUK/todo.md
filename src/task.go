package main

import (
	"os"
	"fmt"
	"strings"
)

type Task struct {
	Text string
	Done bool
	Priority int
	Projects []string
	Contexts []string
	Tags map[string]string
	CompletionDate int
	CreationDate int
}

func (re Task) Save() string {
	line := fmt.Sprintf(
		"%s%s%s%s %s%s%s%s",
		encodeDone(re),
		encodePriority(re),
		encodeDate(re.CompletionDate),
		encodeDate(re.CreationDate),
		re.Text,
		encodeProjects(re),
		encodeContexts(re),
		encodeTags(re),
	)
	return line
}

func (re Task) Load(s string) Task {
	re = decodeDone(re, s)
	re = decodePriority(re, s)
	re = decodeProjects(re, s)
	re = decodeContexts(re, s)
	re = decodeTags(re, s)
	re = decodeCompletion(re, s)
	re = decodeCreation(re, s)
	re = decodeText(re, s)
	return re
}

func (re Task) New(fn string) error {
	_, err := os.Create(fn)
	return err
}

func (re Task) List(fn string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	for n, line := range lines {
		out(n+1, line)
	}
	return nil
}

func (re Task) Add(fn string, task []string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	re = re.Load("[ ] " + strings.Join(task, " "))
	lines = append(lines, re.Save())
	err = putListToFile(fn, lines)
	if err != nil {
		return err
	}
	return nil
}

func (re Task) Del(fn string, id int) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	var newLines []string
	for _, line := range lines {
		id--
		if id != 0 {
			newLines = append(newLines, line)
		}
	}
	newLines = append(newLines, "")
	err = putListToFile(fn, newLines)
	if err != nil {
		return err
	}
	return nil
}

func (re Task) SetComplete(fn string, id int, done bool) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	var newLines []string
	for _, line := range lines {
		id--
		if id == 0 {
			line = setDone(line, done)
		}
		newLines = append(newLines, line)
	}
	err = putListToFile(fn, newLines)
	if err != nil {
		return err
	}
	return nil
}

//
// the bit at the start with the parens, not the later bit with the +
func (re Task) SetPriority(fn string, id int, pri string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	var newLines []string
	for _, line := range lines {
		id--
//		if id == 0 {
//			line = modifyLine(line, enable, name, '+')
//		}
		newLines = append(newLines, line)
	}
	newLines = append(newLines, "")
	err = putListToFile(fn, newLines)
	if err != nil {
		return err
	}
	return nil
}
//

//
func (re Task) SetProject(
		fn string,
		id int,
		enable bool,
		name string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	var newLines []string
	for _, line := range lines {
		id--
		if id == 0 {
			line = modifyTag(line, enable, name, '+')
		}
		newLines = append(newLines, line)
	}
	newLines = append(newLines, "")
	err = putListToFile(fn, newLines)
	if err != nil {
		return err
	}
	return nil
/*
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	for n, line := range lines {
		out(n+1, line)
	}
*/
/*
	t := []string{}
	for id, line := range lines {
		id--
		if id == 0 {
			re = re.Load(line)
			re = re.changeProjects(enable, name)
			line = re.Save()
		}
		t = append(t, line)
	}
	fileAsLongString := strings.Join(t, "\n") + "\n"
	err = os.WriteFile(fn, []byte(fileAsLongString), 0666)
	if err != nil {
		return err
	}
*/
}
//

//
func (re Task) SetContext(
		fn string,
		id int,
		enable bool,
		name string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	var newLines []string
	for _, line := range lines {
		id--
		if id == 0 {
			line = modifyTag(line, enable, name, '@')
		}
		newLines = append(newLines, line)
	}
	newLines = append(newLines, "")
	err = putListToFile(fn, newLines)
	if err != nil {
		return err
	}
	return nil
/*
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	for n, line := range lines {
		out(n+1, line)
	}
*/
/*
	t := []string{}
	for _, line := range lines {
		id--
		if id == 0 {
			re = re.Load(line)
			re = re.changeContexts(enable, name)
			line = re.Save()
		}
		t = append(t, line)
	}
	fileAsLongString := strings.Join(t, "\n") + "\n"
	err = os.WriteFile(fn, []byte(fileAsLongString), 0666)
	if err != nil {
		return err
	}
*/
}
//

//
func (re Task) SetTag(
		fn string,
		id int,
		enable bool,
		key string,
		value string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	for n, line := range lines {
		out(n+1, line)
	}
//
	return nil
}
//

//
func (re Task) SortByPriority(fn string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	for n, line := range lines {
		out(n+1, line)
	}

	return nil
}
//

//
// hmm. How do we define x>y for this?
func (re Task) SortByContext(fn string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	for n, line := range lines {
		out(n+1, line)
	}

	return nil
}
//

//
// hmm. How do we define x>y for this?
func (re Task) SortByProject(fn string) error {
	lines, err := getListFromFile(fn)
	if err != nil {
		return err
	}
	for n, line := range lines {
		out(n+1, line)
	}

	return nil
}
//
