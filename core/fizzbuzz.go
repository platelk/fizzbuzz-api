package core

import "fmt"

func FizzBuzz(from, to, multiple1, multiple2 int, s1, s2 string) ([]string, error) {
	if from > to {
		return nil, fmt.Errorf("from can't be greater than to (%d > %d)", from, to)
	}
	if multiple1 < 1 || multiple2 < 1 {
		return nil, fmt.Errorf("multiple can't be inferior to 1")
	}

	resp := make([]string, to - from + 1)
	for i := 0; from <= to; from++ {
		tmp := ""
		if from % multiple1 == 0 {
			tmp += s1
		}
		if from % multiple2 == 0 {
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
