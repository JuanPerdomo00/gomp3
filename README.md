# 🎵 gomp3 - YouTube MP3 Downloader

**Author:** Jakepys  
**License:** [GNU GPL v3](LICENSE)  
**Version:** 1.0  

GoMP3 is a **command-line tool** written in Go that downloads audio from YouTube in **MP3 format**.  
It reads a text file containing one or more YouTube video or playlist links and downloads the audio to a specified folder.  

Powered by [yt-dlp](https://github.com/yt-dlp/yt-dlp) + [FFmpeg](https://ffmpeg.org/).


## ⚙️ Requirements

Before using GoMP3, make sure you have:

- **Go** (1.20+ recommended) → [Download Go](https://go.dev/dl/)  
- **yt-dlp** → [yt-dlp GitHub](https://github.com/yt-dlp/yt-dlp)  
- **FFmpeg** → [FFmpeg Downloads](https://ffmpeg.org/download.html)  

---

## 📥 Installation

### 🐧 Linux

| Distro       | Install Command |
|--------------|----------------|
| **Arch Linux** | `sudo pacman -S go yt-dlp ffmpeg` |
| **Fedora**     | `sudo dnf install golang yt-dlp ffmpeg` |
| **Ubuntu/Debian** | `sudo apt update && sudo apt install golang yt-dlp ffmpeg` |

---

### 🪟 Windows

1. Install Go → [Download Go](https://go.dev/dl/)  
2. Install yt-dlp:  
   - Download from [yt-dlp releases](https://github.com/yt-dlp/yt-dlp/releases/latest)  
   - Place `yt-dlp.exe` in a folder and add it to your **PATH**.  
3. Install FFmpeg:  
   - Download from [FFmpeg official site](https://ffmpeg.org/download.html)  
   - Extract and add the `bin` folder to your **PATH**.  

---

## 🔨 Compilation

Clone the repository:
```bash
git clone https://github.com/usuario/gomp3.git
cd gomp3
```

Compile for your system:
```bash
go build -o gomp3 main.go
```

Compile with a custom name:
```bash
go build -o mydownloader main.go
```

---

## Usage

Syntax:
```bash
./gomp3 file_with_links.txt output_directory
```

Example:
```bash
./gomp3 links.txt ./music
```

Where:
- **`links.txt`** → Text file with YouTube URLs (one per line)  
- **`./music`** → Output folder for downloaded MP3 files  

Example `links.txt`:
```
https://www.youtube.com/watch?v=R_KkeckyFsI
https://www.youtube.com/watch?v=dQw4w9WgXcQ
```

---

## Notes

- If you pass a **playlist URL**, yt-dlp will download **all videos** in the playlist.  
- If you only want a single song from a playlist link, **use the direct video URL**.  

Example:  
Playlist link →  
```
https://www.youtube.com/watch?v=R_KkeckyFsI&list=PL1234
```
 Single video link →  
```
https://www.youtube.com/watch?v=R_KkeckyFsI
```

- Make sure `yt-dlp` and `ffmpeg` are **in your PATH**.

---

## License

This project is licensed under the **GNU General Public License v3.0 (GPLv3)**.  
You are free to use, modify, and distribute it, but **derivative works must also be GPLv3**.

---

