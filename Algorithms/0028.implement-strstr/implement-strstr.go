package problem0028

func strStr(haystack string, needle string) int {
	hlen, nlen := len(haystack), len(needle)
	for i := 0; i <= hlen - nlen; i++ {
		if haystack[i:i+nlen] == needle {
			return i
		}
	}
	return -1
}