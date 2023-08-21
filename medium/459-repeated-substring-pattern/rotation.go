package RepeatedSubstringPattern

// This question is actually an Easy, but it's so good
func repeatedSubstringPattern(s string) bool {
	t := s + s
	for i := 1; i+len(s) < len(t); i++ {
		if t[i:i+len(s)] == s {
			return true
		}
	}
	return false
}
