package log_test

import (
	"encoding/json"
	"math/rand"
	"testing"

	"github.com/meduvia/godolog/log"
)

func TestCreateLog(t *testing.T) {
	product := "App"
	service := "Service"
	location := "EU01-D01-R01-N01"
	level := log.FlagDebug
	code := "OP-R403"
	message := "Rest API Error 403"
	logobj_create := log.CreateLog(product, service, location, level, code, message)
	logobj_test := &log.Log{
		Product:   product,
		Service:   service,
		Location:  location,
		Timestamp: logobj_create.Timestamp, // this line is to bypass the error caused by timestamp missmatch
		Level:     level,
		Code:      code,
		Message:   message,
	}
	if *logobj_create != *logobj_test {
		t.Error("Error field attribution")
	}
}

func TestLogToJSONString(t *testing.T) {
	product := "App"
	service := "Service"
	location := "EU01-D01-R01-N01"
	level := log.FlagDebug
	code := "OP-R403"
	message := "Rest API Error 403"
	logobj_create := log.CreateLog(product, service, location, level, code, message)
	logobj_test := &log.Log{
		Product:   product,
		Service:   service,
		Location:  location,
		Timestamp: logobj_create.Timestamp, // this line is to bypass the error caused by timestamp missmatch
		Level:     level,
		Code:      code,
		Message:   message,
	}
	logobj_create_jsonstring, err := logobj_create.ToJSONString()
	if err != nil {
		t.Error("Error json string func")
	}
	logobj_test_jsonstring, err := json.Marshal(logobj_test)
	if logobj_create_jsonstring != string(logobj_test_jsonstring) {
		t.Error("Error field attribution")
	}
}

func FuzzLogToJSONString(f *testing.F) {
	log_level_avaliable := []log.FlagLevel{log.FlagVerbose, log.FlagDebug, log.FlagWarn, log.FlagError, log.FlagFatal}
	product := "App"
	service := "Service"
	location := "EU01-D01-R01-N01"
	level := 10
	code := "OP-R403"
	message := "Rest API Error 403"
	f.Add(product, service, location, level, code, message)
	f.Fuzz(func(t *testing.T, product string, service string, location string, level_raw int, code string, message string) {
		level := log_level_avaliable[rand.Intn(len(log_level_avaliable))]
		logobj_create := log.CreateLog(product, service, location, level, code, message)
		logobj_test := &log.Log{
			Product:   product,
			Service:   service,
			Location:  location,
			Timestamp: logobj_create.Timestamp, // this line is to bypass the error caused by timestamp missmatch
			Level:     level,
			Code:      code,
			Message:   message,
		}
		logobj_create_jsonstring, err := logobj_create.ToJSONString()
		if err != nil {
			f.Error("Error json string func")
		}
		logobj_test_jsonstring, err := json.Marshal(logobj_test)
		if logobj_create_jsonstring != string(logobj_test_jsonstring) {
			f.Error("Error field attribution")
		}
		println(logobj_create_jsonstring)
	})
}

