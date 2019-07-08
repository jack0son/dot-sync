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
	var repo *Repo
	repoPath := loadConfig()

	// @ fix rename to repoDir
	if repoPath == "" {
		// No existing config pointing to repo
		// Create a new repository config file
		repoPath, err := Init()
		if err != nil {
			return err
		}
		defer createConfig(repoPath)

		repo = NewRepo(repoPath)
		defer repo.Store()

	} else {
		// Use existing config file
		var err error
		repo, err = LoadRepo(repoPath)
		if err != nil {
			return err
		}

		// 2. if app exists return error
		if app, ok := repo.findApp(appName); ok {
			return fmt.Errorf("App \"%s\" already exists. Try track command.", app.Name)
		}
		defer repo.Store()
	}

	fmt.Println("Calling add on repo:  ", repo)
	if err := repo.Add(appName, paths); err != nil {
		return err
	}

	return nil
}

/**
 * Initialise dotsync configuration by setting the repo directory
 *
 * @dev Need to be able to defer writing the config file until we know that
 * that the associated repo config has been written to the the user's
 * dot-files repository.
 *
 * @returns (repoPath)
 */
func Init() (string, error) {
	// Refactor this function - doesn't fit in this file and doesn't stand alone
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Configuring dot-sync. Enter the path to your dot files repository.")
	fmt.Print("  repo path: ")
	scanner.Scan()
	path, _ := filepath.Abs(scanner.Text())

	if !IsValidDir(path) {
		// Must have write permissions for this test
		return "", errors.New("Provided repository path invalid")
	}

	return path, nil
}

// Load dotsync configuration file and return repo directory
func loadConfig() string {
	byteString, err := ioutil.ReadFile(conf_path)
	if err != nil {
		return ""
	}

	return string(byteString)
}

func createConfig(repoPath string) error {
	// ~/.dotsync/config
	cf, err := os.Create(conf_path)
	if err != nil {
		return errors.New("Failed to create dotsync config file")
	}
	defer cf.Close()
	cf.WriteString(filepath.Clean(repoPath))

	fmt.Println("Wrote config file to ", conf_path)
	return nil
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
