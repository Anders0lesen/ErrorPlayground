package logging

import "log"

func Setup() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
