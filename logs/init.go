package logs

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

type Log struct {
	logger      zerolog.Logger
	serviceName string
}

var appLogger Log

func init() {
	appLogger = NewLogger("app")
}

func getDirectory(location *time.Location, serviceName string) io.Writer {
	appName, ok := os.LookupEnv("APP_NAME")

	if !ok {
		panic("APP_NAME not found")
	}

	appVersion, ok := os.LookupEnv("APP_VERSION")

	if !ok {
		panic("APP_VERSION not found")
	}

	logRootDir, ok := os.LookupEnv("LOG_ROOT_DIR")

	if !ok {
		panic("LOG_ROOT_DIR not found")
	}

	appDir := fmt.Sprintf("%s-%s", appName, appVersion)
	logDir := fmt.Sprintf("%s/%s/%s.log", logRootDir, appDir, serviceName)
	dir := path.Dir(logDir)

	// Generate Root Directory
	if _, err := os.Stat(logRootDir); os.IsNotExist(err) {
		err = os.Mkdir(logRootDir, os.ModePerm)

		if err != nil {
			panic(err.Error())
		}
	}

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)

		if err != nil {
			panic(err.Error())
		}
	}

	f, err := os.OpenFile(logDir, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err.Error())
	}

	return io.MultiWriter(f, os.Stdout)
}

func NewLogger(serviceName string) (logger Log) {
	loc, err := time.LoadLocation("Asia/Jakarta")

	if err != nil {
		panic(err)
	}

	//pwd, err := os.Getwd()
	//pwd, err := filepath.Abs("./")

	//if err != nil {
	//	panic(err)
	//}

	writer := getDirectory(loc, serviceName)

	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(loc)
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: true}
	output.FormatLevel = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("| ")
	}
	output.FormatFieldValue = func(i interface{}) string {
		return strings.ToUpper(fmt.Sprintf("%s", i))
	}
	output.FormatMessage = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatCaller = func(i interface{}) string {
		//fmt.Println(pwd)
		//fmt.Println(strings.SplitAfter(fmt.Sprintf("%s", i), pwd)[1])
		//if enums.Environment(os.Getenv("ENVIRONMENT")) == enums.Production {
		//	return strings.SplitAfter(fmt.Sprintf("%s", i), pwd)[1]
		//}

		return fmt.Sprintf("%s", i)
	}
	output.Out = writer

	//multi := zerolog.MultiLevelWriter(output, os.Stdout)
	fileLogger := zerolog.New(output).
		With().
		CallerWithSkipFrameCount(3).
		Str("services", serviceName).
		Timestamp().
		Logger()

	logger = Log{
		logger:      fileLogger,
		serviceName: serviceName,
	}

	return logger
}

func GetAppLogger() Collections {
	return appLogger
}

func (log Log) Info(message string) {
	log.logger.Info().Msg(message)
}

func (log Log) InfoInterface(data interface{}) {
	marshaledData, _ := json.Marshal(data)
	log.logger.Info().Msg(string(marshaledData))
}

func (log Log) Error(message string) {
	log.logger.Error().Msg(message)
}

func (log Log) Debug(message string) {
	log.logger.Debug().Msg(message)
}
