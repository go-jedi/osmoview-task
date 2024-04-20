package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-jedi/osmoview-task/internal/app"
	"github.com/go-jedi/osmoview-task/pkg/changefield"
	"github.com/go-jedi/osmoview-task/pkg/logger"
)

type Data1 struct {
	Name string
	Type string
}

type Data2 struct {
	ID  string
	Arr []int
	Map map[string]interface{}
	Any interface{}
}

func main() {
	ctx := context.Background()

	err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to init app: %s", err.Error())
	}

	data1 := Data1{Name: "Name", Type: "Type"}
	data2 := Data2{ID: "1", Arr: []int{1, 2, 3}, Map: map[string]interface{}{"test": "test"}, Any: "any"}

	err = changefield.ChangeField(&data1, "Name", "NewName")
	if err != nil {
		logger.Error("Ошибка изменения поля data1")
		return
	}

	err = changefield.ChangeField(&data2, "ID", "2")
	if err != nil {
		logger.Error("Ошибка изменения поля data2")
	}

	fmt.Println("data1:", data1)
	fmt.Println("data2:", data2)
}
