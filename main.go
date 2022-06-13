package main

import (
	"flag"
	"io"
	"os"
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
	choseOperation := flag.String("operation", "", "chose operation")
	inputBody := flag.String("item", "", "body for add")
	choseFile := flag.String("fileName", "", "chose file")
	flag.Parse()
	mp := Arguments{}
	mp["operation"] = *choseOperation
	mp["item"] = *inputBody
	mp["fileName"] = *choseFile
	return mp
}

func GetInfo() {
}

func AddNewItem() {
}

func RemoveUser() {
}

func FindByID() {
}
