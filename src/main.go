package main

import (
	"os"
	"log"
	"strings"
	"path/filepath"
)

const docoptAppVer = "{{ .Name }} v{{ .Version }}"

const docoptString = docoptAppVer + `

Usage:
  {{ .Arg0 }} new <name>
  {{ .Arg0 }} list <name>
  {{ .Arg0 }} add <name> <task>...
  {{ .Arg0 }} del <name> <line>
  {{ .Arg0 }} setpriority <name> <line> <priority>
  {{ .Arg0 }} setproject <name> <line> (+|-) <key>
  {{ .Arg0 }} setcontext <name> <line> (+|-) <key>
  {{ .Arg0 }} settag <name> <line> (+|-) <key> <value>
  {{ .Arg0 }} complete <name> <line>
  {{ .Arg0 }} reopen <name> <line>
  {{ .Arg0 }} sort <name> [context | project | priority]
  {{ .Arg0 }} -h | --help
  {{ .Arg0 }} -v | --version

Options:
  -h --help     Show this screen
  -v --version  Show version
  <name>        Name of Todo.md file in .todo.md
  <task>        Task; as long as you like, but only one line
  <line>        Line as shown in list
  <priority>    Letter priority, A-Z descending.
  <key>         Indexable key; any text but no spaces
  <value>       Any text, spaces allowed
`

func route(logs chan string, args argMap, verb string, fn actFunc) {
	if args[verb].(bool) {
		fn(logs, args)
	}
}

// If help or version passed, args will be empty.
func threadMainLoop(args argMap, logs chan string) {
	defer close(logs)
	if len(args) == 0 {
		return
	}
	route(logs, args, "new", fnNew)
	route(logs, args, "list", fnList)
	route(logs, args, "add", fnAdd)
	route(logs, args, "del", fnDel)
	route(logs, args, "setpriority", fnSetPriority)
	route(logs, args, "setproject", fnSetProject)
	route(logs, args, "setcontext", fnSetContent)
	route(logs, args, "settag", fnSetTag)
	route(logs, args, "complete", fnComplete)
	route(logs, args, "reopen", fnReopen)
	route(logs, args, "sort", fnSort)
}

func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func main() {
	dotv := struct {
		Name string
		Version string
		Arg0 string
	}{APP_NAME, VERSION, filepath.Base(os.Args[0])}
	args, err := runDocopt(dotv)
	if err != nil {
		log.Println(err)
		return
	}
	logs := make(chan string)
	go threadMainLoop(args, logs)
	for msg := range logs {
		msgs := strings.Split(msg, "\\n")
		for _, m := range removeEmptyStrings(msgs) {
			log.Println(m)
		}
	}
}
