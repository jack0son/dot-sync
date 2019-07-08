package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jgitgud/dot-sync/app"
	"os"
	"path/filepath"
)

const CONFIG_FILE_NAME = "dotsync.json"

// Repository
type Repo struct {
	apps []app.App
	Dir  string // Path to repo directory
	//config RepoConfig
}

func (r *Repo) Add(appName string, paths []string) error {
	a, err := app.NewApp(appName, paths)
	if err != nil {
		return err
	}
	r.apps = append(r.apps, *a)
	return nil
}

func (r *Repo) getApps() []app.App {
	return r.apps
}

func (r *Repo) findApp(name string) (*app.App, bool) {
	for _, a := range r.apps {
		if a.Name == name {
			return &a, true
		}
	}

	return nil, false
}

func NewRepo(repoPath string) *Repo {
	repo := new(Repo)
	repo.Dir = repoPath

	// @fix poor initialisation
	//repo.apps = make([]app.App, 0)

	return repo
}

func LoadRepo(configPath string) (*Repo, error) {
	cf, err := os.Open(filepath.Join(configPath, CONFIG_FILE_NAME))
	if err != nil {
		return nil, errors.New("could not locate config file")
	}
	// Load config info into repoConfig struct
	//fmt.Println(cf)

	decoder := json.NewDecoder(cf)
	repoConfig := new(RepoConfig)

	if err := decoder.Decode(repoConfig); err != nil {
		return nil, fmt.Errorf("Failed to decode config json: %v", err)
	}

	repo := repoConfig.load()

	return repo, nil
}

func (r *Repo) Store() error {
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
		Dir:  r.Dir,
		Apps: appConfigs,
	}

	rcJson, err := json.MarshalIndent(rc, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to marshal config json with error %v", err)
	}

	file, err := os.Create(filepath.Join(r.Dir, CONFIG_FILE_NAME))
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
	Dir  string               `json:"dir"` // this will change to profile
	Apps map[string]AppConfig `json:"apps"`
}

type AppConfig []FileConfig

type FileConfig struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
	//hash uint32 `json:"i"`
}

func (fc *FileConfig) getName() string {
	return fc.Name
}

func (fc *FileConfig) getDir() string {
	return fc.Dir
}

// This style contradicts app load
func (rc *RepoConfig) load() *Repo {
	repo := new(Repo)
	repo.Dir = rc.Dir

	//apps := make([]app.App, len(rc.Apps))
	var apps []app.App
	for name, appConfig := range rc.Apps {
		apps = append(apps, *appConfig.load(name))
	}

	repo.apps = apps

	return repo
}

func (ac AppConfig) load(name string) *app.App {
	a := new(app.App)
	a.Name = name

	a.Files = make([]app.File, len(ac))
	for i, fc := range ac {
		file := new(app.File)
		file.Name = fc.Name
		file.Dir = fc.Dir

		a.Files[i] = *file //append(a.Files, *file)
	}

	return a
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
		"profile": light
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
