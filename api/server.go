package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"io"
	"os"
	"github.com/gorilla/mux"
	"bytes"
	"go-webhook/api/model"
)

var DiscordUrl string = "https://discord.com/api/webhooks/1046725677140934676/haqiK8EMDhTFQcflNzBjxpRXFaOKdd93IWVN5zOYFfBhzdUWzCTJr389tql7yQ_BRrGb"

type Server struct {
	*mux.Router
}

func NewServer() *Server {
	s := &Server{
		Router: mux.NewRouter(),
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	s.HandleFunc("/xcloud-webhook", s.xcloudToDiscord()).Methods("POST")
}

func (s *Server) xcloudToDiscord() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		reqBody, _ := ioutil.ReadAll(r.Body)
		var xcloud model.Xcloud
		err := json.Unmarshal(reqBody, &xcloud)
		if err != nil {
			fmt.Println(err)
			return
		}

		commitField := model.Field{
			Name:  "Commit Sha",
			Value: xcloud.CiBuildRun.Attributes.SourceCommit.CommitSha + "[See Changes](" + xcloud.CiBuildRun.Attributes.SourceCommit.HTMLURL + ")",
		}

		fields := []model.Field{commitField}

		embed := model.Embeds{
			Title:       xcloud.CiProduct.Attributes.Name,
			URL:         "",
			Description: "Lates Commit by **" + xcloud.CiBuildRun.Attributes.SourceCommit.Author.DisplayName + "**",
			Color:       15258703,
			Fields:      fields,
			Thumbnail: struct {
				URL string "json:\"url\""
			}{
				URL: "https://e7.pngegg.com/pngimages/375/318/png-clipart-webhook-slack-form-web-application-world-wide-web-text-logo-thumbnail.png",
			},
			Footer: struct {
				Text    string "json:\"text\""
				IconURL string "json:\"icon_url\""
			}{
				Text:    "Be Awesome Today!",
				IconURL: "https://cdn.icon-icons.com/icons2/788/PNG/512/cool_icon-icons.com_64891.png",
			},
		}

		embeds := []model.Embeds{embed}

		discord := model.Discord{
			Username:  "Xcode Cloud",
			AvatarURL: "https://developer.apple.com/assets/elements/icons/xcode-cloud/xcode-cloud-128x128_2x.png",
			Content:   "",
			Embeds:    embeds}

		payloadBuf := new(bytes.Buffer)
		json.NewEncoder(payloadBuf).Encode(discord)
		req, _ := http.NewRequest("POST", DiscordUrl, payloadBuf)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, _ := client.Do(req)

		defer res.Body.Close()

		fmt.Println("response Status:", res.Status)
		// Print the body to the stdout
		io.Copy(os.Stdout, res.Body)
	}
}
