# Release Process

Follow these steps to release a new version of **WinCU**.

## 1. Prerequisites

- **Go** (1.20+)
- **Inno Setup** (for installer)
- **Git**

## 2. Update Version

Run the automation script to update the version number in all files (`main.go`, `manifest`, `installer`, `CHANGELOG`).

```powershell
# Example: Update to version 1.1.0
go run scripts/update_version.go 1.1.0
```

This script will automatically set the installer output filename to: `wincu-installer-v1.1.0.exe`.

## 3. Commit Changes

Commit the version bump changes.

```powershell
git add .
git commit -m "chore: release v1.1.0"
git tag v1.1.0
git push origin v1.1.0
```

## 4. Build Executable (Binary)

Build the standalone executable with the requested naming pattern: `wincu-v<version>-windows-amd64.exe`.

```powershell
# Replace 1.1.0 with your actual version
go build -ldflags="-s -w" -o wincu-v1.1.0-windows-amd64.exe cmd/main.go
```

## 5. Build Installer

Compile the `installer.iss` script using Inno Setup.

1.  Open `installer.iss` with Inno Setup Compiler.
2.  Click **Build** > **Compile**.
3.  The output file will be generated in the project root as `wincu-installer-v1.1.0.exe`.

## 6. Publish to GitHub

1.  Go to the [GitHub Releases](https://github.com/phanumat44/wincu/releases) page.
2.  Draft a new release.
3.  Select the tag `v1.1.0`.
4.  Copy the notes from `CHANGELOG.md`.
5.  **Upload Assets**:
    - `wincu-v1.1.0-windows-amd64.exe`
    - `wincu-installer-v1.1.0.exe`
6.  Publish!
