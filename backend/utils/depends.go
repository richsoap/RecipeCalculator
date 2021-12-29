package utils

import (
	"strconv"
	"strings"

	"github.com/richsoap/RecipeCalculator/errors"
)

func StringDependsToMap(in string) (map[uint64]int, error) {
	pairs := strings.Split(in, ";")
	result := make(map[uint64]int)
	for _, pair := range pairs {
		kv := strings.Split(pair, ":")
		if len(kv) != 2 {
			return nil, errors.BROKEN_DATA
		}
		id, err := strconv.Atoi(strings.TrimSpace(kv[0]))
		if err != nil {
			return nil, err
		}
		amount, err := strconv.Atoi(strings.TrimSpace(kv[1]))
		if err != nil {
			return nil, err
		}
		result[uint64(id)] = amount
	}
	return result, nil
}

func MapDependsToString(in map[uint64]int64) string {
	var sb strings.Builder
	first := true
	for k, v := range in {
		if !first {
			sb.WriteByte(';')
		}
		first = false
		sb.WriteString(strconv.Itoa(int(k)))
		sb.WriteByte(':')
		sb.WriteString(strconv.Itoa(int(v)))
	}
	return sb.String()
}
