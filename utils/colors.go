package utils

import "fmt"

// Colors variables
var (
	Black       = Color("\033[0;30m%s\033[0m")
	Red         = Color("\033[0;31m%s\033[0m")
	Green       = Color("\033[0;32m%s\033[0m")
	Orange      = Color("\033[0;33m%s\033[0m")
	Blue        = Color("\033[0;34m%s\033[0m")
	Purple      = Color("\033[0;35m%s\033[0m")
	Cyan        = Color("\033[0;36m%s\033[0m")
	LightGray   = Color("\033[0;37m%s\033[0m")
	DarkGray    = Color("\033[1;30m%s\033[0m")
	LightRed    = Color("\033[1;31m%s\033[0m")
	LightGreen  = Color("\033[1;32m%s\033[0m")
	Yellow      = Color("\033[1;33m%s\033[0m")
	LightBlue   = Color("\033[1;34m%s\033[0m")
	LightPurple = Color("\033[1;35m%s\033[0m")
	LightCyan   = Color("\033[1;36m%s\033[0m")
	White       = Color("\033[1;37m%s\033[0m")
)

// Color return string in a color specified
func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

// func main() {
// 	fmt.Println(Red("hello, world!"))
// }
