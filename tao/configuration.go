package tao

import (
	"encoding/json"
	"os"

	"golang.org/x/oauth2"
)

type Configuration struct {
	Oauth2     oauth2.Config
	ServerURL  string
	DevicePath string
}

func (c *Configuration) Load() {
	c.Oauth2 = oauth2.Config{
		RedirectURL: "https://localhost:4443/oauthCallback/",
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://beta.trainasone.com/oauth/authorise",
			TokenURL: "https://beta.trainasone.com/oauth/token",
		},
		Scopes: []string{"TAO_MOBILE"},
	}

	err := c.loadOauth2Configuration()
	if err != nil {
		panic(err)
	}

	err = c.loadAppConfiguration()
	if err != nil {
		panic(err)
	}
}

func (c *Configuration) loadOauth2Configuration() error {
	oauth2Conf, err := os.Open("./config/oauth2.conf")
	if err != nil {
		panic(err)
	}
	defer oauth2Conf.Close()

	decoder := json.NewDecoder(oauth2Conf)

	return decoder.Decode(&c.Oauth2)
}

func (c *Configuration) loadAppConfiguration() error {
	appConf, err := os.Open("./config/app.conf")
	if err != nil {
		panic(err)
	}
	defer appConf.Close()

	decoder := json.NewDecoder(appConf)

	return decoder.Decode(&c)
}
