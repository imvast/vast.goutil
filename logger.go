/**
 * Author: imvast
 * File: logger.go
 */

package goutil

import (
	"fmt"
	"time"
	"os"
	"os/exec"
	"bufio"
	"strconv"
	"math/rand"
	// "syscall"
	// "unsafe"
	"github.com/fatih/color"
)


func getTime(c bool) string {
	currentTime := time.Now().Format("15:04:05")
	if c == true {
		return color.HiBlackString(currentTime)
	} else {
		return currentTime
	}
	
}

/**
* unfinished *
func PrettyPrint(col any, content ...interface{}) {
	if content == nil {
		fmt.Printf("[%s] %s", getTime(), color.CyanString(fmt.Sprintln(col)))
	} else {
		if col == 1 {
			fmt.Printf("[%s] %s", getTime(), color.BlueString(fmt.Sprintln(content...)))
		} else { // TODO: handle extra color values
			fmt.Printf("%s", col) 
		}
	}
}

* causes err on certain devices *
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


* unused due to errors *
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
*/

// experimental since i havent read the docs for it lol
// func ToString(content ...interface{}) string {
// 	msg, _ := strconv.Atoi(content...)
// 	return msg
// }
// func ToInt(content ...interface{}) int {
// 	return strconv.Iota(content...)
// }

func Print(content...interface{}) {
	fmt.Printf("[%s] %s %s", color.BlueString(getTime(false)), color.HiBlackString("➜"), color.CyanString(fmt.Sprintln(content...)))
}

func Logger(status int, content ...interface{}) {
	if status == 0 {
		fmt.Printf("[%s] %s %s", color.BlueString(getTime(false)), color.GreenString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else if status == 1 {
		fmt.Printf("[%s] %s %s", color.BlueString(getTime(false)), color.YellowString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else if status == 2 {
		fmt.Printf("[%s] %s %s", color.BlueString(getTime(false)), color.RedString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else {
		return
	}
}

func StatLog(col string, status int, content ...interface{}) {
	// idk how tf to do this :/
    // switch status.(type) {
	// 	case string:
	// 		return
	// 	case int32:
	// 		status := strconv.Itoa(status)
	// 	default:
	// 		return
	// }
	statc := strconv.Itoa(status)
	if col == "green" {
		fmt.Printf("%s [%v] %s %s", getTime(false), color.GreenString(statc), color.GreenString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else if col == "red" {
		fmt.Printf("%s [%v] %s %s", getTime(false), color.RedString(statc), color.RedString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else if col == "yellow" {
		fmt.Printf("%s [%v] %s %s", getTime(false), color.YellowString(statc), color.YellowString("➜"), color.CyanString(fmt.Sprintln(content...)))
		} else if col == "magent"{
		fmt.Printf("%s [%v] %s %s", getTime(false), color.MagentaString(statc), color.MagentaString("➜"), color.CyanString(fmt.Sprintln(content...)))
	} else {
		return
	}
}

func Question(content string) string {
    fmt.Printf("[ %s ] %s", color.BlueString("?"), color.BlueString(fmt.Sprintf(content)))
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
	answer := scanner.Text()
    err := scanner.Err()
	if err != nil {
		Logger(2, err)
	}
	return answer
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
func RandStr(n int) string {
    // make sure to include this: func init() { rand.Seed(time.Now().UnixNano()) }
    b := make([]rune, n)
    for i := range b {
        b[i] = letterRunes[rand.Intn(len(letterRunes))]
    }
    return string(b)
}


func ClearConsole() {
	cmd := exec.Command("cmd", "/C", "clear||cls")
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
