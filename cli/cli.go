/*
 * gomp3 - CLI utilities for logging and usage display
 * Copyright (C) 2025  Juan Perdomo
 * Licensed under GNU GPL v3 or later
 */

package cli

import (
	"fmt"
	"os"
)

// Usage shows how to use the program and exits
func Usage() {
	LogWarning("Usage: gomp3 file[urls] [dir]")
	fmt.Fprintf(os.Stderr, "")
	os.Exit(1)
}

