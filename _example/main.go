package main

import (
	"github.com/b4b4r07/go-ap"
	// "github.com/k0kubun/pp"
)

func main() {
	wifi := ap.NewMacWifiScanner()
	lines, err := wifi.Scan()
	if err != nil {
		panic(err)
	}
	// pp.Println(lines)
	// for _, line := range lines {
	// 	log.Println(line)
	// }
	ap.Parse(lines)
}
