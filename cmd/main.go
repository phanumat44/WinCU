package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"wincu/cleaner"
	"wincu/utils"
)

const Version = "1.1.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// scan command
	scanCmd := flag.NewFlagSet("scan", flag.ExitOnError)
	scanJson := scanCmd.Bool("json", false, "Output in JSON format")
	scanForce := scanCmd.Bool("force", false, "Force scan (try to elevate permissions)")

	// clean command
	cleanCmd := flag.NewFlagSet("clean", flag.ExitOnError)
	cleanAll := cleanCmd.Bool("all", false, "Clean all targets")
	cleanTemp := cleanCmd.Bool("temp", false, "Clean temporary files (User Temp, Windows Temp)")
	cleanPrefetch := cleanCmd.Bool("prefetch", false, "Clean Prefetch")
	cleanUpdate := cleanCmd.Bool("update", false, "Clean Windows Update Cache")
	cleanRecycle := cleanCmd.Bool("recyclebin", false, "Clean Recycle Bin")
	cleanChrome := cleanCmd.Bool("chrome", false, "Clean Google Chrome Cache")
	cleanEdge := cleanCmd.Bool("edge", false, "Clean Microsoft Edge Cache")
	cleanBrowser := cleanCmd.Bool("browser", false, "Clean all browser caches")
	cleanDryRun := cleanCmd.Bool("dry-run", false, "Simulate cleaning")
	cleanThreads := cleanCmd.Int("threads", runtime.NumCPU(), "Number of threads")
	cleanJson := cleanCmd.Bool("json", false, "Output in JSON format")
	cleanForce := cleanCmd.Bool("force", false, "Force delete (ignore errors)")
	versionFlag := flag.Bool("version", false, "Show version")

	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	if *versionFlag { // This won't work easily with subcommands logic unless we parse global flags first or check args
		// Simple check for version arg
	}
	// Actually, let's just check os.Args
	for _, arg := range os.Args {
		if arg == "--version" || arg == "-v" {
			fmt.Println("wincu version " + Version)
			os.Exit(0)
		}
	}

	switch os.Args[1] {
	case "scan":
		scanCmd.Parse(os.Args[2:])
		utils.InitLogger(*scanJson)

		// Check for force flag in scan
		// If force is used and not admin, try to elevate
		if *scanForce && !utils.IsAdmin() {
			utils.Info("Force flag detected. Attempting to elevate privileges...")
			err := utils.RunAsAdmin()
			if err != nil {
				utils.Error("Failed to elevate: " + err.Error())
				os.Exit(1)
			}
			os.Exit(0) // Exit original process
		}

		targets := cleaner.GetTargets()
		results := cleaner.ScanTargets(targets)

		var totalSize int64
		for _, res := range results {
			if *scanJson {
				utils.Info("Scan Result", res)
			} else {
				utils.Info(fmt.Sprintf("%-25s: %s (%d files)", res.Target, utils.FormatBytes(res.Size), res.Count))
			}
			totalSize += res.Size
		}
		if !*scanJson {
			utils.Info(fmt.Sprintf("Total reclaimable size: %s", utils.FormatBytes(totalSize)))
		}

	case "clean":
		cleanCmd.Parse(os.Args[2:])
		utils.InitLogger(*cleanJson)

		// If force is used and not admin, try to elevate
		if *cleanForce && !utils.IsAdmin() {
			utils.Info("Force flag detected. Attempting to elevate privileges...")
			err := utils.RunAsAdmin()
			if err != nil {
				utils.Error("Failed to elevate: " + err.Error())
				os.Exit(1)
			}
			os.Exit(0) // Exit original process
		}

		if !*cleanAll && !*cleanTemp && !*cleanRecycle && !*cleanPrefetch && !*cleanUpdate && !*cleanChrome && !*cleanEdge && !*cleanBrowser {
			fmt.Println("Please specify targets to clean (e.g., --all, --temp, --browser)")
			cleanCmd.PrintDefaults()
			os.Exit(1)
		}

		workerCount := *cleanThreads
		if workerCount < 1 {
			workerCount = 1
		}

		targets := cleaner.GetTargets()
		var selectedTargets []cleaner.Target

		if *cleanAll {
			selectedTargets = targets
		} else {
			for _, t := range targets {
				if *cleanTemp && (t.Name == "User Temp" || t.Name == "Windows Temp") {
					selectedTargets = append(selectedTargets, t)
				}
				if *cleanPrefetch && t.Name == "Prefetch" {
					selectedTargets = append(selectedTargets, t)
				}
				if *cleanUpdate && t.Name == "Windows Update Cache" {
					selectedTargets = append(selectedTargets, t)
				}
				if *cleanRecycle && t.Name == "Recycle Bin" {
					selectedTargets = append(selectedTargets, t)
				}
				if (*cleanChrome || *cleanBrowser) && t.Name == "Google Chrome Cache" {
					selectedTargets = append(selectedTargets, t)
				}
				if (*cleanEdge || *cleanBrowser) && t.Name == "Microsoft Edge Cache" {
					selectedTargets = append(selectedTargets, t)
				}
			}
		}

		// Pass *cleanForce to NewCleaner
		c := cleaner.NewCleaner(selectedTargets, *cleanDryRun, *cleanForce, workerCount)
		c.Run()

	default:
		printUsage()
		os.Exit(1)
	}

	// If we are running with force and not in JSON mode, pause so user can see output
	// (Useful when auto-elevated which opens a new window)
	if (*scanForce || *cleanForce) && (!*scanJson && !*cleanJson) {
		fmt.Println("\nPress Enter to exit...")
		fmt.Scanln()
	}
}

func printUsage() {
	fmt.Println("Usage: wincu <command> [flags]")
	fmt.Println("\nCommands:")
	fmt.Println("  scan   Scan for junk files")
	fmt.Println("  clean  Clean junk files")
	fmt.Println("\nGlobal Flags:")
	fmt.Println("  --help     Show help")
	fmt.Println("  --version  Show version")

	fmt.Println("\nScan Flags:")
	fmt.Println("  --json     Output in JSON format")
	fmt.Println("  --force    Force scan (elevate if needed)")

	fmt.Println("\nClean Flags:")
	fmt.Println("  --all          Clean all targets")
	fmt.Println("  --temp         Clean temporary files")
	fmt.Println("  --recyclebin   Clean Recycle Bin")
	fmt.Println("  --prefetch     Clean Prefetch")
	fmt.Println("  --update       Clean Windows Update Cache")
	fmt.Println("  --chrome       Clean Google Chrome Cache")
	fmt.Println("  --edge         Clean Microsoft Edge Cache")
	fmt.Println("  --browser      Clean all browser caches")
	fmt.Println("  --dry-run      Simulate cleaning")
	fmt.Println("  --force        Force delete (elevate if needed)")
	fmt.Println("  --threads <n>  Number of threads")
	fmt.Println("  --json         Output in JSON format")
}
