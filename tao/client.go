package tao

import (
	"os"
	"strings"
)

type Client struct {
	Config Configuration
}

func (c *Client) Init() {
	c.Config.Load()
}

func (c *Client) SaveNextWorkoutTo(path string) (string, error) {
	// TODO: Rewrite to use the api later and download the file to specified path directly
	file, err := os.Create(path + "workout" + c.getWorkoutFileType())
	if err != nil {
		panic(err)
	}

	file.Write([]byte("test"))

	return strings.Trim(file.Name(), path), err
}

// TODO: Probably won't need this later when real code is in place (service)
func (c *Client) getWorkoutFileType() string {
	return "." + c.Config.Workout.FileType
}
