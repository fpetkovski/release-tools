package changelog

import (
	"fmt"
	"strings"
)

type Formatter interface {
	BeforeGroup(changeType ChangeType) string
	AfterGroup(changeType ChangeType) string
	FormatEntry(entry Entry) string
}

type FlatFormatter struct {}

func (s FlatFormatter) BeforeGroup(changeType ChangeType) string {
	return ""
}

func (s FlatFormatter) AfterGroup(changeType ChangeType) string {
	return ""
}

func (s FlatFormatter) FormatEntry(entry Entry) string {
	return fmt.Sprintf("[%s] %s (%s)", strings.ToUpper(string(entry.ChangeType)), entry.ChangeDescription, entry.PullURL)
}

type GroupFormatter struct {}

func (g GroupFormatter) BeforeGroup(changeType ChangeType) string {
	return strings.ToUpper(string(changeType)) + "\n"
}

func (g GroupFormatter) AfterGroup(changeType ChangeType) string {
	return "\n"
}

func (g GroupFormatter) FormatEntry(entry Entry) string {
	descriptionParts := strings.Split(entry.ChangeDescription, " ")
	descriptionParts[0] = strings.Title(descriptionParts[0])
	description := strings.Join(descriptionParts, " ")

	return fmt.Sprintf("%s (%s)\n", description, entry.PullURL)
}
