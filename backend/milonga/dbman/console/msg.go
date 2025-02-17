package console

import "log"

const errorPrefix = "DBMAN > "

func Print(msg string) {
	log.Printf("%s %s\n", errorPrefix, msg)
}
