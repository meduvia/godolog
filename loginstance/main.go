package log

import (
	"github.com/meduvia/godolog/log"
	"github.com/meduvia/godolog/writer"
)

type LogInstance struct {
	writer *writer.WriterInstance
	level log.FlagLevel
}

func NewLogInstance(logbackend *writer.WriterInstance, level log.FlagLevel) *LogInstance{
	return &LogInstance{
		writer: logbackend,
		level: level,
	}
}

func (logger *LogInstance) Log(product string, service string, location string, level log.FlagLevel, code string, message string) error {
	if (log.FlagCompareOneMoreThanTwo(logger.level, level)) {
		return nil // break log level requirement is more than log level
	}
	logobj := log.CreateLog(product, service, location, level,code, message)
	return (*logger.writer).Write(logobj)
}

