package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const endpointProfile string = "https://www.googleapis.com/oauth2/v2/userinfo"

type googleUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	FirstName string `json:"given_name"`
	LastName  string `json:"family_name"`
	Link      string `json:"link"`
	Picture   string `json:"picture"`
}

func fetchGoogleUser(accessToken string) (googleUser, error) {
	response, err := new(http.Client).Get(endpointProfile + "?access_token=" + url.QueryEscape(accessToken))
	if err != nil {
		return googleUser{}, fmt.Errorf("failed to fetch user information: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return googleUser{}, fmt.Errorf("google responded with a %d trying to fetch user information", response.StatusCode)
	}

	var u googleUser

	if err = json.NewDecoder(response.Body).Decode(&u); err != nil {
		return googleUser{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return u, nil
}
