package asciiartfs
func selectFile(filename string) string {
	switch filename {
	case "thinkertoy":
		filename = "thinkertoy.txt"
	case "shadow":
		filename = "shadow.txt"
	default:
		filename = "standard.txt"
	}
	return filename
}