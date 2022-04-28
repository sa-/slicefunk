package slicefunk

func Map[T, U any](s []T, f func(T) U) []U {
	r := make([]U, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}

func Filter[T any](s []T, f func(T) bool) []T {
	r := make([]T, len(s))
	counter := 0
	for i := 0; i < len(s); i++ {
		if f(s[i]) {
			r[counter] = s[i]
			counter++
		}
	}
	return r[:counter]
}

func Unique[T comparable](s []T) []T {
    inResult := make(map[T]bool)
    var result []T
    for _, str := range s {
        if _, ok := inResult[str]; !ok {
            inResult[str] = true
            result = append(result, str)
        }
    }
    return result
}
