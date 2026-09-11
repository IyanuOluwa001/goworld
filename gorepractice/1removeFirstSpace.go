package main

func removeFirstSpace(s string) string{
	for i := 0; i < len(s); i++{
		if s[i] ==' '{
			return s[:i] + s[i+1:]
		}
	}
	return s
}

func removeFostSpace(s string) string{
	for i, ch := range s {
		if ch == ' ' {
			return s[:i] + s[i+1:]
		}
	}
	return s
}

