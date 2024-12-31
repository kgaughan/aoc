package day4

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func adventCoin(prefix string, nZeroes, startFrom, endAt int) int {
	hexPrefix := strings.Repeat("0", nZeroes)

	for i := startFrom; i <= endAt; i++ {
		h := md5.New()
		_, _ = io.WriteString(h, prefix)
		_, _ = io.WriteString(h, fmt.Sprint(i))

		hex := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(hex, hexPrefix) {
			return i
		}
	}
	return 0
}
