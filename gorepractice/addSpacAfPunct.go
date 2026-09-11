package main

import "strings"


func addSpaceAfterPunc(s string) string {
	var sb strings.Builder
	for i := 0; i < len(s); i++ {
		sb.WriteByte(s[i])
		if strings.ContainsRune(".,!?:;", rune(s[i])) {
			sb.WriteByte(' ')
		}
	}
	return sb.String()
}
