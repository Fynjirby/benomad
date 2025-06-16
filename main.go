package main

import (
	"fmt"
	"os"

	"github.com/Slaves-Corp/benomad/core"
)

func main() {
	var needMoreArgs string = "Provide more arguments. See help with 'benomad help'"

	if len(os.Args) < 2 {
		fmt.Println(needMoreArgs)
		return
	}

	var helpMsg = `
  Welcome to benomad!

  Benomad is a bash script manager written on Go
  Visit our GitHub repo to see full command list, aliases, guidelines, etc
  github.com/Slaves-Corp/benomad
  
  Help:
   install <ben> - install some benomad metadata files (bens)
   remove <ben> - remove benomad 
   list - list all bens installed
   run <ben> - run a ben's script
   info <ben> - see information about any ben
   edit <ben> - edit ben file or script of ben
   
  Thanks for using benomad!`

	do := os.Args[1]
	switch do {
	default:
		fmt.Println("Command not found. See help with 'benomad help'")
	case "install", "add":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}

		err := core.InstallBen(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "list", "ls":
		err := core.ListBen()
		if err != nil {
			fmt.Println(err)
			return
		}
	case "remove", "delete":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}

		err := core.RemoveBen(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "info":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}

		err := core.InfoBen(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "edit":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}

		err := core.EditBen(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "run", "exec":
		if len(os.Args) < 3 {
			fmt.Println(needMoreArgs)
			return
		}
		err := core.RunBen(os.Args[2])
		if err != nil {
			fmt.Println(err)
			return
		}
	case "help", "man":
		fmt.Println(helpMsg)
	}
}
