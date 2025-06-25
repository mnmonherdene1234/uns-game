package gameengine

import "os"

type Logger struct {
	LogFilePath string // Name of the log file
}

func NewLogger(logFilepath string) *Logger {
	return &Logger{
		LogFilePath: logFilepath,
	}
}

func (l *Logger) FilePrintln(message string) {
	if l.LogFilePath == "" {
		return // If no log file path is set, do nothing
	}

	// check file exists
	if _, err := os.Stat(l.LogFilePath); os.IsNotExist(err) {
		// If the file does not exist, create it
		file, err := os.Create(l.LogFilePath)
		if err != nil {
			println("Failed to create log file:", err.Error())
			return
		}
		defer file.Close()
	}

	// Open the log file in append mode
	file, err := os.OpenFile(l.LogFilePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		println("Failed to open log file:", err.Error())
		return
	}

	file.WriteString(message + "\n") // Write the message to the log file
	if err := file.Sync(); err != nil {
		println("Failed to sync log file:", err.Error())
		return
	}

	defer file.Close()
}

func (l *Logger) Info(message string) {
	// Implement logging logic here, e.g., print to console or write to a file
	log := "[INFO]" + message
	println(log)
	l.FilePrintln(log)
}

func (l *Logger) Error(message string) {
	// Implement error logging logic here, e.g., print to console or write to a file
	log := "[ERROR]" + message
	println(log)
	l.FilePrintln(log)
}
