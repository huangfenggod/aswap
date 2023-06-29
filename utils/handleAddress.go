package utils

func HandleAddress(address string)string  {
	if len(address)<8 {
		return address
	}
	return substr(address,0,4)+"***" + substr(address,len(address)-4,len(address))
}

func substr(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}
