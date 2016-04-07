package docker

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type JSONLog struct {
	Log     string    `json:"log,omitempty"`
	Stream  string    `json:"stream,omitempty"`
	Created time.Time `json:"time"`
}

// AddContainerLog appends a single line of log to container's logfile
// For this version of daoker, only json-file is supported
// TODO: support more log drivers.
func AddContainerLog(logPath string, content string) error {
	if logPath == "" {
		return fmt.Errorf("logPath is empty, log driver for this container is not json-file")
	}

	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}

	defer f.Close()

	log := JSONLog{
		Log:     content,
		Stream:  "stdout",
		Created: time.Now().UTC(),
	}

	jsonLog, err := json.Marshal(log)
	context := string(jsonLog) + "\n"

	// Start to append container log.
	if _, err = f.WriteString(context); err != nil {
		return err
	}

	return nil
}
