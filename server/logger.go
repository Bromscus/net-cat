package server

import (
	"io"
	"log"
	"os"
)

func Logger() (*os.File, error) {
	logFile, err := os.OpenFile(
		"server.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		return nil, err
	}

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))
	return logFile, nil
}
