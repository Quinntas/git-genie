package stringUtils

func ArrayToString(arr []string) string {
	var str string
	for _, v := range arr {
		str += v + " "
	}
	return str
}
