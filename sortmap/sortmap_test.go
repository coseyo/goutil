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
	}
	var m SortMaps
	p := m.Sort(params)
	fmt.Println(p)
}
