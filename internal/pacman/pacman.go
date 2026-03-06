package pacman

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Client struct {
	pacmanPath string
	sudoPath   string
}

type Package struct {
	Name        string
	Version     string
	Description string
	Repository  string
}

func NewClient(pacmanPath, sudoPath string) *Client {
	return &Client{
		pacmanPath: pacmanPath,
		sudoPath:   sudoPath,
	}
}

func (c *Client) IsInstalled(pkgName string) bool {
	cmd := exec.Command(c.pacmanPath, "-Qi", pkgName)
	return cmd.Run() == nil
}

func (c *Client) Search(query string) ([]Package, error) {
	cmd := exec.Command(c.pacmanPath, "-Ss", query)
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var packages []Package
	lines := strings.Split(string(output), "\n")
	
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
	
	return packages, nil
}

func (c *Client) Install(pkgNames []string, noConfirm bool) error {
	args := []string{c.pacmanPath, "-S"}
	if noConfirm {
		args = append(args, "--noconfirm")
	}
	args = append(args, pkgNames...)
	
	cmd := exec.Command(c.sudoPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func (c *Client) Remove(pkgNames []string, noConfirm bool) error {
	args := []string{c.pacmanPath, "-Rs"}
	if noConfirm {
		args = append(args, "--noconfirm")
	}
	args = append(args, pkgNames...)
	
	cmd := exec.Command(c.sudoPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func (c *Client) Update(noConfirm bool) error {
	args := []string{c.pacmanPath, "-Syu"}
	if noConfirm {
		args = append(args, "--noconfirm")
	}
	
	cmd := exec.Command(c.sudoPath, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func (c *Client) GetInstalledPackages() ([]Package, error) {
	cmd := exec.Command(c.pacmanPath, "-Q")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var packages []Package
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			packages = append(packages, Package{
				Name:    parts[0],
				Version: parts[1],
			})
		}
	}
	
	return packages, scanner.Err()
}

func (c *Client) GetForeignPackages() ([]Package, error) {
	cmd := exec.Command(c.pacmanPath, "-Qm")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	var packages []Package
	scanner := bufio.NewScanner(strings.NewReader(string(output)))
	
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			packages = append(packages, Package{
				Name:    parts[0],
				Version: parts[1],
			})
		}
	}
	
	return packages, scanner.Err()
}