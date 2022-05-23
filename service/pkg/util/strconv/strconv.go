package strconv

import "strconv"

// UItoa is equivalent to FormatUint(uint64(i), 10).
func UItoa(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func ParseUint32(s string) (uint32, error) {
	n, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(n), nil
}
