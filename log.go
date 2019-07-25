package seed

import "log"

//Log is a wrapper around log.Println
func Log(v ...interface{}) {
	log.Println(v...)
}
