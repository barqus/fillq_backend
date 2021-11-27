package users

import "time"

type UserAccessInformation struct {
	AccessToken  string   `json:"access_token"`
	ExpiresIn    int      `json:"expires_in"`
	RefreshToken string   `json:"refresh_token"`
	Scope        []string `json:"scope"`
	TokenType    string   `json:"token_type"`
}

type TwitchChannelInformation struct {
	ID              string    `json:"id"`
	Login           string    `json:"login"`
	DisplayName     string    `json:"display_name"`
	Type            string    `json:"type"`
	BroadcasterType string    `json:"broadcaster_type"`
	Description     string    `json:"description"`
	ProfileImageURL string    `json:"profile_image_url"`
	OfflineImageURL string    `json:"offline_image_url"`
	ViewCount       int       `json:"view_count"`
	Email           string    `json:"email"`
	CreatedAt       time.Time `json:"created_at"`
}

type TwitchData struct {
	Data []TwitchChannelInformation `json:"data"`
}

type User struct {
	ID              string `json:"id"`
	DisplayName     string `json:"display_name"`
	ProfileImageURL string `json:"profile_image_url"`
	Email           string `json:"email"`
	TwitchCode      string `json:"twitch_code"`
	Role            string `json:"role"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	JWTToken 		string `json:"jwt_token"`
}
