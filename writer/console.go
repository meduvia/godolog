package writer

import "github.com/meduvia/godolog/log"

type ConsoleWriterInterface interface {
	// we did not add verbose because in major logger backend verbose does not exist
	Debug(string) // verbose + debug
	Info(string)  // info
	Warn(string)  // warn
	Error(string) // error + fatal
	// We did not add fatal to prevent logger backend from killing the running program
}

type ConsoleWriter struct {
	writer *ConsoleWriterInterface
}

func NewConsoleWriter(LogBackend *ConsoleWriterInterface) *ConsoleWriter {
	return &ConsoleWriter{
		writer: LogBackend,
	}
} 

func (consolewriter *ConsoleWriter) Write(logobj *log.Log) {
	logstring, err := logobj.ToJSONString()
	if err != nil {
		(*consolewriter.writer).Error("Error Writing log")
	}
	switch logobj.Level {
	case log.FlagVerbose, log.FlagDebug:
		(*consolewriter.writer).Debug(logstring)
	case log.FlagInfo:
		(*consolewriter.writer).Info(logstring)
	case log.FlagWarn:
		(*consolewriter.writer).Warn(logstring)
	case log.FlagError, log.FlagFatal:
		(*consolewriter.writer).Error(logstring)
	default:
		(*consolewriter.writer).Debug(logstring)
	}

}
