package day04

import (
	"strings"
)

func Parse(input []byte) []map[string]string {
	var res []map[string]string

	for _, v := range strings.Split(string(input), "\n\n") {
		m := make(map[string]string)

		v = strings.Replace(v, "\n", " ", -1)
		v := strings.Split(v, " ")

		for _, vv := range v {
			parts := strings.Split(vv, ":")
			m[parts[0]] = parts[1]
		}

		res = append(res, m)
	}

	return res
}

func Valid(input []map[string]string, mandatory []string) int {
	count := 0
	for _, item := range input {
		remain := len(mandatory)
		for _, m := range mandatory {
			for k := range item {
				if k == m {
					remain--
				}
			}
		}

		if remain == 0 {
			count++
		}
	}

	return count
}
