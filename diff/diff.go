package diff

func findLine(text []string, s string) (pos int, ok bool) {
	for i, l := range text {
		if l == s {
			return i, true
		}
	}
	return -1, false
}

func findSameLine(text1 []string, text2 []string) (text1pos int, text2pos int, ok bool) {
	for text1pos, s := range text1 {
		if text2pos, ok := findLine(text2, s); ok {
			return text1pos, text2pos, ok
		}
	}
	return -1, -1, false
}

func LineDiff(oldText []string, newText []string) (result LineDiffResult) {
	var addBlock = func(opos, npos int, oldText, newText []string) {
		for i := 0; i < opos; i++ {
			result.AddLine(oldText[i], REMOVED)
		}
		for i := 0; i < npos; i++ {
			result.AddLine(newText[i], ADDED)
		}
		result.AddLine(oldText[opos], UNMODIFIED)
	}
	//add blocks of {removedLines,addedLines,anUnmodifiedLine}
	oldProcPos, newProcPos := 0, 0
	for oldProcPos < len(oldText) && newProcPos < len(newText) {
		if opos, npos, ok := findSameLine(oldText[oldProcPos:], newText[newProcPos:]); ok {
			addBlock(opos, npos, oldText[oldProcPos:], newText[newProcPos:])
			oldProcPos, newProcPos = oldProcPos+opos+1, newProcPos+npos+1
		} else {
			break
		}
	}
	//add rest lines
	for ; oldProcPos < len(oldText); oldProcPos++ {
		result.AddLine(oldText[oldProcPos], REMOVED)
	}
	for ; newProcPos < len(newText); newProcPos++ {
		result.AddLine(newText[newProcPos], ADDED)
	}
	result.CalOrder()
	result.RemoveLineBreak()
	return
}
