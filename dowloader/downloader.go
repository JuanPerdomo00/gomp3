/*
 * gomp3 - Downloader using yt-dlp to extract MP3 audio
 * Copyright (C) 2025  Juan Perdomo
 * Licensed under GNU GPL v3 or later
 */

package dowloader

import (
	"fmt"
	"os"
	"os/exec"
)

// DownloadMP3 downloads audio in MP3 format from a given YouTube URL
func DownloadMP3(url, outputDir string) error {
	cmd := exec.Command(
		"yt-dlp",
		"--no-playlist", 
		"-x", "--audio-format", "mp3",
		"-o", fmt.Sprintf("%s/%%(title)s.%%(ext)s", outputDir),
		url,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

