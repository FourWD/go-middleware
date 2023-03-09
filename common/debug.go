package common

import (
	"fmt"
	"time"
)

const colorReset string = "\033[0m"
const colorGreen string = "\033[32m"
const colorRed string = "\033[31m"

func Print(label string, value string) {
	fmt.Println(string(colorReset), ``)
	fmt.Println(string(colorGreen), `=========================================`)
	fmt.Println(string(colorGreen), time.Now().Format("2017-09-07 17:06:06")+" : "+label)
	fmt.Println(string(colorGreen), value)
	fmt.Println(string(colorGreen), `=========================================`)
	fmt.Println(string(colorReset), ``)
}

func PrintError(label string, value string) {
	fmt.Println(string(colorReset), ``)
	fmt.Println(string(colorRed), `=========================================`)
	fmt.Println(string(colorRed), time.Now().Format("2017-09-07 17:06:06")+" : "+label)
	fmt.Println(string(colorRed), value)
	fmt.Println(string(colorRed), `=========================================`)
	fmt.Println(string(colorReset), ``)
}
