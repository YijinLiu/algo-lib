// We add file / line to the log.

package logging

import (
	"fmt"
	"log"
	"path"
	"runtime"
)

func getFileLinePrefix() string {
	if _, file, line, ok := runtime.Caller(2); ok {
		return fmt.Sprintf("[%s:%d] ", path.Base(file), line)
	}
	return ""
}

func Fatal(v ...interface{}) {
	log.Fatal(getFileLinePrefix(), fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) {
	log.Fatal(getFileLinePrefix(), fmt.Sprintf(format, v...))
}

func Fatalln(v ...interface{}) {
	log.Fatalln(getFileLinePrefix(), fmt.Sprint(v...))
}

func Panic(v ...interface{}) {
	log.Panic(getFileLinePrefix(), fmt.Sprint(v...))
}

func Panicf(format string, v ...interface{}) {
	log.Panic(getFileLinePrefix(), fmt.Sprintf(format, v...))
}

func Panicln(v ...interface{}) {
	log.Panicln(getFileLinePrefix(), fmt.Sprint(v...))
}

func Print(v ...interface{}) {
	log.Print(getFileLinePrefix(), fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	log.Print(getFileLinePrefix(), fmt.Sprintf(format, v...))
}

func Println(v ...interface{}) {
	log.Println(getFileLinePrefix(), fmt.Sprint(v...))
}
