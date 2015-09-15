package sortmap

import "sort"

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

func (this SortMaps) Sort(params map[string]interface{}) map[string]interface{} {
	for k, v := range params {
		this = append(this, sortMap{Key: k, Value: v})
	}
	sort.Sort(this)
	m := map[string]interface{}{}
	for _, v := range this {
		m[v.Key] = v.Value
	}
	return m
}
