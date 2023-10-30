package main

import (
	"github.com/remotehy/go-std-mod/internal/log"
	"github.com/remotehy/go-std-mod/lib/math"
	string2 "github.com/remotehy/go-std-mod/lib/string"
)

func main() {
	log.StdLog("%s start", "prog-a")
	log.StdLog("AddOne to 100 is %d", math.AddOne(100))
	log.StdLog("res: %s", string2.Duplicate("apple", 3, ","))
	log.StdLog("%s done", "prog-a")
}
