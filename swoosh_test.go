package swoosh

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testAddr = ":7799"
)

func TestSwoosh_EnableLog(t *testing.T) {
	sln, err := ListenAndServe("tcp", testAddr)
	assert.NoErrorf(t, err, "is address %s unavailable ?", testAddr)

	t.Run("log level should be TRACE and report caller should be on",
		func(t *testing.T) {
			sln.EnableLog(TraceLevel)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, TraceLevel, actualLevel, "expected %d, got %d",
				TraceLevel, actualLevel)

			assert.True(t, sln.logger.ReportCaller, "report caller should be on")
		})

	t.Run("log level should be DEBUG and report caller should be on",
		func(t *testing.T) {
			sln.EnableLog(DebugLevel)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, DebugLevel, actualLevel, "expected %d, got %d",
				DebugLevel, actualLevel)
		})

	t.Run("log level should be INFO and report caller should be of",
		func(t *testing.T) {
			sln.EnableLog(InfoLevel)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, InfoLevel, actualLevel, "expected %d, got %d",
				InfoLevel, actualLevel)
		})

	t.Run("log level should be ERROR and report caller should be of",
		func(t *testing.T) {
			sln.EnableLog(ErrorLevel)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, ErrorLevel, actualLevel, "expected %d, got %d",
				ErrorLevel, actualLevel)
		})

	t.Run("log level should be WARN and report caller should be of",
		func(t *testing.T) {
			sln.EnableLog(WarnLevel)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, WarnLevel, actualLevel, "expected %d, got %d",
				WarnLevel, actualLevel)
		})

	t.Run("log level should be FATAL and report caller should be of",
		func(t *testing.T) {
			sln.EnableLog(FatalLevel)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, FatalLevel, actualLevel, "expected %d, got %d",
				FatalLevel, actualLevel)
		})

	t.Run("log level should be FATAL when unknown level provided",
		func(t *testing.T) {
			sln.EnableLog(-1)

			actualLevel := sln.GetLogLevel()
			assert.Equalf(t, FatalLevel, actualLevel, "expected %d, got %d",
				FatalLevel, actualLevel)
		})
}
