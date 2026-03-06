package system

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type OSInfo struct {
	Type         string // linux, darwin, windows
	Distribution string // ubuntu, fedora, arch, etc.
	Version      string
	PackageManager string
}

type PackageManager struct {
	Name        string
	InstallCmd  []string
	RemoveCmd   []string
	UpdateCmd   []string
	SearchCmd   []string
	ListCmd     []string
	InfoCmd     []string
	SudoNeeded  bool
}

var packageManagers = map[string]PackageManager{
	"pacman": {
		Name:       "pacman",
		InstallCmd: []string{"pacman", "-S"},
		RemoveCmd:  []string{"pacman", "-Rs"},
		UpdateCmd:  []string{"pacman", "-Syu"},
		SearchCmd:  []string{"pacman", "-Ss"},
		ListCmd:    []string{"pacman", "-Q"},
		InfoCmd:    []string{"pacman", "-Si"},
		SudoNeeded: true,
	},
	"apt": {
		Name:       "apt",
		InstallCmd: []string{"apt", "install"},
		RemoveCmd:  []string{"apt", "remove"},
		UpdateCmd:  []string{"apt", "update", "&&", "apt", "upgrade"},
		SearchCmd:  []string{"apt", "search"},
		ListCmd:    []string{"apt", "list", "--installed"},
		InfoCmd:    []string{"apt", "show"},
		SudoNeeded: true,
	},
	"dnf": {
		Name:       "dnf",
		InstallCmd: []string{"dnf", "install"},
		RemoveCmd:  []string{"dnf", "remove"},
		UpdateCmd:  []string{"dnf", "upgrade"},
		SearchCmd:  []string{"dnf", "search"},
		ListCmd:    []string{"dnf", "list", "installed"},
		InfoCmd:    []string{"dnf", "info"},
		SudoNeeded: true,
	},
	"yum": {
		Name:       "yum",
		InstallCmd: []string{"yum", "install"},
		RemoveCmd:  []string{"yum", "remove"},
		UpdateCmd:  []string{"yum", "update"},
		SearchCmd:  []string{"yum", "search"},
		ListCmd:    []string{"yum", "list", "installed"},
		InfoCmd:    []string{"yum", "info"},
		SudoNeeded: true,
	},
	"zypper": {
		Name:       "zypper",
		InstallCmd: []string{"zypper", "install"},
		RemoveCmd:  []string{"zypper", "remove"},
		UpdateCmd:  []string{"zypper", "update"},
		SearchCmd:  []string{"zypper", "search"},
		ListCmd:    []string{"zypper", "search", "--installed-only"},
		InfoCmd:    []string{"zypper", "info"},
		SudoNeeded: true,
	},
	"apk": {
		Name:       "apk",
		InstallCmd: []string{"apk", "add"},
		RemoveCmd:  []string{"apk", "del"},
		UpdateCmd:  []string{"apk", "upgrade"},
		SearchCmd:  []string{"apk", "search"},
		ListCmd:    []string{"apk", "info"},
		InfoCmd:    []string{"apk", "info"},
		SudoNeeded: true,
	},
	"brew": {
		Name:       "brew",
		InstallCmd: []string{"brew", "install"},
		RemoveCmd:  []string{"brew", "uninstall"},
		UpdateCmd:  []string{"brew", "upgrade"},
		SearchCmd:  []string{"brew", "search"},
		ListCmd:    []string{"brew", "list"},
		InfoCmd:    []string{"brew", "info"},
		SudoNeeded: false,
	},
	"port": {
		Name:       "port",
		InstallCmd: []string{"port", "install"},
		RemoveCmd:  []string{"port", "uninstall"},
		UpdateCmd:  []string{"port", "upgrade", "outdated"},
		SearchCmd:  []string{"port", "search"},
		ListCmd:    []string{"port", "installed"},
		InfoCmd:    []string{"port", "info"},
		SudoNeeded: true,
	},
	"choco": {
		Name:       "choco",
		InstallCmd: []string{"choco", "install"},
		RemoveCmd:  []string{"choco", "uninstall"},
		UpdateCmd:  []string{"choco", "upgrade", "all"},
		SearchCmd:  []string{"choco", "search"},
		ListCmd:    []string{"choco", "list", "--local-only"},
		InfoCmd:    []string{"choco", "info"},
		SudoNeeded: false,
	},
	"scoop": {
		Name:       "scoop",
		InstallCmd: []string{"scoop", "install"},
		RemoveCmd:  []string{"scoop", "uninstall"},
		UpdateCmd:  []string{"scoop", "update", "*"},
		SearchCmd:  []string{"scoop", "search"},
		ListCmd:    []string{"scoop", "list"},
		InfoCmd:    []string{"scoop", "info"},
		SudoNeeded: false,
	},
}

