package main

import (
	"fmt"
	"net/http"

	"github.com/mattias/TAOWorkoutDownloader/tao"
	"golang.org/x/oauth2"
)

func main() {
	var client tao.Client
	client.Init()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")

		fmt.Printf("We got code: %s\n", code)

		_, err := client.SaveNextWorkoutTo("./", code)
		if err != nil {
			panic(err)
		}

	})
	// TODO: Save token after we got it for future uses and skip this if we have token.
	// Only run this if token becomes invalid
	url := client.Config.Oauth2.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v\n", url)

	fmt.Println(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))
}
