WinCU (Windows Cleaner Utility) v1.0
====================================

🚀 A fast Windows cleaner CLI written in Go
Supports cache, temp, recycle bin, Windows Update cleanup.
Multi-threaded & safe by default.

GETTING STARTED
---------------
1. Run wincu.exe from the command line (PowerShell or CMD).
2. For best results, run as Administrator to clean system caches.

COMMANDS
--------

1. Scan for junk files (check reclaimable space):
   wincu.exe scan

2. Preview what would be deleted (Dry Run):
   wincu.exe clean --temp --dry-run

3. Clean specific targets:
   wincu.exe clean --temp          (Cleans User & Windows Temp)
   wincu.exe clean --recyclebin    (Empties Recycle Bin)

4. Clean EVERYTHING (Requires Admin):
   wincu.exe clean --all --force

Opitons:
  --json       Output logs in JSON format
  --threads N  Set number of concurrent threads (default: CPU cores)
  --version    Show version info

TROUBLESHOOTING
---------------
- "Permission denied": Run terminal as Administrator.
- "Access is denied": Use --force flag to delete read-only files.

Enjoy a cleaner Windows!
