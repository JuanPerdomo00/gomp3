/*
 * gomp3 - Logging functions with timestamps and colors
 * Copyright (C) 2025  Juan Perdomo
 * Licensed under GNU GPL v3 or later
 */

package cli

import (
	"fmt"
	"time"
)

// LogWarning shows a warning message
func LogWarning(text string) {
	logWithTime(BrightYellow, "WARNING", text)
}

// LogInfo shows an info message
func LogInfo(text string) {
	logWithTime(BrightBlue, "INFO", text)
}

// LogError shows an error message
func LogError(text string) {
	logWithTime(BrightRed, "ERROR", text)
}

// LogSuccess shows a success message
func LogSuccess(text string) {
	logWithTime(BrightGreen, "SUCCESS", text)
}

// logWithTime prints log with timestamp and color
func logWithTime(color, label, text string) {
	time := time.Now().Format("15:04:05")
	fmt.Println(Print(color, fmt.Sprintf("[%s] [%s] --> %s", time, label, text)))
}

