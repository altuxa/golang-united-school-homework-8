package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	operationMissingErr = errors.New("-operation flag has to be specified")
	fileNameMissingErr  = errors.New("-fileName flag has to be specified")
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
	file, err := os.Open(fileName)
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

func FindByID(args Arguments, writer io.Writer) error {
	return nil
}
