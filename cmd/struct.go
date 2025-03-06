package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/alexeyco/simpletable"
)

type Todo struct {
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func validate(id, cap int) error {
	if id >= cap || id < 0 {
		return errors.New("index out of range")
	}

	return nil

}

func toggle(id int, state bool) error {
	t := Todos{}
	load(&t)

	if err := validate(id, len(t)); err != nil {
		return err
	}
	if state {
		completedAt := time.Now()
		t[id].CompletedAt = &completedAt
	} else {
		t[id].CompletedAt = nil
	}

	t[id].Completed = state
	save(t)
	return nil
}

func add(title string) {
	t := Todos{}
	load(&t)

	newTask := Todo{
		Description: title,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}

	t = append(t, newTask)

	save(t)
}

func delete(id int) error {
	t := Todos{}
	load(&t)
	if err := validate(id, len(t)); err != nil {
		return err
	}

	t = append(t[:id], t[id+1:]...)
	save(t)
	return nil
}

func print() {
	t := Todos{}
	load(&t)

	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: "#"},
			{Align: simpletable.AlignLeft, Text: "Task"},
			{Align: simpletable.AlignLeft, Text: "Completed"},
			{Align: simpletable.AlignLeft, Text: "CreatedAt"},
			{Align: simpletable.AlignLeft, Text: "CompletedAt"},
		},
	}

	for i := range t {
		title := t[i].Description
		id := strconv.Itoa(i)

		completed := "NO"
		completedAt := ""
		if t[i].Completed {
			completed = "YES"
			if t[i].CompletedAt != nil {
				completedAt = (*t[i].CompletedAt).Format(time.RFC1123)
			}
		}
		createdAt := t[i].CreatedAt.Format(time.RFC1123)

		table.Body.Cells = append(table.Body.Cells, []*simpletable.Cell{
			{Align: simpletable.AlignLeft, Text: id},
			{Align: simpletable.AlignLeft, Text: title},
			{Align: simpletable.AlignLeft, Text: completed},
			{Align: simpletable.AlignLeft, Text: createdAt},
			{Align: simpletable.AlignLeft, Text: completedAt},
		})

	}
	table.SetStyle(simpletable.StyleCompactClassic)
	fmt.Println(table.String())

}
