package problem0028

func strStr(haystack string, needle string) int {
	hLen, nLen := len(haystack), len(needle)
	for i := 0; i <= hLen - nLen; i++ {
		if haystack[i:i+nLen] == needle {
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

func failTable(p string) []int {
	pLen := len(p)
	ret  := make([]int, pLen)
	ret[0] = -1
	for k, j := -1, 0; j < pLen - 1; {
		if k == -1 || p[k] == p[j] {
			k, j = k+1, j+1
			// 对"ababab"类字符串优化
			if p[k] != p[j] {
				ret[j] = k
			} else {
				ret[j] = ret[k]
			}
		} else {
			k = ret[k]
		}
	}
	return ret
}

func strStrKMP(haystack string, needle string) int {
	n := len(needle)
	if n == 0 {
		return 0
	}
	if n > len(haystack) {
		return -1
	}
	failJumps := failTable(needle)
	i, j := 0, 0
	for ; i < len(haystack) && j < n; {
		if j == -1 || haystack[i] == needle[j] {
			i, j = i+1, j+1
		} else {
			j = failJumps[j]
		}
	}
	if j == n {
		return i - j
	}
	return -1
}

func strStrSunday(haystack string, needle string) int {
	m, n := len(haystack), len(needle)
	posTable := make([]int, 128)
	for i := 0; i < 128; i++ {
		posTable[i] = -1
	}
	for i := 0; i < n; i++ {
		posTable[needle[i]] = i
	}

	i, j := 0, 0
	for i <= m - n {
		j = 0
		for j < n {
			if haystack[i] == needle[j] {
				i++
				j++
			} else {
				index := i + n - j
				if index >= m {
					return -1
				}
				if posTable[haystack[index]] == -1 {
					i = index + 1
				} else {
					i = index - posTable[haystack[index]]
				}
				break
			}
		}
		if j == n {
			return i - n
		}
	}
	return -1
}