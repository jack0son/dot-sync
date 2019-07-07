package lib

import (
	"github.com/jgitgud/dot-sync/app"
	"github.com/jgitgud/dot-sync/repo"
)

const conf_path = "./data/dotsync.conf"

func Add(appName string, paths []string) error {
	// 1. initialise repository from config
	repo := repo.LoadRepo(conf_path)

	// 2. if app exists return error
	if app, ok := repo.findApp(appName); ok {
		return error.new("App already exists. Use track.")
	}

	// 3.
	return nil
}

func Track(appName string, paths []string) error {
	// 1. Initialise repository from config

	// 2. Check app exists in repository

	// 3. Add paths to app
	return nil
}

func Sync(appName string, paths []string) error {
	return nil
}

func Clone(appName string, paths []string) error {
	return nil
}

func CloneAll() error {
	return nil
}

func SyncAll() error {
	return nil
}
