/*
 * gomp3 - ANSI colors for terminal output
 * Copyright (C) 2025  Juan Perdomo
 * Licensed under GNU GPL v3 or later
 */

package cli

import "fmt"

const (
	Reset = "\033[0m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
)

// Print returns a string with ANSI color applied
func Print(color string, text string) string {
	return fmt.Sprintf("%s%s%s", color, text, Reset)
}

