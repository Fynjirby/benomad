package main

import (
	"fmt"
	"os"

	"github.com/Fynjirby/benomad/core"
)

func main() {
	var needMoreArgs string = "Provide more arguments. See help with 'benomad help'"

	var helpMsg = `
  Welcome to benomad!

  Benomad is a bash script manager written on Go
  Visit our GitHub repo to see full command list, aliases, guidelines, etc
  github.com/Fynjirby/benomad
  
  Help:
   install <url> - install some script by url
   remove <script> - remove script
   list - list all scripts installed
   run <script> - run a script
   edit <script> - edit script
   
  Thanks for using benomad!`

	if len(os.Args) < 2 {
		fmt.Println(helpMsg)
		return
	}

	do := os.Args[1]
	switch do {
	default:
		fmt.Println("Command not found. See help with 'benomad help'")
	case "install", "add", "i":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}

		err := core.Install(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "list", "ls", "l":
		err := core.List()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "remove", "delete", "rm", "rem", "r":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}

		err := core.Remove(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "edit":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}

		err := core.Edit(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "run", "exec":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}

		var args []string
		if len(os.Args) > 3 {
			args = os.Args[3:]
		} else {
			args = []string{}
		}

		err := core.Run(os.Args[2], args)
		if err != nil {
			fmt.Println(err)
			return
		}
	case "help", "man", "h":
		fmt.Println(helpMsg)
	}
}
