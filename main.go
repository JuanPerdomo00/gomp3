/*
 * gomp3 - Download YouTube videos as MP3
 * Copyright (C) 2025  Juan Perdomo
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package main

import (
	"gomp3/cli"
	"gomp3/dowloader"
	"gomp3/utils"
	"os"
)

func main() {
	// Check arguments
	args := os.Args
	if len(args) != 3 {
		cli.Usage()
		os.Exit(1)
	}

	file := args[1]
	dir := args[2]

	// Read links from file
	links, err := utils.ReadFile(file)
	if err != nil {
		cli.LogError("Error reading file: " + err.Error())
		os.Exit(1)
	}

	// Create output directory if not exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		cli.LogError("Error creating directory: " + err.Error())
		os.Exit(1)
	}

	// Process each link
	for _, link := range links {
		cleanURL := utils.RemovePlaylistParam(link.URL)
		if err := dowloader.DownloadMP3(cleanURL, dir); err != nil {
			cli.LogError("Failed to download: " + cleanURL)
		} else {
			cli.LogSuccess("Downloaded: " + cleanURL)
		}
	}
}
