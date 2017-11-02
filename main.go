package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/elgs/gostrgen"
	"github.com/mattias/TAOWorkoutDownloader/cache"
	"github.com/mattias/TAOWorkoutDownloader/tao"
	"golang.org/x/oauth2"
)

func main() {
	var client tao.Client
	charsToGenerate := 20
	charSet := gostrgen.LowerUpperDigit
	oauthStateString, err := gostrgen.RandGen(charsToGenerate, charSet, "", "")
	if err != nil {
		panic(err)
	}

	client.Init()

	http.HandleFunc("/oauthCallback/", func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		if state != oauthStateString {
			log.Printf("Invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
			return
		}

		code := r.FormValue("code")
		if code == "" {
			log.Printf("We got no code! Exiting...")
			os.Exit(1)
		}

		var token *oauth2.Token

		token, err = client.Config.Oauth2.Exchange(client.Context, code)
		if err != nil {
			log.Fatalf("Code exchange failed with '%s'\n", err)
		}

		client.Token = token
		err = cache.Save("token", token)
		if err != nil {
			panic(err)
		}

		err = client.SaveNextWorkout()
		if err != nil {
			panic(err)
		}

	})

	// TODO: Save token after we got it for future uses and skip this if we have token.
	// Only run this if token becomes invalid
	if client.Token == nil {
		url := client.Config.Oauth2.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
		open(url)
	} else {
		err = client.SaveNextWorkout()
		if err != nil {
			panic(err)
		}
	}

	fmt.Println(http.ListenAndServeTLS(":4443", "server.crt", "server.key", nil))
}

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
