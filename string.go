package jb

// TruncateString truncates the given string limiting it by the given num
func TruncateString(str string, num int) string {
	s := str
	if len(str) > num {
		if num > 3 {
			num -= 3
		}
		s = str[0:num] + "..."
	}
	return s
}
