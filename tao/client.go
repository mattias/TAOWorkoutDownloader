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
	file, err := os.Create(path + "workout.fit")
	if err != nil {
		panic(err)
	}

	file.Write([]byte("test"))

	return strings.Trim(file.Name(), path), err
}
