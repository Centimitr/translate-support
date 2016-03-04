package diff

import ()

type ResultLine struct {
	Text         string
	IsUnmodified bool
	IsAdded      bool
	IsRemoved    bool
}

type Result struct {
	Lines []ResultLine
}

const (
	UNMODIFIED = iota
	ADDED
	REMOVED
)

func (r *Result) add(line string, lineType int) {
	switch lineType {
	case UNMODIFIED:
		r.Lines = append(r.Lines, ResultLine{line, true, false, false})
	case ADDED:
		r.Lines = append(r.Lines, ResultLine{line, false, true, false})
		//here need to swap the last 2 Lines
		//because the iteration below is on the oldText
		r.Lines[len(r.Lines)-2], r.Lines[len(r.Lines)-1] = r.Lines[len(r.Lines)-1], r.Lines[len(r.Lines)-2]
	case REMOVED:
		r.Lines = append(r.Lines, ResultLine{line, false, false, true})
	}
}

func findLine(text []string, s string) (int, bool) {
	for i, l := range text {
		if l == s {
			return i, true
		}
	}
	return 0, false
}

func ModifiedLines(oldText []string, newText []string) (result Result) {
	newTextLine := 0
	for _, curStr := range oldText {
		if foundLine, ok := findLine(newText[newTextLine:], curStr); !ok {
			result.add(curStr, REMOVED)
			continue
		} else {
			//found situation
			for i := newTextLine; i < newTextLine+foundLine; i++ {
				result.add(newText[i], ADDED)
			}
			newTextLine = newTextLine + foundLine + 1
			result.add(curStr, UNMODIFIED)
		}
	}
	//rest newText's lines are added
	for i := newTextLine; i < len(newText); i++ {
		result.add(newText[i], ADDED)

	}
	return
}
