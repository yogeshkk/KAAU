package logger

import "fmt"

var (
	code, errorMessage string
)

// LogOut this is sjdlksd
func LogOut(out ...string) {

	fmt.Println(out)
}

// LogOutCritical Exported
func LogOutCritical(out ...string) {

	fmt.Printf("ERROR:  %v\n", out)
}

// LogOutInfo Exported
func LogOutInfo(out ...string) {

	fmt.Printf("INFO:  %v \n", out)
}
