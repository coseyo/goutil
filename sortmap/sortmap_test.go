package sortmap

import (
	"fmt"
	"testing"
)

func TestSortMaps(t *testing.T) {
	params := map[string]interface{}{
		"aa": 11,
		"cc": 22,
		"dd": "asdf",
		"bb": "dfe",
		"ff": map[string]interface{}{
			"ff3": 1,
			"ff2": 2,
			"ff4": 3,
			"ff1": map[string]interface{}{
				"ff1_a": 1,
				"ff1_e": 2,
				"ff1_d": 3,
				"ff1_b": 4,
			},
		},
		"ee": 2231,
	}
	var m SortMaps
	p := m.Sort(params)
	fmt.Println(p)

	md5 := MapToMD5(params)
	fmt.Println(md5)
}
