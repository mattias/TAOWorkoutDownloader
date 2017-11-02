package tao

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/mattias/TAOWorkoutDownloader/cache"
	"golang.org/x/oauth2"
)

type Client struct {
	Config  Configuration
	Token   *oauth2.Token
	Context context.Context
}

func (c *Client) Init() {
	c.Config.Load()
	c.Context = context.Background()
	c.loadToken()
}

func (c *Client) loadToken() {
	var token = new(oauth2.Token)

	err := cache.Load("token", token)
	if err == nil {
		c.Token = token
	}
}

func (c *Client) SaveNextWorkout() error {
	client := c.Config.Oauth2.Client(c.Context, c.Token)
	resp, err := client.Get(c.Config.ServerURL + "/api/mobile/plannedWorkout?access_token=" + c.Token.AccessToken)
	if err != nil {
		panic(err)
		// new token?
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("../" + "workout.fit")
	if err != nil {
		panic(err)
	}

	file.Write(body)

	return err
}
