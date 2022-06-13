package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	operationMissingErr = errors.New("-operation flag has to be specified")
	fileNameMissingErr  = errors.New("-fileName flag has to be specified")
	idMissingErr        = errors.New("-id flag has to be specified")
)

type Arguments map[string]string

func Perform(args Arguments, writer io.Writer) error {
	if args["operation"] == "add" {
		err := AddNewItem(args, writer)
		if err != nil {
			return err
		}
	} else if args["operation"] == "list" {
		err := GetInfo(args, writer)
		if err != nil {
			return err
		}
	} else if args["operation"] == "findById" {
		err := FindByID(args, writer)
		if err != nil {
			return err
		}
	} else if args["operation"] == "remove" {
		err := RemoveUser(args, writer)
		if err != nil {
			return err
		}
	} else if len(args["operation"]) == 0 {
		return operationMissingErr
	} else {
		return fmt.Errorf("Operation %s not allowed!", args["operation"])
	}

	return nil
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

func parseArgs() Arguments {
	_ = os.Args[1:]
	id := flag.String("id", "", "chose user id")
	inputBody := flag.String("item", "", "body for add")
	choseOperation := flag.String("operation", "", "chose operation")
	choseFile := flag.String("fileName", "", "chose file")
	flag.Parse()
	mp := Arguments{}
	mp["id"] = *id
	mp["operation"] = *choseOperation
	mp["item"] = *inputBody
	mp["fileName"] = *choseFile
	return mp
}

func GetInfo(args Arguments, writer io.Writer) error {
	fileName := args["fileName"]
	if len(fileName) == 0 {
		return fileNameMissingErr
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}
	data = append(data, '\n')
	writer.Write(data)
	return nil
}

func AddNewItem(args Arguments, writer io.Writer) error {
	return nil
}

func RemoveUser(args Arguments, writer io.Writer) error {
	return nil
}

type User struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func FindByID(args Arguments, writer io.Writer) error {
	input := []User{}
	id := args["id"]
	fileName := args["fileName"]
	if len(id) == 0 {
		return idMissingErr
	}
	if len(fileName) == 0 {
		return fileNameMissingErr
	}
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &input)
	if err != nil {
		return err
	}
	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	ind := 0
	check := false
	for index, i := range input {
		if userId == i.Id {
			ind = index
			check = true
		}
	}
	if !check {
		writer.Write([]byte(""))
		return nil
	}
	out, err := json.Marshal(input[ind])
	if err != nil {
		return err
	}
	out = append(out, '\n')
	writer.Write(out)
	return nil
}
