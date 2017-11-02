package tao

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/oauth2"
)

const serverURL string = "https://stage.trainasone.com"

type Client struct {
	Config  Configuration
	Token   *oauth2.Token
	Context context.Context
}

func (c *Client) Init() {
	c.Config.Load()
	c.Context = context.Background()
}

func (c *Client) SaveNextWorkoutTo(path string) (string, error) {
	client := c.Config.Oauth2.Client(c.Context, c.Token)
	resp, err := client.Get(serverURL + "/api/mobile/plannedWorkout?access_token=" + c.Token.AccessToken)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Body: %s", body)

	file, err := os.Create(path + "workout" + c.getWorkoutFileType())
	if err != nil {
		panic(err)
	}

	file.Write([]byte("test"))

	return strings.Trim(file.Name(), path), err
}

func (c *Client) getWorkoutFileType() string {
	return "." + c.Config.Workout.FileType
}
