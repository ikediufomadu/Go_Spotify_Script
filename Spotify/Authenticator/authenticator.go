package Authenticator

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/zmb3/spotify/v2/auth"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify/v2"
)

// Loads environment variables from a .env file
var (
	loadVars     = godotenv.Load()
	clientID     = os.Getenv("CLIENT_ID")
	clientSecret = os.Getenv("CLIENT_SECRET")
)

// Spotify's authentication configuration
var (
	auth = spotifyauth.New(
		spotifyauth.WithRedirectURL(redirectURI),
		spotifyauth.WithScopes(spotifyauth.ScopePlaylistModifyPublic, spotifyauth.ScopePlaylistModifyPrivate),
		spotifyauth.WithClientID(clientID),
		spotifyauth.WithClientSecret(clientSecret),
	)
	ch             = make(chan *spotify.Client)
	searcherClient *spotify.Client
	userID         string
)

// OAuth constants
const (
	state       = "abc123"
	redirectURI = "http://localhost:8080/callback"
)

// StartLocalServer starts a local HTTP server for Spotify authentication
func StartLocalServer() {
	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Initiate the Spotify login process
	loginErr := loginUser()
	if loginErr != nil {
		os.Exit(1)
	}
}

// Initiates the Spotify login process
func loginUser() error {
	if clientID == "" || clientSecret == "" {
		return fmt.Errorf("CLIENT_ID or CLIENT_SECRET is not set in the environment")
	}

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:\n", url)

	// wait for auth to complete
	client := <-ch
	setClient(client)

	User, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	setUserID(User.ID)
	fmt.Println("You are logged in as:", User.ID)

	return nil
}

// Handles the callback from the Spotify authentication process
func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}

	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	ch <- client
}

func setClient(client *spotify.Client) {
	searcherClient = client
}

func GetClient() (*spotify.Client, error) {
	return searcherClient, nil
}

func setUserID(user string) {
	userID = user
}

func GetUserID() (string, error) {
	return userID, nil
}