func TestFlagCompareOneMoreThanTwo(t *testing.T) {
	t.Run("Test Verbose always false", func(t *testing.T) {
		one := log.FlagVerbose
		tocompare_up := []log.FlagLevel{log.FlagVerbose, log.FlagDebug, log.FlagInfo, log.FlagWarn, log.FlagError, log.FlagFatal}
		tocompare_down := []log.FlagLevel{}
		// upper or eq flag => log should be displayed
		for _, flag := range tocompare_up {
			if log.FlagCompareOneMoreThanTwo(one, flag) != false {
				t.Fail()
			}
		}
		// down => log should be discarded
		for _, flag := range tocompare_down {
			if log.FlagCompareOneMoreThanTwo(one, flag) != true {
				t.Fail()
			}
		}
	})
	t.Run("Test Debug  ", func(t *testing.T) {
		one := log.FlagDebug
		tocompare_up := []log.FlagLevel{log.FlagDebug, log.FlagInfo, log.FlagWarn, log.FlagError, log.FlagFatal}
		tocompare_down := []log.FlagLevel{log.FlagVerbose}
		// upper or eq flag => log should be displayed
		for _, flag := range tocompare_up {
			if log.FlagCompareOneMoreThanTwo(one, flag) != false {
				t.Fail()
			}
		}
		// down => log should be discarded
		for _, flag := range tocompare_down {
			if log.FlagCompareOneMoreThanTwo(one, flag) != true {
				t.Fail()
			}
		}
	})
	t.Run("Test Info", func(t *testing.T) {
		one := log.FlagInfo
		tocompare_up := []log.FlagLevel{log.FlagInfo, log.FlagWarn, log.FlagError, log.FlagFatal}
		tocompare_down := []log.FlagLevel{log.FlagVerbose, log.FlagDebug}
		// upper or eq flag => log should be displayed
		for _, flag := range tocompare_up {
			if log.FlagCompareOneMoreThanTwo(one, flag) != false {
				t.Fail()
			}
		}
		// down => log should be discarded
		for _, flag := range tocompare_down {
			if log.FlagCompareOneMoreThanTwo(one, flag) != true {
				t.Fail()
			}
		}
	})
	t.Run("Test Warn", func(t *testing.T) {
		one := log.FlagWarn
		tocompare_up := []log.FlagLevel{log.FlagWarn, log.FlagError, log.FlagFatal}
		tocompare_down := []log.FlagLevel{log.FlagVerbose, log.FlagDebug, log.FlagInfo}
		// upper or eq flag => log should be displayed
		for _, flag := range tocompare_up {
			if log.FlagCompareOneMoreThanTwo(one, flag) != false {
				t.Fail()
			}
		}
		// down => log should be discarded
		for _, flag := range tocompare_down {
			if log.FlagCompareOneMoreThanTwo(one, flag) != true {
				t.Fail()
			}
		}
	})
	t.Run("Test Error", func(t *testing.T) {
		one := log.FlagError
		tocompare_up := []log.FlagLevel{log.FlagError, log.FlagFatal}
		tocompare_down := []log.FlagLevel{log.FlagVerbose, log.FlagDebug, log.FlagInfo, log.FlagWarn}
		// upper or eq flag => log should be displayed
		for _, flag := range tocompare_up {
			if log.FlagCompareOneMoreThanTwo(one, flag) != false {
				t.Fail()
			}
		}
		// down => log should be discarded
		for _, flag := range tocompare_down {
			if log.FlagCompareOneMoreThanTwo(one, flag) != true {
				t.Fail()
			}
		}
	})
	t.Run("Test Fatal", func(t *testing.T) {
		one := log.FlagFatal
		tocompare_up := []log.FlagLevel{log.FlagFatal}
		tocompare_down := []log.FlagLevel{log.FlagVerbose, log.FlagDebug, log.FlagInfo, log.FlagWarn, log.FlagError}
		// upper or eq flag => log should be displayed
		for _, flag := range tocompare_up {
			if log.FlagCompareOneMoreThanTwo(one, flag) != false {
				t.Fail()
			}
		}
		// down => log should be discarded
		for _, flag := range tocompare_down {
			if log.FlagCompareOneMoreThanTwo(one, flag) != true {
				t.Fail()
			}
		}
	})
	t.Run("Test None == true", func(t *testing.T) {
		one := log.FlagNone
		tocompare_up := []log.FlagLevel{}
		tocompare_down := []log.FlagLevel{log.FlagVerbose, log.FlagDebug, log.FlagInfo, log.FlagWarn, log.FlagError, log.FlagFatal}
		// upper or eq flag => log should be displayed
		for _, flag := range tocompare_up {
			if log.FlagCompareOneMoreThanTwo(one, flag) != false {
				t.Fail()
			}
		}
		// down => log should be discarded
		for _, flag := range tocompare_down {
			if log.FlagCompareOneMoreThanTwo(one, flag) != true {
				t.Fail()
			}
		}
	})
	t.Run("Test Default", func(t *testing.T) {
		var one log.FlagLevel = "default"
		tocompare_up := []log.FlagLevel{log.FlagVerbose, log.FlagDebug, log.FlagInfo, log.FlagWarn, log.FlagError, log.FlagFatal}
		tocompare_down := []log.FlagLevel{}
		// upper or eq flag => log should be displayed
		for _, flag := range tocompare_up {
			if log.FlagCompareOneMoreThanTwo(one, flag) != false {
				t.Fail()
			}
		}
		// down => log should be discarded
		for _, flag := range tocompare_down {
			if log.FlagCompareOneMoreThanTwo(one, flag) != true {
				t.Fail()
			}
		}
	})

}
