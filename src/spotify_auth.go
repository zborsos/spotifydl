package spotifydl

import (
	"context"
	"fmt"
	"log"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2/clientcredentials"
)

// UserData is a struct to hold all variables
type UserData struct {
	UserClient      *spotify.Client
	TrackList       []spotify.FullTrack
	SimpleTrackList []spotify.SimpleTrack
	YoutubeIDList   []string
}

// InitAuth starts Authentication
func InitAuth() *spotify.Client {
	ctx := context.Background()
	// Please do not misuse :)
	config := &clientcredentials.Config{
		ClientID:     "608f81b355634c058a5c76ba2badf94a",
		ClientSecret: "0da7ecf3bd7f4934b34cdce9cf5a253b",
		TokenURL:     spotifyauth.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}
	fmt.Println("Token ", token)

	httpClient := spotifyauth.New().Client(ctx, token)
	client := spotify.New(httpClient, spotify.WithBaseURL("https://api.spotify.com/v1/"))

	return client
}
