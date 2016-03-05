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

func (r *Result) AddLine(text string, lineType int) {
	switch lineType {
	case UNMODIFIED:
		r.Lines = append(r.Lines, ResultLine{text, true, false, false})
	case ADDED:
		r.Lines = append(r.Lines, ResultLine{text, false, true, false})
	case REMOVED:
		r.Lines = append(r.Lines, ResultLine{text, false, false, true})
	}
}

func (r *Result) String() (s string) {
	for _, l := range r.Lines {
		switch {
		case l.IsUnmodified:
			s += ("   ")
		case l.IsAdded:
			s += (" + ")
		case l.IsRemoved:
			s += (" - ")
		}
		s += l.Text + "\n"
	}
	return
}
