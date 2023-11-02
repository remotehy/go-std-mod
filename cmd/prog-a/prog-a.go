package main

import (
	"github.com/remotehy/go-std-mod/internal/log"
	"github.com/remotehy/go-std-mod/lib/math"
	"github.com/remotehy/go-std-mod/lib/string"
)

func main() {
	log.StdLog("%s start", "prog-a")
	log.StdLog("AddOne to 100 is %d", mymath.AddOne(100))
	log.StdLog("res: %s", mystring.Duplicate("apple", 3, ","))
	log.StdLog("%s done", "prog-a")
}
