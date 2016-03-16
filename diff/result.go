package diff

import (
	"strings"
)

type Line struct {
	Text         string `json:"text"`
	IsUnmodified bool   `json:"isUnmodified"`
	IsAdded      bool   `json:"isAdded"`
	IsRemoved    bool   `json:"isRemoved"`
	NewOrder     int    `json:"newOrder"`
	OldOrder     int    `json:"oldOrder"`
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
		r.Lines = append(r.Lines, Line{text, true, false, false, 0, 0})
	case ADDED:
		r.Lines = append(r.Lines, Line{text, false, true, false, 0, 0})
	case REMOVED:
		r.Lines = append(r.Lines, Line{text, false, false, true, 0, 0})
	}
}

func (r *LineDiffResult) CalOrder() {
	newOrder := 1
	oldOrder := 1
	for i, l := range r.Lines {
		switch {
		case l.IsUnmodified:
			r.Lines[i].NewOrder = newOrder
			r.Lines[i].OldOrder = oldOrder
			newOrder++
			oldOrder++
		case l.IsAdded:
			r.Lines[i].NewOrder = newOrder
			newOrder++
		case l.IsRemoved:
			r.Lines[i].OldOrder = oldOrder
			oldOrder++
		}
	}
}

func (r *LineDiffResult) TabReplace(replaceStr string) {
	for i, l := range r.Lines {
		r.Lines[i].Text = strings.Replace(l.Text, "\u0009", replaceStr, -1)
	}
}
func (r *LineDiffResult) RemoveLineBreak() {
	for i, l := range r.Lines {
		s := strings.Replace(l.Text, "\n", "=", -1)
		s = strings.Replace(s, "\r", "=", -1)
		r.Lines[i].Text = s
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
