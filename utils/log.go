package utils

import (
	"fmt"
	//"os"
	"runtime"
	//"strconv"
)

/*
ANSI Colors for printf
*/
var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
)

/*
ANSI Colors for printf
*/
var (
	Info    = Teal
	Warn    = Yellow
	Fata    = Red
	Success = Green
)

/*
Common function to print logs throughout the service
Define prefix and/or postfix for consistent logs
Replace all fmt.Println with this utility

works only if debug=true
*/
func Println(messages ...interface{}) {

	//Default debug false
	//var debug bool

	//debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

	if true {
		prefix := "[=>]"
		postfix := "\n"
		for _, message := range messages {
			fmt.Println(prefix, message, postfix)
		}
	}
}

/*
Common function to print logs throughout the service
Define prefix and/or postfix for consistent logs
Replace all fmt.Println with this utility

works only if debug=true
*/
func PrintSuccess(messages ...interface{}) {
	Println(Success(messages))
}

/*
Common function to print logs throughout the service
Define prefix and/or postfix for consistent logs
Replace all fmt.Println with this utility

works only if debug=true
*/
func PrintWarn(messages ...interface{}) {
	Println(Warn(messages))
}

/*
Common function to print logs throughout the service
Define prefix and/or postfix for consistent logs
Replace all fmt.Println with this utility

works only if debug=true
*/
func PrintInfo(messages ...interface{}) {
	Println(Info(messages))
}

/*
Common function to print logs throughout the service
Define prefix and/or postfix for consistent logs
Replace all fmt.Println with this utility

works only if debug=true
*/
func PrintFatal(messages ...interface{}) {
	Println(Fata(messages))
}

/*
Color is
*/
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

//this logs the function name as well.
func HandleError(err error, params ...interface{}) {
	// notice that we're using 1, so it will actually log the where
	// the error happened, 0 = this function, we don't want that.
	PrintFatal("----------------------")
	pc, fn, line, _ := runtime.Caller(1)

	PrintFatal(fmt.Sprintf("Error in %s[%s:%d]", runtime.FuncForPC(pc).Name(), fn, line))
	PrintFatal("[Error] ", err)
	for _, param := range params {
		PrintFatal(param)
	}
	PrintFatal("----------------------")
}
