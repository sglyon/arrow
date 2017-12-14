package arrow

type Schema struct {
	fields      []Field
	nameToIndex map[string]int
	metadata    KeyValueMetadata
}
