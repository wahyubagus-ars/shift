package config

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"os"
)

func InitLog() {
	log.SetLevel(getLoggerLevel(os.Getenv("LOG_LEVEL")))
	log.SetReportCaller(true)
	log.SetFormatter(&nested.Formatter{
		HideKeys:        true,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		ShowFullLevel:   true,
		CallerFirst:     true,
	})

}

func getLoggerLevel(value string) log.Level {
	fmt.Println("log level :: " + value)
	switch value {
	case "DEBUG":
		return log.DebugLevel
	case "TRACE":
		return log.TraceLevel
	default:
		return log.InfoLevel
	}
}
