package logger

type emptyWriter struct {
}

func (e *emptyWriter) Write(p []byte) (n int, err error)  {
	return len(p),nil
}
