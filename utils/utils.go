/*
 * gomp3 - Utility functions for file reading and URL processing
 * Copyright (C) 2025  Juan Perdomo
 * Licensed under GNU GPL v3 or later
 */

package utils

import (
	"bufio"
	"gomp3/cli"
	"os"
	"strings"
)

type Link struct {
	URL string
}

// ReadFile reads a file containing URLs and returns a slice of Link
func ReadFile(filename string) ([]Link, error) {
	file, err := os.Open(filename)
	if err != nil {
		cli.LogError("Error to read file")
		return nil, err
	}
	defer file.Close()

	var links []Link
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			links = append(links, Link{
				URL: line,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		cli.LogError("Error to read file")
		return nil, err
	}

	return links, nil
}

// RemovePlaylistParam removes "list=" parameter from YouTube URLs
func RemovePlaylistParam(url string) string {
	if strings.Contains(url, "list=") {
		parts := strings.Split(url, "&")
		var filtered []string
		for _, p := range parts {
			if !strings.HasPrefix(p, "list=") {
				filtered = append(filtered, p)
			}
		}
		return strings.Join(filtered, "&")
	}
	return url
}

