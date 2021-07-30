package commontools

import "github.com/kacpekwasny/pyfuncs"

// InSlice f
func InSlice(match string, slice []string) (int, bool) {
	for i, str := range slice {
		if str == match {
			return i, true
		}
	}
	return 0, false
}

// InStr d
func InStr(str, match string) bool {
	ln := len(match)
	// Avoid error
	if ln > len(str) {
		return false
	}
	// Shortcut
	if str == match {
		return true
	}
	// Special case: "" is not in any string unless string is ""
	if match == "" && str != "" {
		return false
	}
	for i := range pyfuncs.Range(float64(len(str)) - float64(ln)) {
		j := int(i)
		if str[j:j+ln] == match {
			return true
		}
	}
	return false

}
