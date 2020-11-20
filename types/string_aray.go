package types

import "container/list"

//StringArray string array sArray
type StringArray struct {
	Data []string
}

//IndexOf return index of value in array
func (sArray *StringArray) IndexOf(param string) (int, bool) {
	for i, s := range sArray.Data {
		if s == param {
			return i, true
		}
	}
	return -1, false
}

//Unique returns only unique values from string array
func (sArray *StringArray) Unique() []string {
	keys := make(map[string]bool)
	responce := []string{}
	for _, str := range sArray.Data {
		if str == "" {
			continue
		}
		if _, value := keys[str]; !value {
			keys[str] = true
			responce = append(responce, str)
		}
	}
	return responce
}

//List ...
func (sArray *StringArray) List() *list.List {
	cList := list.New()
	for _, s := range sArray.Data {
		cList.PushBack(s)
	}
	return cList
}
