package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-webhook/api/model"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

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

		discordUrl := xcloud.Webhook.URL
		commitField := model.Field{
			Name:  "Commit Sha",
			Value: xcloud.CiBuildRun.Attributes.SourceCommit.CommitSha + " [See Changes](" + xcloud.CiBuildRun.Attributes.SourceCommit.HTMLURL + ")",
		}

		fields := []model.Field{commitField}

		executionProgress := xcloud.CiBuildRun.Attributes.ExecutionProgress
		fmt.Println(executionProgress)

		var embedTitle string = ""
		buildNumber := strconv.Itoa(xcloud.CiBuildRun.Attributes.Number)

		if executionProgress == "RUNNING" {
			embedTitle = xcloud.CiProduct.Attributes.Name + " (" + buildNumber + ")" + " is starting build :construction_site:"
		} else if executionProgress == "COMPLETE" {
			embedTitle = xcloud.CiProduct.Attributes.Name + " (" + buildNumber + ")" + " is ready on TestFlight :rocket:"

			for _, action := range xcloud.CiBuildActions {
				durationFormatted := getDuration(action.Attributes.StartedDate, action.Attributes.FinishedDate)
				field := model.Field{
					Name:  action.Attributes.Name,
					Value: durationFormatted,
				}

				fields = append(fields, field)
			}
		} else {
			fmt.Println("No needs to send notification")
			return
		}

		embed := model.Embeds{
			Title:       embedTitle,
			URL:         "",
			Description: "Author: **" + xcloud.CiBuildRun.Attributes.SourceCommit.Author.DisplayName + "**\n Lates Commit: **" + xcloud.CiBuildRun.Attributes.SourceCommit.Committer.DisplayName + "**",
			Color:       15258703,
			Fields:      fields,
			Thumbnail: struct {
				URL string "json:\"url\""
			}{
				URL: "https://images.bareksa.com/logo/1.0.0/default-image-news.jpg",
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
		req, _ := http.NewRequest("POST", discordUrl, payloadBuf)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, _ := client.Do(req)

		defer res.Body.Close()

		fmt.Println("response Status:", res.Status)
		// Print the body to the stdout
		io.Copy(os.Stdout, res.Body)
	}
}

// Time Util
func getDuration(start string, end string) string {
	t1, err := time.Parse(time.RFC3339, start)

	if err != nil {
		fmt.Println(err)
	}

	t2, err := time.Parse(time.RFC3339, end)

	if err != nil {
		fmt.Println(err)
	}

	totalSeconds := int(t2.Sub(t1).Seconds())

	hours := totalSeconds / 3600 % 60
	minutes := totalSeconds / 60 % 60
	seconds := totalSeconds % 60

	if hours > 0 {
		return fmt.Sprintf("Duration: **%d Hour %d Minutes %d Seconds**", hours, minutes, seconds)
	} else if minutes > 0 {
		return fmt.Sprintf("Duration: **%d Minutes %d Seconds**", minutes, seconds)
	} else {
		return fmt.Sprintf("Duration: **%d Seconds**", seconds)
	}
}
