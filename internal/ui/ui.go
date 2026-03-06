package ui

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

var (
	green     = color.New(color.FgGreen, color.Bold).SprintFunc()
	yellow    = color.New(color.FgYellow, color.Bold).SprintFunc()
	red       = color.New(color.FgRed, color.Bold).SprintFunc()
	cyan      = color.New(color.FgCyan, color.Bold).SprintFunc()
	blue      = color.New(color.FgBlue, color.Bold).SprintFunc()
	magenta   = color.New(color.FgMagenta, color.Bold).SprintFunc()
	white     = color.New(color.FgWhite, color.Bold).SprintFunc()
	gray      = color.New(color.FgHiBlack).SprintFunc()
	bold      = color.New(color.Bold).SprintFunc()
	underline = color.New(color.Underline).SprintFunc()
)

func getTerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	output, err := cmd.Output()
	if err != nil {
		return 80
	}
	
	width, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil || width < 40 {
		return 80
	}
	
	if width > 120 {
		return 120
	}
	
	return width
}

func Success(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", green("✓"), msg)
}

func Info(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", cyan("→"), msg)
}

func Warning(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", yellow("⚠"), msg)
}

func Error(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	fmt.Printf("%s %s\n", red("✗"), msg)
}

func Progress(current, total int, name string) {
	width := getTerminalWidth()
	barWidth := min(width-25, 30)
	percent := float64(current) / float64(total) * 100
	filled := int(percent * float64(barWidth) / 100)
	
	bar := strings.Repeat("█", filled)
	empty := strings.Repeat("░", barWidth-filled)
	
	fmt.Printf("\r%s [%s%s] %s %s", 
		blue("⚡"), 
		green(bar), 
		gray(empty), 
		white(fmt.Sprintf("%.0f%%", percent)),
		name)
}

func Header(text string) {
	width := getTerminalWidth()
	line := strings.Repeat("═", min(width, 60))
	
	fmt.Printf("\n%s\n", gray(line))
	fmt.Printf("%s %s\n", magenta("◆"), bold(underline(text)))
	fmt.Printf("%s\n", gray(line))
}

func Box(title, content string) {
	fmt.Printf("\n%s %s\n", blue("▶"), bold(title))
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		fmt.Printf("  %s\n", line)
	}
	fmt.Println()
}

func CategoryBox(category, count string, items []string) {
	fmt.Printf("\n%s %s %s\n", 
		Repository(category), 
		bold(category), 
		gray(fmt.Sprintf("(%s)", count)))
	
	width := getTerminalWidth()
	line := strings.Repeat("─", min(width-4, 50))
	fmt.Printf("  %s\n", gray(line))
	
	for i, item := range items {
		lines := strings.Split(item, "\n")
		for j, line := range lines {
			if j == 0 {
				fmt.Printf("  %s %s\n", cyan("•"), line)
			} else {
				fmt.Printf("    %s\n", gray(line))
			}
		}
		if i < len(items)-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func InstallPlanBox(repoPackages, communityPackages []string) {
	fmt.Printf("\n%s %s\n", blue("▶"), bold("KURULUM PLANI"))
	
	width := getTerminalWidth()
	line := strings.Repeat("═", min(width-4, 50))
	fmt.Printf("  %s\n", gray(line))
	
	if len(repoPackages) > 0 {
		fmt.Printf("  %s %s %s\n", 
			Repository("REPO"), 
			bold("Resmi Repo"), 
			gray(fmt.Sprintf("(%d paket)", len(repoPackages))))
		
		for _, pkg := range repoPackages {
			fmt.Printf("    %s %s\n", cyan("•"), Bold(pkg))
		}
		fmt.Println()
	}
	
	if len(communityPackages) > 0 {
		fmt.Printf("  %s %s %s\n", 
			Repository("TOPLULUK"), 
			bold("Topluluk"), 
			gray(fmt.Sprintf("(%d paket)", len(communityPackages))))
		
		for _, pkg := range communityPackages {
			fmt.Printf("    %s %s\n", cyan("•"), Bold(pkg))
		}
		fmt.Println()
	}
	
	totalText := fmt.Sprintf("Toplam: %d paket", len(repoPackages)+len(communityPackages))
	fmt.Printf("  %s %s\n", cyan("→"), bold(totalText))
	fmt.Println()
}

func Separator() {
	width := getTerminalWidth()
	fmt.Printf("%s\n", gray(strings.Repeat("─", min(width, 60))))
}

func Banner(text string) {
	width := getTerminalWidth()
	line := strings.Repeat("═", min(width, 60))
	
	fmt.Printf("\n%s\n", blue(line))
	fmt.Printf("%s %s\n", magenta("◆"), bold(text))
	fmt.Printf("%s\n", blue(line))
}

func SimpleList(items []string) {
	for _, item := range items {
		fmt.Printf("  %s %s\n", cyan("•"), item)
	}
}

func NumberedList(items []string) {
	for i, item := range items {
		fmt.Printf("  %s %s\n", cyan(fmt.Sprintf("%d.", i+1)), item)
	}
}

func Section(title string) {
	fmt.Printf("\n%s %s\n", magenta("▶"), bold(title))
}

func Indent(level int, format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	indent := strings.Repeat("  ", level)
	fmt.Printf("%s%s\n", indent, msg)
}

func Bold(text string) string {
	return bold(text)
}

func Highlight(text string) string {
	return cyan(text)
}

func Package(name, version string) string {
	return fmt.Sprintf("%s %s", bold(name), gray(version))
}

func Repository(repo string) string {
	switch strings.ToLower(repo) {
	case "core":
		return blue("CORE")
	case "extra":
		return green("EXTRA")
	case "community":
		return yellow("COMMUNITY")
	case "multilib":
		return magenta("MULTILIB")
	case "aur", "topluluk":
		return magenta("TOPLULUK")
	case "repo":
		return blue("REPO")
	default:
		return cyan(strings.ToUpper(repo))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}