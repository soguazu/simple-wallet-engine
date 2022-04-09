package logger

import (
	"fmt"
	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
	"github.com/soguazu/boilerplate_golang/internals/core/ports"
	"github.com/soguazu/boilerplate_golang/pkg/config"
	"gopkg.in/sohlich/elogrus.v3"
	"io"
	"os"
	"runtime"
	"strings"
)

type logger struct {
	logger *log.Logger
}

// NewLogger create an instance of the logger
func NewLogger(l *log.Logger) ports.ILogger {
	return &logger{logger: l}
}

func (l *logger) SetFormater() {
	formatter := &log.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05", // the "time" field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			// this function is required when you want to introduce your custom format.
			// In my case I wanted file and line to look like this `file="engine.go:141`
			// but f.File provides a full path along with the file name.
			// So in `formatFilePath()` function I just trimmed everything before the file name
			// and added a line number in the end
			return "", fmt.Sprintf("%s:%d", formatFilePath(f.File), f.Line)
		},
	}

	l.logger.SetReportCaller(true)
	l.logger.SetFormatter(formatter)
}

// MakeLogger creates an instance of a logger
func (l *logger) MakeLogger(filename string, display bool) *log.Logger {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err.Error())
	}

	l.SetFormater()

	if display {
		l.logger.SetOutput(io.MultiWriter(os.Stdout, f))
	} else {
		l.logger.SetOutput(io.MultiWriter(f))
	}
	return l.logger

}

func (l *logger) Hook() *log.Logger {
	client, err := elastic.NewClient(elastic.SetURL(config.Instance.ElasticURL))
	if err != nil {
		log.Panic(err)
	}
	hook, err := elogrus.NewAsyncElasticHook(client, "localhost", log.DebugLevel, "mylog")
	if err != nil {
		log.Panic(err)
	}
	l.SetFormater()
	l.logger.Hooks.Add(hook)
	return l.logger
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
