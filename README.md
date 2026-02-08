# WinCU — Windows Cleaner Utility

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go&logoColor=white)
![Platform](https://img.shields.io/badge/Platform-Windows-0078D6?logo=windows&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-green)
![Release](https://img.shields.io/github/v/release/phanumat44/wincu?include_prereleases)
![Downloads](https://img.shields.io/github/downloads/phanumat44/wincu/total)

🚀 **A fast, safe, and modern Windows cleaner CLI written in Go**

WinCU is a high-performance command-line utility for cleaning junk files on Windows.  
Designed with **speed**, **safety**, and **automation** in mind — perfect for power users, sysadmins, and scripts.

![WinCU Banner](assets/wincu.jpg)

---

## ✨ Features

### 🧹 Deep Cleaning

- User Temp
- Windows Temp
- Prefetch
- Windows Update Cache
- Recycle Bin

### ⚡ High Performance

- Concurrent worker pool
- Optimized I/O for large directories
- Scales with CPU threads

### 🛡️ Safety First

- **Dry-run mode** to preview changes
- Skips critical system paths automatically
- Graceful handling of permission errors

### 🔐 Auto-Elevation (UAC)

- Automatically requests Administrator privileges when required
- Triggered only when using `--force`

### 🤖 Automation Ready

- Optional **JSON output**
- Script-friendly CLI design

---

## 📦 Installation

### Option 1 — Installer (Recommended)

Download and run:

```
wincu_installer.exe
```

During installation:

- ✅ Select **“Add to PATH”** to use `wincu` globally

---

### Option 2 — Portable Executable

1. Download `wincu.exe`
2. Place it anywhere in your `PATH`

No installation required.

---

### Option 3 — Build from Source

**Requirements**

- Go ≥ 1.20
- Windows OS

```powershell
git clone https://github.com/phanumat44/wincu.git
cd wincu
go build -ldflags="-s -w" -o wincu.exe cmd/main.go
```

#### Build with Icon & Manifest (Optional)

```powershell
go install github.com/akavel/rsrc@latest
rsrc -manifest wincu.manifest -ico assets/wincu.ico -o rsrc.syso
go build -ldflags="-s -w" -o wincu.exe cmd/main.go
```

---

## 🚀 Usage

### Scan for Junk Files

```powershell
wincu scan
```

---

### Clean Selected Targets

```powershell
wincu clean --temp --recyclebin
```

---

### Dry-Run (Safe Preview)

```powershell
wincu clean --all --dry-run
```

---

### Force Clean (Admin Mode)

```powershell
wincu clean --all --force
```

---

## ⚙️ Command Flags

| Flag            | Description                              |
| --------------- | ---------------------------------------- |
| `--all`         | Clean all supported targets              |
| `--temp`        | Clean User & Windows Temp                |
| `--recyclebin`  | Empty Recycle Bin                        |
| `--prefetch`    | Clean Prefetch files                     |
| `--update`      | Clean Windows Update Cache               |
| `--chrome`      | Clean Google Chrome Cache                |
| `--edge`        | Clean Microsoft Edge Cache               |
| `--browser`     | Clean all browser caches                 |
| `--dry-run`     | Preview deletions without removing files |
| `--force`       | Force delete (Admin + read-only files)   |
| `--threads <n>` | Number of concurrent workers             |
| `--json`        | Output logs in JSON format               |
| `--version`     | Show version information                 |

---

## 🗂️ Project Structure

```
wincu/
├── cmd/
│   └── main.go
├── cleaner/
├── worker/
├── utils/
├── assets/
└── installer.iss
```

---

## 📄 License

MIT License © 2026
