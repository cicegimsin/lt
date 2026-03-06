package universal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/cicegimsin/lt/internal/system"
)

type UniversalManager struct {
	OSInfo         *system.OSInfo
	PackageManager system.PackageManager
	SudoPath       string
}

type Package struct {
	Name        string
	Version     string
	Description string
	Repository  string
	Source      string
}

func NewUniversalManager() (*UniversalManager, error) {
	osInfo, err := system.DetectOS()
	if err != nil {
		return nil, err
	}
	
	pm, exists := system.GetPackageManager(osInfo.PackageManager)
	if !exists {
		return nil, fmt.Errorf("desteklenmeyen paket yöneticisi: %s", osInfo.PackageManager)
	}
	
	sudoPath := "/usr/bin/sudo"
	if osInfo.Type == "windows" || !pm.SudoNeeded {
		sudoPath = ""
	}
	
	return &UniversalManager{
		OSInfo:         osInfo,
		PackageManager: pm,
		SudoPath:       sudoPath,
	}, nil
}

func (um *UniversalManager) Search(query string) ([]Package, error) {
	var allPackages []Package
	
	// Resmi paket yöneticisinden ara
	officialPackages, err := um.searchOfficial(query)
	if err == nil {
		allPackages = append(allPackages, officialPackages...)
	}
	
	if um.OSInfo.SupportsCommunityRepos() {
		communityPackages, err := um.searchCommunity(query)
		if err == nil {
			allPackages = append(allPackages, communityPackages...)
		}
	}
	
	return allPackages, nil
}

func (um *UniversalManager) searchOfficial(query string) ([]Package, error) {
	cmd := um.PackageManager.SearchCmd
	args := append(cmd[1:], query)
	
	execCmd := exec.Command(cmd[0], args...)
	output, err := execCmd.Output()
	if err != nil {
		return nil, err
	}
	
	return um.parseSearchOutput(string(output), "official"), nil
}

func (um *UniversalManager) searchCommunity(query string) ([]Package, error) {
	return []Package{}, nil
}

func (um *UniversalManager) parseSearchOutput(output, source string) []Package {
	var packages []Package
	lines := strings.Split(output, "\n")
	
	switch um.PackageManager.Name {
	case "pacman":
		packages = um.parsePacmanOutput(lines, source)
	case "apt":
		packages = um.parseAptOutput(lines, source)
	case "dnf", "yum":
		packages = um.parseDnfOutput(lines, source)
	case "brew":
		packages = um.parseBrewOutput(lines, source)
	default:
		packages = um.parseGenericOutput(lines, source)
	}
	
	return packages
}

func (um *UniversalManager) parsePacmanOutput(lines []string, source string) []Package {
	var packages []Package
	
	for i := 0; i < len(lines); i++ {
		line := strings.TrimSpace(lines[i])
		if line == "" {
			continue
		}
		
		if !strings.HasPrefix(line, " ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				repoPkg := strings.Split(parts[0], "/")
				if len(repoPkg) == 2 {
					pkg := Package{
						Name:       repoPkg[1],
						Version:    parts[1],
						Repository: repoPkg[0],
						Source:     source,
					}
					
					if i+1 < len(lines) {
						desc := strings.TrimSpace(lines[i+1])
						if strings.HasPrefix(desc, " ") {
							pkg.Description = strings.TrimSpace(desc)
						}
					}
					
					packages = append(packages, pkg)
				}
			}
		}
	}
	
	return packages
}

func (um *UniversalManager) parseAptOutput(lines []string, source string) []Package {
	var packages []Package
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "WARNING") {
			continue
		}
		
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			name := parts[0]
			if strings.Contains(name, "/") {
				name = strings.Split(name, "/")[0]
			}
			
			version := parts[1]
			description := ""
			if len(parts) > 2 {
				description = strings.Join(parts[2:], " ")
			}
			
			packages = append(packages, Package{
				Name:        name,
				Version:     version,
				Description: description,
				Source:      source,
			})
		}
	}
	
	return packages
}

