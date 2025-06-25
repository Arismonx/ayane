package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// Struct for todo
type Todo struct {
	Task       string
	Check      bool
	CreatedAt  string
	CompleteAt string
}

// Slice of todo name AYANEEEEEEEEEEEEEEEEEE! >w<
type Ayane []Todo

// Function for "Add" new task
func (ayane *Ayane) add(task string) {
	currentTime := time.Now()
	formatTtime := currentTime.Format("Mon, 02 Jan 2006, 3:04PM")

	todo := Todo{
		Task:       task,
		Check:      false,
		CreatedAt:  formatTtime,
		CompleteAt: "",
	}

	*ayane = append(*ayane, todo)
	fmt.Println("> Create Task Complete!")
}

// Function "validate" index for check befor actuate
func (ayane *Ayane) validate(index int) error {
	if index < 0 || index >= len(*ayane) {
		err := errors.New("error: invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

// Function for "Delete" by use index of task in ayane
func (ayane *Ayane) delete(index int) error {
	aya := *ayane

	if err := aya.validate(index); err != nil {
		return err
	}
	*ayane = append(aya[:index], aya[index+1:]...)
	fmt.Printf("> Delete Complete!\n")
	return nil
}

// Function for "Edit" task by use index and newTask
func (ayane *Ayane) edit(index int, newTask string) error {
	aya := *ayane

	if err := aya.validate(index); err != nil {
		return err
	}

	aya[index].Task = newTask

	return nil
}

// Function for "check" task by use index
func (ayane *Ayane) check(index int) error {
	aya := *ayane

	if err := aya.validate(index); err != nil {
		return err
	}

	currentTime := time.Now()
	formatTtime := currentTime.Format("Mon, 02 Jan 2006, 3:04PM")
	aya[index].Check = true
	aya[index].CompleteAt = formatTtime
	fmt.Printf("> Mark Complete!\n")
	return nil
}

// Function for "print" output to table from package "github.com/aquasecurity/table"
func (ayane *Ayane) print() {
	t := table.New(os.Stdout)
	t.SetHeaders("ID", "Task", "Check", "CreatedAt", "CompleteAt")
	t.SetHeaderStyle(table.StyleBold)
	t.SetLineStyle(table.StyleBlue)

	for i, value := range *ayane {
		checklist := "[✗]"

		if value.Check {
			checklist = "[✓]"
		}

		t.AddRow(
			strconv.Itoa(i),
			value.Task,
			checklist,
			value.CreatedAt,
			value.CompleteAt,
		)
	}
	t.Render()
}
