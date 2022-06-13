package main

import (
	"errors"
	"flag"
	"io"
	"os"
)

var (
	operationMissingErr = errors.New("-operation flag has to be specified")
	fileNameMissingErr  = errors.New("-fileName flag has to be specified")
)

type Arguments map[string]string

func Perform(args Arguments, writer io.Writer) error {
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

func GetInfo() {
}

func AddNewItem(args Arguments) error {
}

func RemoveUser() {
}

func FindByID() {
}
