package changelog

func New() *ChangeLog {
	return &ChangeLog{
		entries: make(map[ChangeType][]Entry),
	}
}

type ChangeLog struct {
	entries map[ChangeType][]Entry
}

func (c *ChangeLog) AddEntry(entry Entry) {
	if entry.ChangeType == ChangeTypeNone {
		return
	}

	if len(c.entries[entry.ChangeType]) == 0 {
		c.entries[entry.ChangeType] = []Entry{}
	}

	c.entries[entry.ChangeType] = append(c.entries[entry.ChangeType], entry)
}

func (c *ChangeLog) String() string {
	output := ""
	for changeType, entries := range c.entries {
		output += string(changeType) + "\n"
		for _, entry := range entries {
			output += entry.String() + "\n"
		}
	}

	return output
}

func (c *ChangeLog) Write(formatter Formatter) string {
	output := ""
	for changeType, entries := range c.entries {
		output += formatter.BeforeGroup(changeType)
		for _, entry := range entries {
			output += formatter.FormatEntry(entry)
		}
		output += formatter.AfterGroup(changeType)
	}

	return output
}
