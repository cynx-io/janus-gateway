package helper

import "errors"

func StringToint32(s string) (int32, error) {
	if s == "" {
		return 0, nil
	}
	var result int32
	for _, c := range s {
		if c < '0' || c > '9' {
			return 0, errors.New("failed parsing int32")
		}
		result = result*10 + int32(c-'0')
	}
	return result, nil
}

func PtrOrDefault[T any](p *T, def T) T {
	if p != nil {
		return *p
	}
	return def
}
