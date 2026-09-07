package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

/*
vi /var/lib/dpkg/status
Package: pika-baseos-minimal
find nano in Depends and delete
*/

func getListFromFile(fn string) ([]string, error) {
	listBytes, err := ioutil.ReadFile(fn)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(listBytes), "\n")
	size := len(lines)
	if size > 0 {
		lines = lines[0:size - 1]
	}
	return lines, nil
}

func putListToFile(fn string, lines []string) error {
	fileAsString := strings.Join(lines, "\n") + "\n"
	return os.WriteFile(fn, []byte(fileAsString), 0666)
}

func out(n int, msg string) {
	fmt.Println(fmt.Sprintf("%5d: %s", n ,msg))
}

func setDone(line string, done bool) string {
	re := Task{}.Load(line)
	re.Done = done
	ret := re.Save()
	return ret
}

func encodeDone(re Task) string {
	if re.Done {
		return "[x]"
	}
	return "[ ]"
}

func encodePriority(re Task) string {
	if re.Priority == -1 {
		return ""
	}
	return fmt.Sprintf(" (%c)", rune('A' + re.Priority))
}

func encodeDate(i int) string {
	if i == -1 {
		return ""
	}
	return " 2025-11-29" // DEBUG
}

func encodeProjects(re Task) string {
	if len(re.Projects) == 0 {
		return ""
	}
	return " +" + strings.Join(re.Projects, " +")
}

func encodeContexts(re Task) string {
	if len(re.Contexts) == 0 {
		return ""
	}
	return " @" + strings.Join(re.Contexts, " @")
}

func encodeTags(re Task) string {
	if len(re.Tags) == 0 {
		return ""
	}
	tags := ""
	for k, v := range re.Tags {
		if !strings.Contains(v, " ") {
			tags += fmt.Sprintf(" %s:%s", k, v)
//		} else {
//			tags += fmt.Sprintf(" {%s:%s}", k, v)
		}
	}
	return tags
}

// tags (projects, contexts, tags) are part of text
func decodeText(re Task, s string) Task {
	removePrefixLength := 4 // done status
	if re.Priority != -1 {
		removePrefixLength += 4
	}
	if re.CreationDate != -1 {
		removePrefixLength += 11
		if re.CompletionDate != -1 {
			removePrefixLength += 11
		}
	}
	re.Text = s[removePrefixLength:]
	return re
}

func decodeDone(re Task, s string) Task {
	re.Done = (s[1] == 'x')
	return re
}

func decodePriority(re Task, s string) Task {
	re.Priority = -1
	if s[4] == '(' && s[6:8] == ") " {
		re.Priority = int(s[5] - 'A')
	}
	return re
}

// note, these can appear in the middle of text.
// tags are part of text
// when rendering, just put out text, not tags --- but read them anyway just in case
func decodeProjects(re Task, s string) Task {
	for _, v := range strings.Split(s, " ") {
		if v[0] == '+' {
			re.Projects = append(re.Projects, v[1:])
		}
	}
	return re
}

// note, these can appear in the middle of text.
func decodeContexts(re Task, s string) Task {
	for _, v := range strings.Split(s, " ") {
		if v[0] == '@' {
			re.Contexts = append(re.Contexts, v[1:])
		}
	}
	return re
}

// note, these can appear in the middle of text.
func decodeTags(re Task, s string) Task {
//	tagOpen := false
//	for _, v := range strings.Split(s, " ") {
//		// !!! too tired for this one
//	}
	return re
}

func decodeCompletion(re Task, s string) Task {
	re.CompletionDate = -1
	return re
}

func decodeCreation(re Task, s string) Task {
	re.CreationDate = -1
	return re
}

//
func modifyTag(
		line string,
		enable bool,
		name string,
		symbol byte) string {
	//re = re.changeProjects(enable, name)
	return line
}
//

/*
func changeContexts(re Task, enable bool, name string) Task {
	if enable {
		re.Context = append(re.Context, name)
	} else {
		n := len(re.Context)
		if n > 0 {
			temp := 
		}
	}
	return re
}

func changePriority(re Task, c byte) Task {
	re.Priority = int(c - 'A')
	return re
}

func changeProjects(re Task, enable bool, name string) Task {
	if enable {
		re.Project = append(re.Project, name)
	} else {
		n := len(re.Project)
		if n > 0 {
			temp := 
		}
	}
	return re
}
*/
