package core

import "fmt"

func FizzBuzz(from, to, mutiple1, mutiple2 uint, s1, s2 string) ([]string, error) {
	if from > to {
		return nil, fmt.Errorf("from can't be greater than to (%d > %d)", from, to)
	}

	resp := make([]string, to - from + 1)
	for i := 0; from <= to; from++ {
		tmp := ""
		if from % mutiple1 == 0 {
			tmp += s1
		}
		if from % mutiple2 == 0 {
			tmp += s2
		}
		if len(tmp) == 0 {
			tmp = fmt.Sprintf("%d", from)
		}
		resp[i] = tmp
		i++
	}
	return resp, nil
}
