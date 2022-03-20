package log

import (
	"encoding/json"
	"fmt"
	"time"
)

type FlagLevel string

const (
	FlagVerbose FlagLevel = "verb"
	FlagDebug   FlagLevel = "debug"
	FlagInfo    FlagLevel = "info"
	FlagWarn    FlagLevel = "warn"
	FlagError   FlagLevel = "error"
	FlagFatal   FlagLevel = "fatal"
	FlagNone    FlagLevel = "none"
)

type Log struct {
	Product   string    `json:"product"`
	Service   string    `json:"service"`
	Location  string    `json:"location"`
	Timestamp string    `json:"timestamp"`
	Level     FlagLevel `json:"level"`
	Code      string    `json:"code"`
	Message   string    `json:"message"`
}

func CreateLog(product string, service string, location string, level FlagLevel, code string, message string) *Log {
	return &Log{
		Product:   product,
		Service:   service,
		Location:  location,
		Timestamp: fmt.Sprint(time.Now().UnixMilli()),
		Level:     level,
		Code:      code,
		Message:   message,
	}
}

func (log *Log) ToJSONString() (string, error) {
	b, err := json.Marshal(log)
	return string(b), err
}

// type FlagCompareOp int

// const (
// 	FlagL FlagCompareOp = iota
// 	FlagLEQ
// 	FlagEQ
// 	FlagMEQ
// 	FlagM
// )

func FlagCompareOneMoreThanTwo(one FlagLevel, two FlagLevel) bool {
	switch one {
	case FlagVerbose:
		return false
	case FlagDebug:
		return !(two != FlagVerbose)
	case FlagInfo:
		return !(two != FlagVerbose && two != FlagDebug)
	case FlagWarn:
		return !(two != FlagVerbose && two != FlagDebug && two != FlagInfo)
	case FlagError:
		return !(two != FlagVerbose && two != FlagDebug && two != FlagInfo && two != FlagWarn)
	case FlagFatal:
		return !(two != FlagVerbose && two != FlagDebug && two != FlagInfo && two != FlagWarn && two != FlagError)
	case FlagNone:
		return true
	default:
		return false
	}
}
