package ui

import (
	"fmt"
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
	percent := float64(current) / float64(total) * 100
	bar := strings.Repeat("█", int(percent/5))
	spaces := strings.Repeat("░", 20-len(bar))
	fmt.Printf("\r%s [%s%s] %s %s", 
		blue("⚡"), 
		green(bar), 
		gray(spaces), 
		white(fmt.Sprintf("%.0f%%", percent)),
		name)
}

func Header(text string) {
	fmt.Printf("\n%s %s\n", magenta("◆"), underline(text))
}

func Separator() {
	fmt.Println(gray("─────────────────────────────────────────"))
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
	switch repo {
	case "core":
		return blue(repo)
	case "extra":
		return green(repo)
	case "community":
		return yellow(repo)
	case "aur":
		return magenta(repo)
	default:
		return cyan(repo)
	}
}
