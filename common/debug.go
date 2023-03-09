package common

import (
	"fmt"
	"time"
)

func Print(label string, value string) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"

	fmt.Println(string(colorReset), ``)
	fmt.Println(string(colorGreen), `=========================================`)
	fmt.Println(string(colorGreen), time.Now().Format("YYYY-MM-DD hh:mm:ss")+" : "+label)
	fmt.Println(string(colorGreen), value)
	fmt.Println(string(colorGreen), `=========================================`)
	fmt.Println(string(colorReset), ``)
}
