package logs

// Collections is log's collection of function
type Collections interface {
	InfoInterface(data interface{})
	Info(message string)
	Error(message string)
	Debug(message string)
}
