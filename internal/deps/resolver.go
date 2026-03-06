package deps

import (
	"fmt"
	"strings"

	"github.com/cicegimsin/lt/internal/aur"
	"github.com/cicegimsin/lt/internal/pacman"
)

type Resolver struct {
	aurClient    *aur.Client
	pacmanClient *pacman.Client
}

type InstallPlan struct {
	RepoPackages []string
	AURPackages  []string
	Order        []string
}

func NewResolver(aurClient *aur.Client, pacmanClient *pacman.Client) *Resolver {
	return &Resolver{
		aurClient:    aurClient,
		pacmanClient: pacmanClient,
	}
}

func (r *Resolver) ResolveDependencies(pkgName string) (*InstallPlan, error) {
	plan := &InstallPlan{
		RepoPackages: []string{},
		AURPackages:  []string{},
		Order:        []string{},
	}
	
	visited := make(map[string]bool)
	
	if err := r.resolveDeps(pkgName, plan, visited); err != nil {
		return nil, err
	}
	
	return plan, nil
}

func (r *Resolver) resolveDeps(pkgName string, plan *InstallPlan, visited map[string]bool) error {
	if visited[pkgName] {
		return nil
	}
	visited[pkgName] = true
	
	if r.pacmanClient.IsInstalled(pkgName) {
		return nil
	}
	
	repoPackages, err := r.pacmanClient.Search(pkgName)
	if err == nil && len(repoPackages) > 0 {
		for _, pkg := range repoPackages {
			if pkg.Name == pkgName {
				plan.RepoPackages = append(plan.RepoPackages, pkgName)
				plan.Order = append(plan.Order, pkgName)
				return nil
			}
		}
	}
	
	aurPkg, err := r.aurClient.Info(pkgName)
	if err != nil {
		return fmt.Errorf("paket bulunamadı: %s", pkgName)
	}
	
	allDeps := append(aurPkg.Depends, aurPkg.MakeDepends...)
	for _, dep := range allDeps {
		depName := r.extractPackageName(dep)
		if depName != "" {
			if err := r.resolveDeps(depName, plan, visited); err != nil {
				return err
			}
		}
	}
	
	plan.AURPackages = append(plan.AURPackages, pkgName)
	plan.Order = append(plan.Order, pkgName)
	
	return nil
}

func (r *Resolver) extractPackageName(dep string) string {
	dep = strings.TrimSpace(dep)
	if dep == "" {
		return ""
	}
	
	operators := []string{">=", "<=", "=", ">", "<"}
	for _, op := range operators {
		if idx := strings.Index(dep, op); idx != -1 {
			dep = dep[:idx]
			break
		}
	}
	
	return strings.TrimSpace(dep)
}

func (r *Resolver) IsAURPackage(pkgName string) bool {
	_, err := r.aurClient.Info(pkgName)
	return err == nil
}

func (r *Resolver) IsRepoPackage(pkgName string) bool {
	packages, err := r.pacmanClient.Search(pkgName)
	if err != nil {
		return false
	}
	
	for _, pkg := range packages {
		if pkg.Name == pkgName {
			return true
		}
	}
	
	return false
}