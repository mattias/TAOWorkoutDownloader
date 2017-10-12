package tao

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

type Configuration struct {
	Oauth2     *oauth2.Config
	FileType   string
	TargetType string
}

func (c *Configuration) Load() error {
	err := c.loadOauth2()
	if err != nil {
		panic(err)
	}

	err = c.loadAppConfiguration()
	if err != nil {
		panic(err)
	}

	return err
}

func (c *Configuration) loadOauth2() error {
	oauth2Conf, err := os.Open("../config/oauth2.conf")
	if err != nil {
		panic(err)
	}
	defer oauth2Conf.Close()

	decoder := json.NewDecoder(oauth2Conf)
	oauth2config := oauth2.Config{}
	c.Oauth2 = &oauth2config

	return decoder.Decode(&oauth2config)
}

func (c *Configuration) loadAppConfiguration() error {
	appConf, err := os.Open("../config/app.conf")
	if err != nil {
		panic(err)
	}
	defer appConf.Close()

	decoder := json.NewDecoder(appConf)

	return decoder.Decode(&c)
}
