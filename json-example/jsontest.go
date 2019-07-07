package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	rc := RepoConfig{
		Dir: "~/.dotsync",
		Apps: map[string]AppConfig{
			"vim": AppConfig{
				FileConfig{
					Name: ".vimrc",
					Dir:  "~",
				},
			},

			"bash": {
				FileConfig{
					Name: ".bashrc",
					Dir:  "~",
				},
				FileConfig{
					Name: ".bash_aliases",
					Dir:  "~",
				},
			},
		},
	}
	fmt.Println(rc.Dir)

	rcJson, err := json.Marshal(rc)
	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println(string(rcJson))

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
