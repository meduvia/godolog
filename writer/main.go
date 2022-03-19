package writer

import "github.com/meduvia/godolog/log"

type WriterInstance interface {
	Write(*log.Log) error
}

