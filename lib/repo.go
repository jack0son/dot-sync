package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jgitgud/dot-sync/app"
	"io/ioutil"
	"os"
	"path/filepath"
)

const CONFIG_FILE_NAME = "dotsync.conf"

// Repository
type Repo struct {
	apps   []app.App
	dir    string // Path to repo directory
	config RepoConfig
}

func (r *Repo) add(appName string, paths []string) error {
	a, err := app.NewApp(appName, paths)
	if err != nil {
		return err
	}
	r.apps = append(r.apps, *a)
	return nil
}

func (r *Repo) findApp(name string) (*app.App, bool) {
	for _, a := range r.apps {
		if a.Name == name {
			return &a, true
		}
	}

	return nil, false
}

func LoadRepo(configFile string) (*Repo, error) {
	cf, err := os.Open(configFile)
	if err != nil {
		return nil, errors.New("could not locate config file")
	}
	// Load config info into repo object
	fmt.Println(cf)

	decoder := json.NewDecoder(cf)
	repoConfig = new(RepoConfig)

	if err := decoder.Decode(repoConfig); err != nil {
		return nil, errors.New("could not decode config json")
	}

	return nil, nil
}

func (r *Repo) store() error {
	// Save repo config
	//appConfigs := make([]AppConfig, len(r.apps))
	appConfigs := make(map[string]AppConfig)

	for _, app := range r.apps {
		fileConfigs := make([]FileConfig, len(app.Files))
		for i, file := range app.Files {
			fileConfigs[i] = FileConfig{
				Name: file.Name,
				Dir:  file.Dir,
			}
		}
		appConfigs[app.Name] = fileConfigs
	}

	rc := RepoConfig{
		Dir:  r.dir,
		Apps: appConfigs,
	}

	rcJson, err := json.Marshal(rc)
	if err != nil {
		return fmt.Errorf("Failed to marshal config json with error %v", err)
	}

	file, err := os.Create(filepath.Join(r.dir, CONFIG_FILE_NAME))
	if err != nil {
		return fmt.Errorf("Failed to write with error %v", err)
	}
	defer file.Close()

	// @fix use json.NewEncoder(file), encoder.Encode(rc)
	file.WriteString(string(rcJson))
	return nil
}

type RepoConfig struct {
	//repo string `json:"repo"`
	Dir  string               `json:"dir"`
	Apps map[string]AppConfig `json:"apps"`
}

type AppConfig = []FileConfig

type FileConfig struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
	//hash uint32 `json:"i"`
}

func (rc *RepoConfig) Load() *Repo {
	repo := new(repo)
	repo.Dir = rc.Dir

	apps := make([]app.App, len(repo.Apps))
	for name, appConfig := range rc.Apps {

		files := make([]app.File, len(appConfig))
		for i, fileConfig := range appConfig {
			files[i] = app.File{
				fileConfig.Name,
				fileConfig.Dir,
			}
		}

		apps[i] = app.App{
			name,
			files,
		}
	}
}

/*
{"dir":"~/.dotsync","apps":{"bash":[{"name":".bashrc","dir":"~"},{"name":".bash_aliases","dir":"~"}],"vim":[{"name":".vimrc","dir":"~"}]}}


dotsync/
├── config.json
├── bash
│   ├── .bash_aliases
│   └── .bashrc
└── vim
    └── .vimrc

config.json
	"repo": {
		"dir":"~/repostiories/dotfiles/dotsync"
		"apps" : [
		{
			"name": "vim",
			"files": [
				{
					"name": ".vimrc",
					"dir": "~/",
					"hash": "0ab210",
				}]
		},
		{
			"name": "bash"
			"files": [
				{
					"name": ".bashrc",
					"dir": "~/",
					"hash": "0ab210",
				},
				{
					"name": ".bash_aliases",
					"dir": "~/",
					"hash": "0ab210",
				}
			]
		}]
	}

*/
