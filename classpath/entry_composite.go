package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (entry CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, subEntry := range entry {
		data, from, err := subEntry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (entry CompositeEntry) String() string {
	strs := make([]string, len(entry))
	for i, entry := range entry {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
