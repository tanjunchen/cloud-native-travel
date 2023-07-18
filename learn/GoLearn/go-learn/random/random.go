package random

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func Suffix() string {
	rand.Seed(time.Now().UnixNano())
	str := strconv.Itoa(rand.Intn(10000))
	return str
}

func AA() {
	nonblockingTaints := ""
	nonblockingTaintsMap := map[string]struct{}{}
	for _, t := range strings.Split(nonblockingTaints, ",") {
		if strings.TrimSpace(t) != "" {
			nonblockingTaintsMap[strings.TrimSpace(t)] = struct{}{}
		}
	}
	if len(nonblockingTaintsMap) > 0 {
		fmt.Print("aaa")
	} else {
		fmt.Println("bbbb")
	}
}
