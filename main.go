package main

import (
	"github.com/barqus/fillq_backend/cmd"
)

func main() {
	//httpClient := common_http.NewClient(http.DefaultClient)
	//youtube.MustNewService(httpClient).GetChannelsVideos()
	cmd.Execute()
}
