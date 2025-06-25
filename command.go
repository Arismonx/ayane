package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add   string
	Del   int
	Edit  string
	Mark  int
	List  bool
	Help  bool
	Ayane bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo by specify task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a tody by index & specify task")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by indec to delete")
	flag.IntVar(&cf.Mark, "mark", -1, "Specify a todo by indec to check")
	flag.BoolVar(&cf.List, "list", false, "List all todo")
	flag.BoolVar(&cf.Help, "help", false, "All Command")
	flag.BoolVar(&cf.Ayane, "ayane", false, "home page")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(aya *Ayane) {
	switch {
	case cf.List:
		aya.print()
	case cf.Add != "":
		aya.add(cf.Add)
		aya.print()
	case cf.Edit != "":
		part := strings.SplitN(cf.Edit, ":", 2)
		if len(part) != 2 {
			fmt.Println("Error: invalid format for edit.")
			os.Exit(1)
		}

		index, err := strconv.Atoi(part[0])
		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}

		aya.edit(index, part[1])
		aya.print()
	case cf.Mark != -1:
		aya.check(cf.Mark)
		aya.print()
	case cf.Del != -1:
		aya.delete(cf.Del)
		aya.print()
	case cf.Help:
		AllCommand()
		HowtoUseCommand()
	case cf.Ayane:
		Display()
	default:
		fmt.Println("")
	}
}

func AllCommand() {
	fmt.Printf("All Command: \n")
	fmt.Printf("  %-10s %s\n", "-add", "Add a new todo by specify task")
	fmt.Printf("  %-10s %s\n", "-edit", "Edit a tody by index & specify task")
	fmt.Printf("  %-10s %s\n", "-del", "Specify a todo by indec to delete")
	fmt.Printf("  %-10s %s\n", "-list", "List all todo")
	fmt.Printf("  %-10s %s\n", "-mark", "Specify a todo by indec to mark")
	fmt.Printf("  %-10s %s\n", "-help", "All Command")
	fmt.Printf("  %-10s %s\n", "-ayane", "home page")
	fmt.Println()
}

func HowtoUseCommand() {
	fmt.Printf("How to use command: \n")
	fmt.Printf("  %-10s %s\n", "add", `-add "Your Task"`)
	fmt.Printf("  %-10s %s\n", "edit", `-edit id:"Edit Task" (-edit 1:"Go Run")`)
	fmt.Printf("  %-10s %s\n", "del", "-del id (-del 1)")
	fmt.Printf("  %-10s %s\n", "list", "-list")
	fmt.Printf("  %-10s %s\n", "mark", "-mark id (-mark 1)")
	fmt.Printf("  %-10s %s\n", "help", "-help")
	fmt.Printf("  %-10s %s\n", "ayane", "-ayane")
}
