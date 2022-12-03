package goutil

import (
    "fmt"
    "time"
	"os"
	"os/exec"
	"bufio"
	"syscall"
	"unsafe"
    "github.com/fatih/color"
)

func getTime() string {
    currentTime := time.Now().Format("15:04:05")
    return color.BlueString(currentTime)
}

// func PrettyPrint(col any, content ...interface{}) {
// 	if content == nil {
// 		fmt.Printf("[%s] %s", getTime(), color.CyanString(fmt.Sprintln(col)))
// 	} else {
// 		if col == 1 {
// 			fmt.Printf("[%s] %s", getTime(), color.BlueString(fmt.Sprintln(content...)))
// 		} else { // TODO: handle extra color values
// 			fmt.Printf("%s", col) 
// 		}
// 	}
// }

func SetTitle(title string) {
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		return
	}
	defer syscall.FreeLibrary(handle)

	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		return
	}

	syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
}

var (
	Username = "..."
	Password = "..."
)
func TitleThread() {
	for {
		SetTitle(fmt.Sprintf("Made by @github.com/imvast - Logged in as: %s | with id: %s", Username, Password))
		time.Sleep(50 * time.Millisecond)
	}
}

func Print(content...interface{}) {
	fmt.Printf("[%s] %s %s", getTime(), color.HiBlackString("➜"), color.CyanString(fmt.Sprintln(content...)))
}

func Logger(status int, content ...interface{}) {
	if status == 0 {
		fmt.Printf("[%s] %s %s", getTime(), color.GreenString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else if status == 1 {
		fmt.Printf("[%s] %s %s", getTime(), color.YellowString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else if status == 2 {
		fmt.Printf("[%s] %s %s", getTime(), color.RedString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else {
		return
	}
}

func Question(content string) string {
    fmt.Printf("[ %s ] %s", color.BlueString("?"), color.BlueString(fmt.Sprintf(content)))
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	question := scanner.Text()
    err := scanner.Err()
	if err != nil {
		Logger(2, err)
	}
	return question
}

func ClearConsole() {
	cmd := exec.Command("cmd", "/C", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func UpdateConsoleTitle(title string) {
	cmd := exec.Command("cmd", "/C", "title", title)
	err := cmd.Run()
	if err != nil {
		Logger(2, err.Error())
	}
}

func HandleError(err error) {
	// if DEBUG is true
	Logger(2, err.Error())
}

func GoodExit(exitCode int) {
	time.Sleep(3 * time.Second)
	os.Exit(exitCode)
}
