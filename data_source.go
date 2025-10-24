package miniexpr

// DataSource helps with multiple source types.
// We're mostly using it for convenience for the scanner to accept slices of bytes and strings.
type DataSource interface {
	Len() int
	At(i int) byte
}

// ByteSource is a []byte data source.
type ByteSource []byte

// Len tells the number of bytes in slice.
func (s ByteSource) Len() int {
	return len(s)
}

// At gives us the byte at index.
func (s ByteSource) At(i int) byte {
	return s[i]
}

// StringSource is a regular string data source.
type StringSource string

// Len tells us the string length in bytes.
func (s StringSource) Len() int {
	return len(s)
}

// At gives us the byte at index.
func (s StringSource) At(i int) byte {
	return s[i]
}
