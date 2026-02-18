package config

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Language          string `toml:"language"`
	MakeFlags         string `toml:"makeflags"`
	ParallelDownloads int    `toml:"parallel_downloads"`
	ColorScheme       string `toml:"color_scheme"`
	CacheDir          string `toml:"cache_dir"`
	LogDir            string `toml:"log_dir"`
}

func Load(path string) (*Config, error) {
	if path == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(home, ".config", "lt", "config.toml")
	}
	
	cfg := defaultConfig()
	
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			if err := cfg.Save(path); err != nil {
				return nil, err
			}
			return cfg, nil
		}
		return nil, err
	}
	
	if err := toml.Unmarshal(data, cfg); err != nil {
		return nil, err
	}
	
	return cfg, nil
}

func (c *Config) Save(path string) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	
	data, err := toml.Marshal(c)
	if err != nil {
		return err
	}
	
	return os.WriteFile(path, data, 0644)
}

func defaultConfig() *Config {
	home, _ := os.UserHomeDir()
	
	return &Config{
		Language:          "tr",
		MakeFlags:         "-j$(nproc)",
		ParallelDownloads: 5,
		ColorScheme:       "default",
		CacheDir:          filepath.Join(home, ".cache", "lt"),
		LogDir:            filepath.Join(home, ".local", "share", "lt", "logs"),
	}
}
