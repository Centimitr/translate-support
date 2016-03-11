package diff

import ()

type Line struct {
	Text         string
	IsUnmodified bool
	IsAdded      bool
	IsRemoved    bool
}

type LineDiffResult struct {
	Lines []Line
}

const (
	UNMODIFIED = iota
	ADDED
	REMOVED
)

func (r *LineDiffResult) AddLine(text string, lineType int) {
	switch lineType {
	case UNMODIFIED:
		r.Lines = append(r.Lines, Line{text, true, false, false})
	case ADDED:
		r.Lines = append(r.Lines, Line{text, false, true, false})
	case REMOVED:
		r.Lines = append(r.Lines, Line{text, false, false, true})
	}
}

func (r *LineDiffResult) String() (s string) {
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
