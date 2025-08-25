// 代码生成时间: 2025-08-25 15:48:51
package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
    "strings"
    "time"
)

// LogEntry represents a single log entry with timestamp, level, and message
type LogEntry struct {
    Timestamp time.Time
    Level     string
    Message   string
}

// parseLogLine takes a single log line as input and returns a LogEntry struct
func parseLogLine(line string) (*LogEntry, error) {
    // Assume the log line format is: [timestamp] level: message
    // For example: [2023-04-01 12:00:00] INFO: This is a log message
    parts := strings.Fields(line)
    if len(parts) < 3 || parts[0] != "[" || parts[2] != "]": {
        return nil, fmt.Errorf("invalid log line format")
    }

    timestampStr := parts[1] + parts[2]
    timestamp, err := time.Parse(`2006-01-02 15:04:05`, timestampStr)
    if err != nil {
        return nil, err
    }

    level := parts[3]
    message := strings.Join(parts[4:], " ")

    return &LogEntry{Timestamp: timestamp, Level: level, Message: message}, nil
}

// ParseLogFile reads the log file and prints each parsed log entry
func ParseLogFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        entry, err := parseLogLine(line)
        if err != nil {
            log.Printf("Error parsing log line: %s, error: %s", line, err)
            continue
        }
        fmt.Printf("Timestamp: %s, Level: %s, Message: %s
", entry.Timestamp, entry.Level, entry.Message)
    }

    if err := scanner.Err(); err != nil {
        return err
    }
    return nil
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: log_parser <logfile>")
        os.Exit(1)
    }

    filename := os.Args[1]
    if err := ParseLogFile(filename); err != nil {
        fmt.Printf("Error reading log file: %s
", err)
    }
}