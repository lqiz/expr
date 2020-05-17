package expr

import (
	"strconv"
	"strings"
)

func Compare(v1, v2 string) int {
	str1List := strings.Split(v1, ".")
	str2List := strings.Split(v2, ".")
	len1 := len(str1List)
	len2 := len(str2List)
	max := len1
	if len1 < len2 {
		max = len2
	}

	for i := 0; i < max; i++ {
		v1 := uint64(0)
		if i < len1 {
			var err error
			v1, err = strconv.ParseUint(str1List[i], 10, 64)
			if err != nil {
				continue
			}
		}

		v2 := uint64(0)
		if i < len2 {
			var err error
			v2, err = strconv.ParseUint(str2List[i], 10, 64)
			if err != nil {
				continue
			}
		}

		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		}
	}
	return 0
}

// we only compare the number version
func VersionCompare(v1, operator, v2 string) bool {
	com := Compare(v1, v2)
	switch operator {
	case "==", "=":
		if com == 0 {
			return true
		}
	case "<":
		if com == 2 {
			return true
		}
	case ">":
		if com == 1 {
			return true
		}
	case "<=":
		if com == 0 || com == 2 {
			return true
		}
	case ">=":
		if com == 0 || com == 1 {
			return true
		}
	}
	return false
}
