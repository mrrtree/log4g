package log4g

import (
	"testing"
)

func Test_GetLogger_With_NoConf(t *testing.T) {
	logger := GetLogger("log1")
	logger.Debug("Debug Logger")
	logger.Info("Info Logger")
	logger.Error("Error Logger")
	t.Log("test success")
}

func Test_GetLoggerWithId_With_NoConf(t *testing.T) {
	logger := GetLoggerWithId("log2", "12345678")
	logger.Debug("Debug Logger")
	logger.Info("Info Logger")
	logger.Error("Error Logger")
	t.Log("test success")
}

func Test_GetLogger_With_Conf(t *testing.T) {
	logger := GetLogger("mylog")
	logger.Debug("Debug Logger")
	logger.Info("Info Logger")
	logger.Error("Error Logger")
	t.Log("test sucess")
}
