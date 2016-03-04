package diff

import ()

func findLine(text []string, s string) (int, bool) {
	for i, l := range text {
		if l == s {
			return i, true
		}
	}
	return 0, false
}

func ModifiedLines(oldText []string, newText []string) (add, remove []int) {
	newTextLine := 0
	for oldTextLine, curStr := range oldText {
		if foundLine, ok := findLine(newText[newTextLine:], curStr); !ok {
			remove = append(remove, oldTextLine)
			continue
		} else {
			//found situation
			for i := newTextLine; i < newTextLine+foundLine; i++ {
				add = append(add, i)
			}
			newTextLine = newTextLine + foundLine + 1
		}
	}
	//rest newText's lines are added
	for i := newTextLine; i < len(newText); i++ {
		add = append(add, i)
	}
	return
}
