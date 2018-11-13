package main

import (
	"log"
	"os"

	"github.com/Jacobious52/telegram-exporter/pkg/exporter"

	"gopkg.in/alecthomas/kingpin.v2"
)

var filePath = kingpin.Arg("filepath", "filepath to the result json to import").Required().ExistingFile()
var chatName = kingpin.Flag("chat", "chat name to export").Short('c').Required().String()
var escapeText = kingpin.Flag("escape", "escape the text of the csv to url query safe?").Short('e').Default("false").Bool()
var b64Text = kingpin.Flag("b64", "escape the text of the csv to url query safe?").Short('b').Default("false").Bool()

func main() {
	kingpin.Parse()

	inFile, err := os.Open(*filePath)
	if err != nil {
		log.Fatalln("failed to open inFile:", err)
	}
	defer inFile.Close()

	result, err := exporter.DecodeJsonResult(inFile)
	if err != nil {
		log.Fatalln("failed to decode result:", err)
	}

	chat := result.FindChat(*chatName)
	if chat == nil {
		log.Fatalln("could not find chat by name:", *chatName)
	}

	err = exporter.ExportResult(*chat, os.Stdout, *escapeText, *b64Text)
	if err != nil {
		log.Fatalln("failed to output csv:", err)
	}
}
