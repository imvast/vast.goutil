/**
 * Author: imvast
 * File: init.go
 */

package goutil

import (
	"fmt"
	"github.com/fatih/color"
)

func init() {
	str := fmt.Sprintf(color.HiWhiteString("[goutil]") + color.HiBlackString(" initialized | made by vast"))
	fmt.Println(getTime(true), str)
}
