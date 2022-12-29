package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func NewLogger(path string) (*logrus.Logger, func() error) {
	logger := logrus.New()

	currentTime := time.Now()

	formatted := path + fmt.Sprintf("Date:%d.%02d.%02d__Time:%02d:%02d:%02d",
		currentTime.Year(), currentTime.Month(), currentTime.Day(),
		currentTime.Hour(), currentTime.Minute(), currentTime.Second()) + ".log"

	f, err := os.OpenFile(formatted, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Open file error: [%v]", err)
	}

	logger.SetOutput(f)
	logger.Writer()
	logger.SetLevel(logrus.InfoLevel)

	formatter := &logrus.TextFormatter{}
	formatter.DisableQuote = true

	logger.SetFormatter(formatter)

	return logger, f.Close
}
