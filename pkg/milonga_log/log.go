package milonga_log

import (
	"fmt"
	"log"
)

const txtRed = "\033[31m"
const txtGrenn = "\033[32m"
const txtYellow = "\033[33m"
const txtReset = "\033[0m"
const txtBold = "\033[1m"

var LogPrefix = "TANGO"
var msgPrefixSeparator = " > "

const msgEnd = "\n"

func getMsg(msg string) string {
	return fmt.Sprintf("%s %v %s", txtReset, msg, msgEnd)

}

func Print(msg string) {
	log.Printf("%s%s%s %s", txtBold, LogPrefix, msgPrefixSeparator, getMsg(msg))

}

func PrintError(msg string) {
	log.Printf("%s%s [error] %s %s", txtRed, LogPrefix, msgPrefixSeparator, getMsg(msg))

}

func PrintOk(msg string) {
	log.Printf("%s%s%s%s %s%s", txtBold, txtGrenn, LogPrefix, msgPrefixSeparator, txtGrenn, getMsg(msg))

}

func PrintWarning(msg string) {
	log.Printf("%s%s%s%s %s", txtBold, txtYellow, LogPrefix, msgPrefixSeparator, getMsg(msg))

}
