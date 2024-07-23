package logger

// StringerFunc ...
type StringerFunc func() string

// String ...
func (f StringerFunc) String() string {
	return f()
}
