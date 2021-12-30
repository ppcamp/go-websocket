package utils

func StartsWith(str, prefix string) bool {
	lp := len(prefix)
	return str[:lp] == prefix
}
