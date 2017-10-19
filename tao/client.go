package tao

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const serverURL string = "https://beta.trainasone.com"

type Client struct {
	Config Configuration
}

func (c *Client) Init() {
	c.Config.Load()
}

func (c *Client) SaveNextWorkoutTo(path string, code string) (string, error) {
	ctx := context.Background()

	token, err := c.Config.Oauth2.Exchange(ctx, code)
	if err != nil {
		log.Fatalf("Code exchange failed with '%s'\n", err)
	}

	client := c.Config.Oauth2.Client(ctx, token)
	resp, err := client.Get(serverURL + "/api/mobile/plannedWorkout?access_token=" + token.AccessToken)
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
