package helper

import "errors"

func StringToUint64(s string) (uint64, error) {
	if s == "" {
		return 0, nil
	}
	var result uint64
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, errors.New("failed parsing uint64")
		}
		result = result*10 + uint64(c-'0')
	}
	return result, nil
}

func PtrOrDefault[T any](p *T, def T) T {
	if p != nil {
		return *p
	}
	return def
}
