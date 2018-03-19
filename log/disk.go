package log

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/nephele/context"
)

type diskLogger struct {
	path string
}

func (l *diskLogger) Printf(ctx context.Context, level string, format string, values ...interface{}) {
	f, err := os.OpenFile(l.path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	message := fmt.Sprintf(format, values...)
	result := fmt.Sprintf("[%s] %s [%s]\t\"%s\"\n", level, "contextid", time.Now().Format(time.RFC3339), message)
	f.Write([]byte(result))
}

func (l *diskLogger) Printw(ctx context.Context, level string, message string, keysAndValues ...interface{}) {
	f, err := os.OpenFile(l.path, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return
	}
	defer f.Close()
	resultBuffer := bytes.NewBuffer([]byte(fmt.Sprintf("[%s] %s [%s]", level, "contextid", time.Now().Format(time.RFC3339))))
	for i := 0; i < len(keysAndValues)/2; i++ {
		resultBuffer.WriteString(fmt.Sprintf("\t\"%v\"", keysAndValues[i*2]))
		resultBuffer.WriteString(fmt.Sprintf("\t%v", keysAndValues[i*2+1]))
	}
	resultBuffer.WriteString(fmt.Sprintf("\t\"%s\"", message))
	resultBuffer.WriteString("\n")
	f.Write(resultBuffer.Bytes())
}

type diskConfig struct {
	path string `toml:"path"`
}

func (dc *diskConfig) BuildLogger() (Logger, error) {
	return &diskLogger{dc.path}, nil
}