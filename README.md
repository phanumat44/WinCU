# WinCU (Windows Cleaner Utility)

🚀 **A fast Windows cleaner CLI written in Go**  
Supports cache, temp, recycle bin, Windows Update cleanup.  
**Multi-threaded & safe by default.**

![wincu](assets/wincu.jpg)

## Features

- **Deep Cleaning**: Removes User Temp, Windows Temp, Prefetch, Windows Update Cache, and empties the Recycle Bin.
- **Concurrency**: Uses a worker pool to delete files in parallel for maximum speed.
- **Safety First**:
  - Dry-run mode to preview changes.
  - Skips critical system files.
  - Handles permission errors gracefully.
- **Auto-Elevation**: Automatically requests Administrator privileges (UAC) when needed (e.g., using `--force`).
- **Automation Ready**: Supports JSON output for easy integration with scripts.

## Installation

### Option 1: Installer (Recommended)

Download and run the installer `wincu_installer.exe`.

- Select **"Add to PATH"** during installation to run `wincu` from any terminal.

### Option 2: Portable Executable

Download `wincu.exe` and place it anywhere in your PATH.

### Option 3: Build from Source

Requires [Go](https://go.dev/dl/) installed.

```powershell
git clone https://github.com/yourusername/wincu.git
cd wincu
go build -ldflags="-s -w" -o wincu.exe cmd/main.go
```

To embed the icon (requires `rsrc`):

```powershell
go install github.com/akavel/rsrc@latest
rsrc -manifest wincu.manifest -ico assets/wincu.ico -o rsrc.syso
go build -ldflags="-s -w" -o wincu.exe cmd/main.go
```

## Usage

### Commands

**Scan for junk:**

```powershell
wincu scan
```

**Clean specific targets:**

```powershell
wincu clean --temp --recyclebin
```

_(Cleans User Temp, Windows Temp, and Recycle Bin)_

**Preview deletion (Safe Mode):**

```powershell
wincu clean --all --dry-run
```

**Force Clean (Admin + Read-only files):**

```powershell
wincu clean --all --force
```

_(Triggers UAC prompt if not already Admin)_

### Flags

| Flag            | Description                                          |
| :-------------- | :--------------------------------------------------- |
| `--all`         | Clean all supported targets                          |
| `--temp`        | Clean User & Windows Temp folders                    |
| `--recyclebin`  | Empty Recycle Bin                                    |
| `--prefetch`    | Clean Prefetch files                                 |
| `--update`      | Clean Windows Update Cache                           |
| `--dry-run`     | Simulate cleaning without deleting                   |
| `--force`       | Force delete (elevate permissions, delete read-only) |
| `--threads <n>` | Set number of concurrent threads                     |
| `--json`        | Output logs in JSON format                           |
| `--version`     | Show version info                                    |

## Project Structure

- `cmd/main.go`: Application entry point.
- `cleaner/`: Core logic for target scanning and deletion.
- `worker/`: Worker pool for concurrent processing.
- `utils/`: Helper utilities (logging, UAC, etc.).
- `assets/`: Icons and images.

## License

MIT License
