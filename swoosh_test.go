package swoosh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testAddr = ":7799"
)

func TestSwoosh_EnableLog(t *testing.T) {
	sln, err := Listen("tcp", testAddr)
	assert.NoErrorf(t, err, "is address %s unavailable ?", testAddr)

	t.Run("log level should be TRACE and report caller should be on",
		func(t *testing.T) {
			sln.EnableLog(TRACE_LEVEL)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, TRACE_LEVEL, actualLevel, "expected %d, got %d",
				TRACE_LEVEL, actualLevel)

			assert.True(t, sln.logger.ReportCaller, "report caller should be on")
		})

	t.Run("log level should be DEBUG and report caller should be on",
		func(t *testing.T) {
			sln.EnableLog(DEBUG_LEVEL)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, DEBUG_LEVEL, actualLevel, "expected %d, got %d",
				DEBUG_LEVEL, actualLevel)
		})

	t.Run("log level should be INFO and report caller should be of",
		func(t *testing.T) {
			sln.EnableLog(INFO_LEVEL)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, INFO_LEVEL, actualLevel, "expected %d, got %d",
				INFO_LEVEL, actualLevel)
		})

	t.Run("log level should be ERROR and report caller should be of",
		func(t *testing.T) {
			sln.EnableLog(ERROR_LEVEL)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, ERROR_LEVEL, actualLevel, "expected %d, got %d",
				ERROR_LEVEL, actualLevel)
		})

	t.Run("log level should be WARN and report caller should be of",
		func(t *testing.T) {
			sln.EnableLog(WARN_LEVEL)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, WARN_LEVEL, actualLevel, "expected %d, got %d",
				WARN_LEVEL, actualLevel)
		})

	t.Run("log level should be FATAL and report caller should be of",
		func(t *testing.T) {
			sln.EnableLog(FATAL_LEVEL)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, FATAL_LEVEL, actualLevel, "expected %d, got %d",
				FATAL_LEVEL, actualLevel)
		})

	t.Run("log level should be FATAL when unknown level provided",
		func(t *testing.T) {
			sln.EnableLog(-1)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, FATAL_LEVEL, actualLevel, "expected %d, got %d",
				FATAL_LEVEL, actualLevel)
		})
}
