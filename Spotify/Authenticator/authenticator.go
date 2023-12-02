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
		spotifyauth.WithScopes(spotifyauth.ScopePlaylistModifyPublic, spotifyauth.ScopePlaylistModifyPrivate, spotifyauth.ScopePlaylistReadPrivate, spotifyauth.ScopePlaylistReadCollaborative),
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
	loginUser()
}

// Initiates the Spotify login process
func loginUser() {
	if clientID == "" || clientSecret == "" {
		log.Fatal("CLIENT_ID or CLIENT_SECRET is not set in the environment")
	}

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting this link in your browser:\n", url)

	// wait for auth to complete
	client := <-ch
	setClient(client)

	User, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal("Could not retrieve user data.\n", err)
	}
	setUserID(User.ID)
	fmt.Printf("You are logged in as: %s.\n", User.User.DisplayName)
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

func setUserID(user string) {
	userID = user
}

func GetClient() *spotify.Client {
	return searcherClient
}

func GetUserID() string {
	return userID
}
