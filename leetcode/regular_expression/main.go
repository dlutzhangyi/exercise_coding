package main

import "fmt"

func match(s string, p string) bool {
	if len(s) == 0 || len(p) == 0 {
		return false
	}
	return matchCore(s, len(s)-1, p, len(p)-1)
}

func matchCore(s string, i int, p string, j int) bool {
	if j == -1 {
		if i == -1 {
			return true
		} else {
			return false
		}

	}

	if p[j] == '*' {
		if i > -1 && (p[j-1] == '.' || p[j-1] == s[i]) {
			if matchCore(s, i-1, p, j) {
				return true
			}
		}
		return matchCore(s, i, p, j-2)
	}
	if i > -1 && (p[j] == '.' || p[j] == s[i]) {
		return matchCore(s, i-1, p, j-1)
	}

	return false
}

func main() {
	strs := []string{"aa", "a", "aab", "ab", "ab"}
	patterns := []string{"a", "ab*", "c*a*b", ".*", ".*c"}
	for index, str := range strs {
		fmt.Printf("str:%s pattern:%s\n", str, patterns[index])
		fmt.Printf("Is match:%v\n", match(str, patterns[index]))
	}
}
