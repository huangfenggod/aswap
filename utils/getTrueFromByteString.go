package utils

//排除字符串前面的0
func GetTrueString(s string)string {

	for k,v :=range s{
		if v != '0' {
			return s[k:]
		}
	}
	return ""
}

