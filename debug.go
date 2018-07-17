// +build debug

package gosftp

import "log"

func debug(fmt string, args ...interface{}) {
	log.Printf(fmt, args...)
}
