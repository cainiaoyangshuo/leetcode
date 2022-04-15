package _68validIPAddress

import (
	"strconv"
	"strings"
	"unicode"
)

func ValidIPAddress(queryIP string) string {
	if isIpv4(queryIP) {
		return "IPv4"
	}

	if isIpv6(queryIP) {
		return "IPv6"
	}

	return "Neither"
}

func isIpv4(str string) bool {
	strSplit := strings.Split(str, ".")
	if len(strSplit) != 4 {
		return false
	}

	for _, v := range strSplit {
		n, _ := strconv.Atoi(v)
		if v == "" || len(v) > 3 || n > 255 || n < 0 {
			return false
		}
		switch len(v) {
		case 3:
			if n/100 == 0 {
				return false
			}
			break
		case 2:
			if n/10 == 0 {
				return false
			}
			break
		}
	}

	return true
}

func isIpv6(str string) bool {
	strSplit := strings.Split(str, ":")
	if len(strSplit) != 8 {
		return false
	}

	for _, n := range strSplit {
		length := len(n)
		if length == 0 || length > 4 {
			return false
		}

		for _, c := range n {
			if !unicode.IsNumber(c) && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
				return false
			}
		}
	}
	return true
}