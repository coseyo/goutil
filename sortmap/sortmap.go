package sortmap

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"sort"
)

type sortMap struct {
	Key   string
	Value interface{}
}

type SortMaps []sortMap

func (this SortMaps) Len() int {
	return len(this)
}

func (this SortMaps) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this SortMaps) Less(i, j int) bool {
	return this[i].Key < this[j].Key
}

func (this SortMaps) Sort(params map[string]interface{}) SortMaps {
	for k, v := range params {
		vValue := reflect.ValueOf(v)

		if vValue.Kind() == reflect.Map {
			vValueI := vValue.Interface()
			ns := &SortMaps{}
			v = ns.Sort(vValueI.(map[string]interface{}))
		}

		this = append(this, sortMap{Key: k, Value: v})
	}
	sort.Sort(this)

	return this
}

func MapToMD5(params map[string]interface{}) []byte {

	smStruct := &SortMaps{}
	sm := smStruct.Sort(params)
	s := ""

	for _, v := range sm {
		s += fmt.Sprintf("%s=%v&", v.Key, v.Value)
	}

	ctMd5 := md5.New()
	ctMd5.Write([]byte(s))
	return ctMd5.Sum(nil)

}
