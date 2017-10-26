package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/elgs/gostrgen"
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

	fmt.Printf("State string generated: %s\n", oauthStateString)

	client.Init()

	http.HandleFunc("/oauthCallback/", func(w http.ResponseWriter, r *http.Request) {
		state := r.FormValue("state")
		fmt.Printf("State string received: %s\n", state)
		if state != oauthStateString {
			log.Printf("Invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
			return
		}

		code := r.FormValue("code")
		if code == "" {
			log.Printf("We got no code! Exiting...")
			os.Exit(1)
		}
		fmt.Printf("We got code: %s\n", code)

		fmt.Printf("Configuration for Oauth2:\n%+v\n", client.Config.Oauth2)

		var token *oauth2.Token

		token, err = client.Config.Oauth2.Exchange(client.Context, code)
		if err != nil {
			log.Fatalf("Code exchange failed with '%s'\n", err)
		}

		client.Token = token

		fmt.Printf("%+v\n", token)

		_, err = client.SaveNextWorkoutTo("./")
		if err != nil {
			panic(err)
		}

	})

	// TODO: Save token after we got it for future uses and skip this if we have token.
	// Only run this if token becomes invalid
	if client.Token == nil {
		url := client.Config.Oauth2.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
		fmt.Printf("Url generated: %s\n", url)
		open(url)
	} else {
		_, err = client.SaveNextWorkoutTo("./")
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
