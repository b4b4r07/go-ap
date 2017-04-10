package ap

import (
	"bufio"
	"errors"
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

const mc = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport -s"

type WiFi interface {
}

type WifiScanner struct {
	Command string
}

// Alias of WifiScanner for mac
type MacWifiScanner struct {
	WifiScanner
}

type AccessPoint struct {
	ssid, bssid, quality, security string
}

type AccessPoints []AccessPoint

// Render wil return json body
func (ap *AccessPoints) Render() {
}

func NewMacWifiScanner() *MacWifiScanner {
	return &MacWifiScanner{WifiScanner{
		Command: mc,
	}}
}

func (w *WifiScanner) Scan() ([]string, error) {
	var lines []string
	args := strings.Split(w.Command, " ")
	if len(args) == 0 {
		return lines, errors.New("command not found")
	}
	command := args[0]
	options := args[1:]
	_, err := exec.LookPath(command)
	if err != nil {
		return lines, err
	}
	out, err := exec.Command(command, options...).Output()
	if err != nil {
		return lines, err
	}
	scanner := bufio.NewScanner(strings.NewReader(string(out)))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

// TODO
var (
	reSSID = regexp.MustCompile(`^ *(.*?) .*$`)
)

func Parse(lines []string) {
	for _, line := range lines {
		// result := reSSID.FindAllStringSubmatch(line, -1)
		fmt.Printf("%#v\n", line)
	}
}
