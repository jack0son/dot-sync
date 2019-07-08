package app

import (
	//"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// TODO:
//	- should config jsons be an interface of struct they represent
//	- memory usage in json marshalling

// A user named application which tracks a number of files
type App struct {
	Name  string
	Files []File
}

//func LoadApp(name string, fcs []lib.FileConfig) *App {

// Change to LoadApp
func NewApp(name string, paths []string) (*App, error) {
	app := new(App)
	for _, path := range paths {
		file, err := NewFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading path %v: %v", path, err)
			return nil, err
		}
		app.Files = append(app.Files, *file)
	}

	app.Name = name

	return app, nil
}

func (app *App) track(files ...File) error {
	return nil
}

func (app *App) sync(files ...File) error {

	return nil
}

type File struct {
	Name     string
	Dir      string
	contents []byte
	hash     uint32
}

// Change to Load File to match Repo LoadRepo
//func (f *File) Load(path string) (*File, error) {
func NewFile(path string) (*File, error) {
	// 1. attempt to read file
	file := new(File)
	if err := file.read(path); err != nil {
		return nil, err
	}
	return file, nil
}

func (f *File) read(path string) error {
	var err error
	if f.contents, err = ioutil.ReadFile(path); err != nil {
		return err
	}

	// Extact name from path
	f.Dir = filepath.Dir(path)
	f.Name = filepath.Base(path)

	return nil
}

/*func NewFile(path string) *File {
	return File {
		Name: filepath.Base(path),
		Dir: filepath.Dir(path),
		contents: nil,
		hash: nil,
	}
}*/
