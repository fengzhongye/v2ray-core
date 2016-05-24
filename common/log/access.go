package log

import (
	"fmt"

	"github.com/v2ray/v2ray-core/common/log/internal"
)

// AccessStatus is the status of an access request from clients.
type AccessStatus string

const (
	AccessAccepted = AccessStatus("accepted")
	AccessRejected = AccessStatus("rejected")
)

var (
	accessLoggerInstance internal.LogWriter = new(internal.NoOpLogWriter)
)

// InitAccessLogger initializes the access logger to write into the give file.
func InitAccessLogger(file string) error {
	logger, err := internal.NewFileLogWriter(file)
	if err != nil {
		Error("Failed to create access logger on file (", file, "): ", file, err)
		return err
	}
	accessLoggerInstance = logger
	return nil
}

// Access writes an access log.
func Access(from, to fmt.Stringer, status AccessStatus, reason fmt.Stringer) {
	accessLoggerInstance.Log(&internal.AccessLog{
		From:   from,
		To:     to,
		Status: string(status),
		Reason: reason,
	})
}
