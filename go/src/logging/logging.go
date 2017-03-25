// We add file / line to the log.

package logging

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"strings"
)

const (
	COLOR_RED    = "\033[0;31m"
	COLOR_YELLOW = "\033[0;33m"
	COLOR_GREEN  = "\033[0;32m"
	COLOR_NONE   = "\033[0m"
)

var vFlag = flag.Int("v", 1, "")

func Fatal(v ...interface{}) {
	log.Fatal(fileLinePrefix(), fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) {
	log.Fatal(fileLinePrefix(), fmt.Sprintf(format, v...))
}

func Fatalln(v ...interface{}) {
	log.Fatalln(fileLinePrefix(), fmt.Sprint(v...))
}

func Panic(v ...interface{}) {
	log.Panic(fileLinePrefix(), fmt.Sprint(v...))
}

func Panicf(format string, v ...interface{}) {
	log.Panic(fileLinePrefix(), fmt.Sprintf(format, v...))
}

func Panicln(v ...interface{}) {
	log.Panicln(fileLinePrefix(), fmt.Sprint(v...))
}

func Print(v ...interface{}) {
	log.Print(fileLinePrefix(), fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	log.Print(fileLinePrefix(), fmt.Sprintf(format, v...))
}

func Println(v ...interface{}) {
	log.Println(fileLinePrefix(), fmt.Sprint(v...))
}

func GetVerboseLevel() int {
	return *vFlag
}

func SetVerboseLevel(level int) {
	*vFlag = level
}

func Vlog(level int, v ...interface{}) {
	if level <= *vFlag {
		log.Print(fileLinePrefix(), vlogPrefix(level), fmt.Sprint(v...), vlogSuffix(level))
	}
}

func Vlogf(level int, format string, v ...interface{}) {
	if level <= *vFlag {
		log.Print(fileLinePrefix(), vlogPrefix(level), fmt.Sprintf(format, v...), vlogSuffix(level))
	}
}

func vlogPrefix(level int) string {
	if level < 0 {
		return COLOR_RED
	} else if level == 0 {
		return COLOR_YELLOW
	} else if level >= 3 {
		return COLOR_GREEN
	}
	return ""
}

func vlogSuffix(level int) string {
	if level <= 0 || level >= 3 {
		return COLOR_NONE
	}
	return ""
}

func fileLinePrefix() string {
	// Find the first caller outside of package logging.
	for i := 2; ; i++ {
		if _, file, line, ok := runtime.Caller(i); ok {
			file = srcFilePath(file)
			if strings.Contains(file, "/logging/") {
				continue
			}
			return fmt.Sprintf("[%s:%d] ", file, line)
		}
		break
	}
	return "[unknown] "
}

// Returns the source file name after "go/src/".
func srcFilePath(path string) string {
	const START_AFTER = "go/src/"
	if index := strings.LastIndex(path, START_AFTER); index != -1 {
		return path[index+len(START_AFTER):]
	}
	return path
}
