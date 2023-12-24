package main

import (
    "strings"
    "os"
    "github.com/chzyer/readline"
)

// getServerName returns a slice of readline.PrefixCompleterInterface
// that represents the available server names for auto-completion.
func getServerName() []readline.PrefixCompleterInterface {
	return []readline.PrefixCompleterInterface{
		readline.PcItem("server1"),
		readline.PcItem("server2"),
		readline.PcItem("server3"),
	}
}

var completer = readline.NewPrefixCompleter(
    readline.PcItem("server", getServerName()...),
    readline.PcItem("command2"),
    readline.PcItem("show",
		readline.PcItem("config"),
		readline.PcItem("running-config"),
	),
	readline.PcItem("exit"),
)

// main is the entry point of the program.
// It initializes the readline library, sets up auto-completion,
// and reads user input until the user enters "exit".
func main() {
    rl, err := readline.NewEx(&readline.Config{
        Prompt:          "prompt> ",
        HistoryFile:     "/tmp/readline.tmp",
        AutoComplete:    completer,
        InterruptPrompt: "^C",
        EOFPrompt:       "exit",
    })

    if err != nil {
        panic(err)
    }
    defer rl.Close()

    for {
        line, err := rl.Readline()
        if err != nil { // io.EOF
            break
        }
		
		line = strings.TrimSpace(line)
        if line == "exit" {
            os.Exit(0)
        }
    }
}