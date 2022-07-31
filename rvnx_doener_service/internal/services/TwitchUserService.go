package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/nicklaw5/helix/v2"
	"rvnx_doener_service/ent"
	"rvnx_doener_service/ent/twitchuser"
	"strconv"
	"time"
)

func NewTwitchUserService(
	client *ent.Client,
	eventService *EventService,
	clientID, clientSecret string,
) *TwitchUserService {
	return &TwitchUserService{
		twitchClientID:     clientID,
		twitchClientSecret: clientSecret,
		client:             client.TwitchUser,
		context:            context.Background(),
		eventService:       eventService,
	}
}

type TwitchUserService struct {
	twitchClientID, twitchClientSecret string
	client                             *ent.TwitchUserClient
	eventService                       *EventService
	context                            context.Context
}

func (t *TwitchUserService) GetClientID() string {
	return t.twitchClientID
}

func (t *TwitchUserService) GetClientSecret() string {
	return t.twitchClientSecret
}

func (t *TwitchUserService) FinalizeUserLogin(code, redirectURI string) (*ent.TwitchUser, error) {
	helixClient, err := helix.NewClient(&helix.Options{
		ClientID:     t.GetClientID(),
		ClientSecret: t.GetClientSecret(),
		RedirectURI:  redirectURI,
	})
	if err != nil {
		return nil, err
	}

	token, err := helixClient.RequestUserAccessToken(code)
	if err != nil {
		return nil, err
	}

	if token.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%d - %s (%s)", token.StatusCode, token.Error, token.ErrorMessage))
	}

	twitchUser := ent.TwitchUser{
		OauthToken:        token.Data.AccessToken,
		OauthRefreshToken: token.Data.RefreshToken,
	}

	helixClient.SetUserAccessToken(token.Data.AccessToken)
	resp, err := helixClient.GetUsers(&helix.UsersParams{})
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 401 {
		// TODO: implement token refresh (here kinda unlikely though)
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("%d - %s (%s)", resp.StatusCode, resp.Error, resp.ErrorMessage))
	}

	if len(resp.Data.Users) == 0 {
		return nil, errors.New("empty response from twitch")
	}

	u := resp.Data.Users[0]
	twitchUser.ID, err = strconv.ParseInt(u.ID, 10, 64)
	if err != nil {
		return nil, errors.New("invalid twitch user id")
	}

	twitchUser.Login = u.Login
	twitchUser.DisplayName = u.DisplayName
	twitchUser.Email = u.Email
	twitchUser.CreatedAt = time.UnixMicro(u.CreatedAt.UnixMicro())

	return t.CreateOrUpdateUser(&twitchUser)
}

func (t *TwitchUserService) GetTwitchUserData(id int64) (user *helix.User, exists bool, err error) {
	u, err := t.client.Get(t.context, id)
	if ent.IsNotFound(err) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, nil
	}

	helixClient, err := helix.NewClient(&helix.Options{
		ClientID:        t.GetClientID(),
		ClientSecret:    t.GetClientSecret(),
		UserAccessToken: u.OauthToken,
	})
	if err != nil {
		return nil, false, nil
	}

	resp, err := helixClient.GetUsers(&helix.UsersParams{})
	if err != nil {
		return nil, false, nil
	}

	if resp.StatusCode == 401 {
		// TODO: implement token refresh (here kinda unlikely though)
	}

	if resp.StatusCode != 200 {
		return nil, false, errors.New(fmt.Sprintf("%d - %s (%s)", resp.StatusCode, resp.Error, resp.ErrorMessage))
	}

	if len(resp.Data.Users) == 0 {
		return nil, false, errors.New("empty response from twitch")
	}

	return &resp.Data.Users[0], true, nil
}

func (t *TwitchUserService) CreateOrUpdateUser(user *ent.TwitchUser) (*ent.TwitchUser, error) {
	_, err := t.client.Query().Unique(false).Where(twitchuser.ID(user.ID)).First(t.context)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if ent.IsNotFound(err) {
		newUser, err := t.client.Create().
			SetID(user.ID).
			SetCreatedAt(user.CreatedAt).
			SetLogin(user.Login).
			SetDisplayName(user.DisplayName).
			SetEmail(user.Email).
			SetOauthToken(user.OauthToken).
			SetOauthRefreshToken(user.OauthRefreshToken).
			Save(t.context)
		if err != nil {
			return nil, err
		}

		t.eventService.LogFirstTimeUserLogin(newUser)

		return newUser, nil
	}

	_, err = t.client.Update().Where(twitchuser.ID(user.ID)).
		SetCreatedAt(user.CreatedAt).
		SetLogin(user.Login).
		SetDisplayName(user.DisplayName).
		SetEmail(user.Email).
		SetOauthToken(user.OauthToken).
		SetOauthRefreshToken(user.OauthRefreshToken).
		Save(t.context)
	if err != nil {
		return nil, err
	}

	t.eventService.LogUserLogin(user)

	return user, nil
}
