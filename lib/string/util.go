package string

import (
	"strings"

	"github.com/remotehy/go-std-mod/internal/log"
)

func Duplicate(ori string, num int, sep string) string {
	log.StdLog("Duplicate ori string [%s] %d times", ori, num)
	var tmp []string
	for i := 0; i < num; i++ {
		tmp = append(tmp, ori)
	}
	return strings.Join(tmp, sep)
}
