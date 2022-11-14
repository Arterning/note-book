package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {

	log.SetPrefix("黄宁是天才--")
	//log.SetOutput(os.Stdout)
	f, err := os.OpenFile("./yx.log", os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	Trace = log.New(ioutil.Discard,
		"Trace:",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(os.Stdout,
		"Info:",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(os.Stdout,
		"Warning:",
		log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(io.MultiWriter(f, os.Stderr),
		"Error:",
		log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {

	// log.Println("Log")

	// log.Fatal("Fatal")

	// log.Panic("panic!!")

	Trace.Println("Trace")
	Info.Println("Info")
	Warning.Println("Warning")
	Error.Println("Error")

}
