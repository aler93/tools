package helpers

import (
	"fmt"
	"log"
	"strings"
)

// Print: Alias to fmt.Println
func Print(msg ...interface{}) {
	for _, m := range msg {
		_, e := fmt.Println(m)

		Check(e, "Error at Print(msg ...interface{})")
	}
}

// PrintType: Print var type in the terminal
func PrintType(t ...interface{}) {
	for _, m := range t {
		_, e := fmt.Printf("%T\n", m)

		Check(e, "Error at Print(msg ...interface{})")
	}
}

func GetType(msg interface{}) string {
	return fmt.Sprintf("%T", msg)
}

// Echo: Alias to fmt.Print
func Echo(msg ...interface{}) {
	for _, m := range msg {
		_, e := fmt.Print(m)

		Check(e, "Error at Echo(msg ...interface{})")
	}
}

// Echo: Alias to fmt.Print, breaks new line in the end
func EchoBreak(msg ...interface{}) {
	for _, m := range msg {
		_, e := fmt.Print(m)

		Check(e, "Error at Echo(msg ...interface{})")
	}
	fmt.Print("\n")
}

var logPath string

// LogPath: Set path to save logs
func LogPath(path string) string {
	if !EmptyString(path) {
		if path[len(path)-1:] != "/" {
			path += "/"
		}

		logPath = path
	}

	return logPath
}

// Check: Looks for error and prints a custom message; ignores if error == nil
func Check(e error, msg string) {
	if e != nil {
		time := NowString()
		err := e.Error()
		Print(msg, err)
		if !EmptyString(logPath) {
			AppendToFile(logPath+"Check.log", []string{time, msg, err})
		}

		log.Fatal(e)
	}
}

// Panic: Alias to panic; ignores if error == nil
func Panic(e error) {
	if e != nil {
		if !EmptyString(logPath) {
			time := NowString()
			AppendToFile(logPath+"Panic.log", []string{time, e.Error()})
		}

		panic(e)
	}
}

// StayCalm: Don't panic, if has error, just print error.Error() and return true
func StayCalm(e error) bool {
	if e != nil {
		if !EmptyString(logPath) {
			time := NowString()
			AppendToFile(logPath+"Check.log", []string{time, e.Error()})
		}
		Print(e.Error())

		return true
	}

	return false
}

// EmptyString: if empty true; else false
func EmptyString(s string) bool {
	if len(s) <= 0 {
		return true
	}

	return false
}

// Upper: Alias to strings.ToUpper
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower: Alias to strings.ToLower
func Lower(s string) string {
	return strings.ToLower(s)
}

// Title: Alias to strings.Title
func Title(s string) string {
	return strings.Title(s)
}
