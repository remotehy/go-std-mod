package math

import "github.com/remotehy/go-std-mod/internal/log"

func AddOne(ori int) int {
	log.StdLog("AddOne to %d", ori)
	return ori + 1
}
