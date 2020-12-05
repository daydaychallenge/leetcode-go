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

const PrimeRK = 16777619

func HashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash * PrimeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func strStrRK(haystack string, needle string) int {
	hashSub, pow := HashStr(needle)
	n := len(needle)
	var h uint32
	if n > len(haystack) {
		return -1
	}
	for i := 0; i < n; i++ {
		h = h * PrimeRK + uint32(haystack[i])
	}

	if h == hashSub && haystack[:n] == needle {
		return 0
	}

	for i := n; i < len(haystack); {
		h *= PrimeRK
		h += uint32(haystack[i])
		h -= pow * uint32(haystack[i-n])
		i++
		if h == hashSub && haystack[i-n: i] == needle {
			return i-n
		}
	}
	return -1
}