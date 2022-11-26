package _51

var space = byte(' ')

func reverseWords(s string) string {
	//convert to char array
	b := []byte(s)

	//remove extra space
	k := 0
	for i := range b {
		// fix: remove the heading spaces
		if (b[i] == space && k == 0) || (b[i] == space && b[i-1] == space) {
			continue
		}
		b[k] = b[i]
		k++
	}

	//fix: remove the last space
	if b[k-1] == space {
		k--
	}

	// reverse whole string
	reverseSingleWord(b[0:k])

	// reverse each word
	start := 0
	for i := 0; i < k; i++ {
		if b[i] == space {
			// need to reverse word [start: i]
			reverseSingleWord(b[start:i])
			start = i + 1
		}
	}
	//fix: reverse the last word
	reverseSingleWord(b[start:k])

	return string(b[0:k])

}

func reverseSingleWord(b []byte) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
}
