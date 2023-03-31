package stringx

// Ext returns the file name extension used by path.
// The extension is the suffix beginning at the final dot
// it is empty if there is no dot.
func Ext(path string, sub byte) string {
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == sub {
			if i < len(path)-1 {
				return path[i+1:]
			}
			return ""
		}
	}
	return ""
}
