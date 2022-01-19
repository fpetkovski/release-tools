package changelog

import "fmt"

type ChangeType string

const ChangeTypeNone ChangeType = "none"

type Entry struct {
	ChangeType        ChangeType
	ChangeDescription string
	PullURL           string
}

func NewEntry(changeType ChangeType, description, prURL string) Entry {
	return Entry{
		ChangeType:        changeType,
		ChangeDescription: description,
		PullURL:           prURL,
	}
}

func (e Entry) String() string {
	return fmt.Sprintf("%s (%s)", e.ChangeDescription, e.PullURL)
}
