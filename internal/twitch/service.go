package twitch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/barqus/fillq_backend/internal/participants"
	"net/url"
	//"github.com/barqus/fillq_backend/config"
	"github.com/barqus/fillq_backend/internal/common_http"
	//"github.com/dgrijalva/jwt-go"
	//"github.com/sirupsen/logrus"
	//"os"
	//"time"
	"github.com/barqus/fillq_backend/internal/database"
)

func UpdateUserTwitch(client common_http.Client, databaseClient *database.Database) {
	ctx := context.TODO()
	userTwitchURI := "https://api.twitch.tv/helix/search/channels"
	v := url.Values{}
	v.Set("first", "1")

	participantsStorage := participants.MustNewStorage(databaseClient)
	twitchStorage := MustNewStorage(databaseClient)
	allParticipants, err := participantsStorage.GetAllParticipants()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, item := range allParticipants {
		fmt.Println("UPDATING: ", item.Nickname)
		v.Add("query", item.TwitchChannel)
		rawResponse, _, err := client.Get(ctx, userTwitchURI, "02yzzgipgycv741y4nyuueuepfuirk","by7zl6rwazu7ks1z6sby63bnwq1267", v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var responseTwitchReturnedObject TwitchReturnedObject
		err = json.Unmarshal(rawResponse, &responseTwitchReturnedObject)
		if err != nil {
			fmt.Println(err)
			continue
		}

		twitchChannelInfo := &TwitchObject{
			GameName: responseTwitchReturnedObject.Data[0].GameName,
			IsLive: responseTwitchReturnedObject.Data[0].IsLive,
			StartedAt: responseTwitchReturnedObject.Data[0].StartedAt,
			DisplayName:     responseTwitchReturnedObject.Data[0].DisplayName,
			Title:     responseTwitchReturnedObject.Data[0].Title,
			ParticipantID: item.Id,
			TwitchID: responseTwitchReturnedObject.Data[0].Id,
		}

		fmt.Println(twitchChannelInfo)
		twitchExists, err := twitchStorage.twitchAccountAlreadyExists(twitchChannelInfo.TwitchID)
		fmt.Println(*twitchExists)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if *twitchExists == true {
			err = twitchStorage.UpdateTwitchAccountByID(*twitchChannelInfo, twitchChannelInfo.TwitchID)
			if err != nil {
				fmt.Println(err)
				continue
			}
		} else {
			err = twitchStorage.AddTwitchAccount(*twitchChannelInfo)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		v.Del("query")
	}

}