package types

import "container/list"

//NewList ...
func NewList(_data *list.List) *List {
	return &List{
		data: _data,
	}
}

//List ...
type List struct {
	data *list.List
}

//StringArray ...
func (lst *List) StringArray() []string {
	response := make([]string, 0)
	for val := lst.data.Front(); val != nil; val = val.Next() {
		strValue, success := val.Value.(string)
		if !success {
			return []string{}
		}
		response = append(response, strValue)
	}
	return response
}
