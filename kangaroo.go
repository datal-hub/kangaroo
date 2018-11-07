package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type KangarooTask struct {
	x1 int
	x2 int
	v1 int
	v2 int
}

func validParam(param int) bool {
	if param > 10000 || param < -10000 {
		return false
	}
	return true
}

func (t KangarooTask) isValid() bool {
	if validParam(t.x1) && validParam(t.x2) && validParam(t.v1) && validParam(t.v2) {
		return true
	}
	return false
}

func (t KangarooTask) ExistDecision() bool {
	if !t.isValid() {
		return false
	}
	distance := math.Abs(float64(t.x1 - t.x2))
	if distance == 0 && int(math.Abs(float64(t.v1-t.v2))) == 0 {
		return true
	}
	if int(distance)%int(math.Abs(float64(t.v1-t.v2))) != 0 {
		return false
	}
	if (t.x1 < t.x2 && t.v1 <= t.v2) || (t.x1 > t.x2 && t.v1 >= t.v2) {
		return false
	}
	return true
}

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	params := strings.Split(input, " ")
	intParams := make([]int, 4)
	var err error
	for i, s := range params {
		if intParams[i], err = strconv.Atoi(strings.TrimSuffix(s, "\n")); err != nil {
			fmt.Print("NO")
			return
		}
	}
	task := KangarooTask{x1: intParams[0], v1: intParams[1], x2: intParams[2], v2: intParams[3]}
	if !task.ExistDecision() {
		fmt.Print("NO")
		return
	}
	fmt.Print("YES")
	return
}
