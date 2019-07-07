package lib

import (
	"bufio"
	"fmt"
	//"github.com/jgitgud/dot-sync/app"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

const conf_path = "./data/dotsync.conf"

func Add(appName string, paths []string) error {
	// 1. initialise repository from config

	repoPath := LoadConfig()
	if repoPath == "" {
		return Init()
	}

	repo, err := LoadRepo(repoPath)
	if err != nil {
		return err
	}

	// 2. if app exists return error
	if app, ok := repo.findApp(appName); ok {
		fmt.Println(app.Name)
	}

	return errors.New("App already exists. Use track.")
}

// Initialise dotsync configuration by setting the repo directory
func Init() error {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Configuring dot-sync. Please enter the path to your dot files repository.")
	fmt.Print("\trepo path: ")
	scanner.Scan()
	path := scanner.Text()

	if !IsValidDir(path) {
		return errors.New("Provided repository path invalid")
	}

	// ~/.dotsync/config
	cf, err := os.Create(conf_path)
	if err != nil {
		return errors.New("Failed to create dotsync config file")
	}
	defer cf.Close()
	cf.WriteString(path)

	return nil
}

// Load dotsync configuration file and return repo directory
func LoadConfig() string {
	byteString, err := ioutil.ReadFile(conf_path)
	if err != nil {
		return ""
	}

	return string(byteString)
}

func Track(appName string, paths []string) error {
	// 1. Initialise repository from config

	// 2. Check app exists in repository

	// 3. Add paths to app
	return nil
}

func Sync(appName string) error {
	// @todo able to sync individual file
	return nil
}

func Clone(appName string) error {
	return nil
}

func CloneAll() error {
	return nil
}

func SyncAll() error {
	return nil
}

func IsValidDir(dp string) bool {

	fp := filepath.Join(dp, "tmp")
	// Check if file already exists
	if _, err := os.Stat(fp); err == nil {
		return true
	}

	// Attempt to create it
	var d []byte
	if err := ioutil.WriteFile(fp, d, 0644); err == nil {
		os.Remove(fp) // And delete it
		return true
	}
	return false
}
