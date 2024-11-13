/*
The aim of this program is to use slog package to write logs to a log file in JSON format.
We did this to be able to retrieve the logs for later use, since they are in format that we can decode using encoding/json package.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

type LogEntry struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Message string `json:"msg"`
}

func main() {
	// Open the log file for writing
	logFile, err := os.OpenFile("Logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("Error opening log file: %v\n", err)
		return
	}

	// Set up JSON logging to the file

	// We want to change the level to debug, so that debug logs will also appear in log records

	logLevel := &slog.LevelVar{}
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}

	logJSON := slog.New(slog.NewJSONHandler(logFile, opts)) // we can also set (os.Stderr, opts) to output only to stderr
	logLevel.Set(slog.LevelDebug)

	logJSON.Error("Error message in JSON.")
	logJSON.Debug("Debug message in JSON.")
	logJSON.Info("Info message in JSON.")
	logJSON.Warn("Warning message in JSON.")

	// Close the file after writing
	// We need to close the file before reopening it to process the log files
	logFile.Close()

	// Reopen the log file for reading
	logFile, err = os.Open("Logs.log")
	if err != nil {
		fmt.Printf("Error reopening log file: %v\n", err)
		return
	}
	defer logFile.Close()

	// Read and decode the JSON logs from the file
	decoder := json.NewDecoder(logFile)
	var logEntries []LogEntry

	for decoder.More() {
		var entry LogEntry
		err := decoder.Decode(&entry)
		if err != nil {
			fmt.Printf("Error decoding log entry: %v\n", err)
			return
		}
		logEntries = append(logEntries, entry)
	}

	// Use the decoded log entries
	for _, entry := range logEntries {
		fmt.Printf("Time: %s, Level: %s, Message: %s\n", entry.Time, entry.Level, entry.Message)
	}
}
