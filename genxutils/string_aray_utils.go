package genxutils

//StringArrayUtils string array utils
type StringArrayUtils struct {
	Data []string
}

//IndexOf return index of value in array
func (utils *StringArrayUtils) IndexOf(param string) (int, bool) {
	for i, s := range utils.Data {
		if s == param {
			return i, true
		}
	}
	return -1, false
}

//Unique returns only unique values from string array
func (utils *StringArrayUtils) Unique() []string {
	keys := make(map[string]bool)
	responce := []string{}
	for _, str := range utils.Data {
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