func (um *UniversalManager) parseDnfOutput(lines []string, source string) []Package {
	var packages []Package
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "=") || strings.Contains(line, "Matched:") {
			continue
		}
		
		parts := strings.Fields(line)
		if len(parts) >= 3 {
			name := parts[0]
			if strings.Contains(name, ".") {
				name = strings.Split(name, ".")[0]
			}
			
			packages = append(packages, Package{
				Name:        name,
				Version:     parts[1],
				Description: strings.Join(parts[2:], " "),
				Source:      source,
			})
		}
	}
	
	return packages
}

func (um *UniversalManager) parseBrewOutput(lines []string, source string) []Package {
	var packages []Package
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		parts := strings.Fields(line)
		if len(parts) >= 1 {
			packages = append(packages, Package{
				Name:   parts[0],
				Source: source,
			})
		}
	}
	
	return packages
}

func (um *UniversalManager) parseGenericOutput(lines []string, source string) []Package {
	var packages []Package
	
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		
		parts := strings.Fields(line)
		if len(parts) >= 1 {
			packages = append(packages, Package{
				Name:   parts[0],
				Source: source,
			})
		}
	}
	
	return packages
}

func (um *UniversalManager) Install(pkgNames []string, noConfirm bool) error {
	cmd := um.PackageManager.InstallCmd
	args := cmd[1:]
	
	if noConfirm {
		switch um.PackageManager.Name {
		case "pacman":
			args = append(args, "--noconfirm")
		case "apt":
			args = append(args, "-y")
		case "dnf", "yum":
			args = append(args, "-y")
		case "zypper":
			args = append(args, "-y")
		}
	}
	
	args = append(args, pkgNames...)
	
	var execCmd *exec.Cmd
	if um.SudoPath != "" {
		sudoArgs := append([]string{cmd[0]}, args...)
		execCmd = exec.Command(um.SudoPath, sudoArgs...)
	} else {
		execCmd = exec.Command(cmd[0], args...)
	}
	
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	
	return execCmd.Run()
}

func (um *UniversalManager) Remove(pkgNames []string, noConfirm bool) error {
	cmd := um.PackageManager.RemoveCmd
	args := cmd[1:]
	
	if noConfirm {
		switch um.PackageManager.Name {
		case "pacman":
			args = append(args, "--noconfirm")
		case "apt":
			args = append(args, "-y")
		case "dnf", "yum":
			args = append(args, "-y")
		}
	}
	
	args = append(args, pkgNames...)
	
	var execCmd *exec.Cmd
	if um.SudoPath != "" {
		sudoArgs := append([]string{cmd[0]}, args...)
		execCmd = exec.Command(um.SudoPath, sudoArgs...)
	} else {
		execCmd = exec.Command(cmd[0], args...)
	}
	
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	
	return execCmd.Run()
}

func (um *UniversalManager) Update(noConfirm bool) error {
	cmd := um.PackageManager.UpdateCmd
	args := cmd[1:]
	
	if noConfirm {
		switch um.PackageManager.Name {
		case "pacman":
			args = append(args, "--noconfirm")
		case "apt":
			args = append(args, "-y")
		case "dnf", "yum":
			args = append(args, "-y")
		}
	}
	
	var execCmd *exec.Cmd
	if um.SudoPath != "" {
		sudoArgs := append([]string{cmd[0]}, args...)
		execCmd = exec.Command(um.SudoPath, sudoArgs...)
	} else {
		execCmd = exec.Command(cmd[0], args...)
	}
	
	execCmd.Stdin = os.Stdin
	execCmd.Stdout = os.Stdout
	execCmd.Stderr = os.Stderr
	
	return execCmd.Run()
}

func (um *UniversalManager) IsInstalled(pkgName string) bool {
	cmd := um.PackageManager.ListCmd
	args := append(cmd[1:], pkgName)
	
	execCmd := exec.Command(cmd[0], args...)
	return execCmd.Run() == nil
}

func (um *UniversalManager) GetSystemInfo() string {
	return fmt.Sprintf("%s %s (%s)", 
		um.OSInfo.GetDisplayName(), 
		um.OSInfo.Version, 
		um.PackageManager.Name)
}