func DetectOS() (*OSInfo, error) {
	osInfo := &OSInfo{
		Type: runtime.GOOS,
	}
	
	switch runtime.GOOS {
	case "linux":
		return detectLinuxDistro(osInfo)
	case "darwin":
		return detectMacOS(osInfo)
	case "windows":
		return detectWindows(osInfo)
	default:
		return nil, fmt.Errorf("desteklenmeyen işletim sistemi: %s", runtime.GOOS)
	}
}

func detectLinuxDistro(osInfo *OSInfo) (*OSInfo, error) {
	if data, err := os.ReadFile("/etc/os-release"); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "ID=") {
				osInfo.Distribution = strings.Trim(strings.TrimPrefix(line, "ID="), "\"")
			}
			if strings.HasPrefix(line, "VERSION_ID=") {
				osInfo.Version = strings.Trim(strings.TrimPrefix(line, "VERSION_ID="), "\"")
			}
		}
	}
	
	osInfo.PackageManager = detectPackageManager()
	
	return osInfo, nil
}

func detectMacOS(osInfo *OSInfo) (*OSInfo, error) {
	osInfo.Distribution = "macos"
	
	if output, err := exec.Command("sw_vers", "-productVersion").Output(); err == nil {
		osInfo.Version = strings.TrimSpace(string(output))
	}
	
	if commandExists("brew") {
		osInfo.PackageManager = "brew"
	} else if commandExists("port") {
		osInfo.PackageManager = "port"
	} else {
		osInfo.PackageManager = "brew" // varsayılan
	}
	
	return osInfo, nil
}

func detectWindows(osInfo *OSInfo) (*OSInfo, error) {
	osInfo.Distribution = "windows"
	
	if output, err := exec.Command("cmd", "/c", "ver").Output(); err == nil {
		osInfo.Version = strings.TrimSpace(string(output))
	}
	
	if commandExists("choco") {
		osInfo.PackageManager = "choco"
	} else if commandExists("scoop") {
		osInfo.PackageManager = "scoop"
	} else {
		osInfo.PackageManager = "choco" // varsayılan
	}
	
	return osInfo, nil
}

func detectPackageManager() string {
	managers := []string{"pacman", "apt", "dnf", "yum", "zypper", "apk"}
	
	for _, manager := range managers {
		if commandExists(manager) {
			return manager
		}
	}
	
	return "unknown"
}

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func GetPackageManager(name string) (PackageManager, bool) {
	pm, exists := packageManagers[name]
	return pm, exists
}

func (osInfo *OSInfo) GetDisplayName() string {
	switch osInfo.Distribution {
	case "ubuntu":
		return "Ubuntu Linux"
	case "debian":
		return "Debian Linux"
	case "fedora":
		return "Fedora Linux"
	case "centos":
		return "CentOS Linux"
	case "rhel":
		return "Red Hat Enterprise Linux"
	case "arch":
		return "Arch Linux"
	case "manjaro":
		return "Manjaro Linux"
	case "opensuse", "opensuse-leap", "opensuse-tumbleweed":
		return "openSUSE Linux"
	case "alpine":
		return "Alpine Linux"
	case "macos":
		return "macOS"
	case "windows":
		return "Microsoft Windows"
	default:
		return fmt.Sprintf("%s (%s)", strings.Title(osInfo.Distribution), osInfo.Type)
	}
}

func (osInfo *OSInfo) SupportsCommunityRepos() bool {
	return osInfo.Distribution == "arch" || osInfo.Distribution == "manjaro"
}