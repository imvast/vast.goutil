/**
 * Author: imvast
 * File: logger_test.go
 */

package goutil

import (
    "fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(getTime(true), "printout:test?")
}