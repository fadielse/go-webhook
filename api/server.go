package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"io"
	"os"
	"github.com/gorilla/mux"
	"bytes"
)

var DiscordUrl string = "https://discord.com/api/webhooks/1046725677140934676/haqiK8EMDhTFQcflNzBjxpRXFaOKdd93IWVN5zOYFfBhzdUWzCTJr389tql7yQ_BRrGb"

type Xcloud struct {
	Webhook struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"webhook"`
	App struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"app"`
	CiWorkflow struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name               string    `json:"name"`
			Description        string    `json:"description"`
			LastModifiedDate   time.Time `json:"lastModifiedDate"`
			IsEnabled          bool      `json:"isEnabled"`
			IsLockedForEditing bool      `json:"isLockedForEditing"`
		} `json:"attributes"`
	} `json:"ciWorkflow"`
	CiProduct struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name        string    `json:"name"`
			CreatedDate time.Time `json:"createdDate"`
			ProductType string    `json:"productType"`
		} `json:"attributes"`
	} `json:"ciProduct"`
	CiBuildRun struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Number       int       `json:"number"`
			CreatedDate  time.Time `json:"createdDate"`
			StartedDate  time.Time `json:"startedDate"`
			FinishedDate time.Time `json:"finishedDate"`
			SourceCommit struct {
				CommitSha string `json:"commitSha"`
				Author    struct {
					DisplayName string `json:"displayName"`
				} `json:"author"`
				Committer struct {
					DisplayName string `json:"displayName"`
				} `json:"committer"`
				HTMLURL string `json:"htmlUrl"`
			} `json:"sourceCommit"`
			IsPullRequestBuild bool   `json:"isPullRequestBuild"`
			ExecutionProgress  string `json:"executionProgress"`
			CompletionStatus   string `json:"completionStatus"`
		} `json:"attributes"`
	} `json:"ciBuildRun"`
	CiBuildActions []struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name         string    `json:"name"`
			ActionType   string    `json:"actionType"`
			StartedDate  time.Time `json:"startedDate"`
			FinishedDate time.Time `json:"finishedDate"`
			IssueCounts  struct {
				AnalyzerWarnings int `json:"analyzerWarnings"`
				Errors           int `json:"errors"`
				TestFailures     int `json:"testFailures"`
				Warnings         int `json:"warnings"`
			} `json:"issueCounts"`
			ExecutionProgress string `json:"executionProgress"`
			CompletionStatus  string `json:"completionStatus"`
			IsRequiredToPass  bool   `json:"isRequiredToPass"`
		} `json:"attributes,omitempty"`
		Relationships struct {
			Build struct {
				ID         string `json:"id"`
				Type       string `json:"type"`
				Attributes struct {
					Platform string `json:"platform"`
				} `json:"attributes"`
			} `json:"build"`
		} `json:"relationships,omitempty"`
		Attributes0 struct {
			Name         string    `json:"name"`
			StartedDate  time.Time `json:"startedDate"`
			FinishedDate time.Time `json:"finishedDate"`
			IssueCounts  struct {
				AnalyzerWarnings int `json:"analyzerWarnings"`
				Errors           int `json:"errors"`
				TestFailures     int `json:"testFailures"`
				Warnings         int `json:"warnings"`
			} `json:"issueCounts"`
			ExecutionProgress string `json:"executionProgress"`
			CompletionStatus  string `json:"completionStatus"`
			IsRequiredToPass  bool   `json:"isRequiredToPass"`
		} `json:"attributes,omitempty"`
		Relationships0 struct {
		} `json:"relationships,omitempty"`
	} `json:"ciBuildActions"`
	ScmProvider struct {
		Type       string `json:"type"`
		Attributes struct {
			ScmProviderType struct {
				ScmProviderType string `json:"scmProviderType"`
				DisplayName     string `json:"displayName"`
				IsOnPremise     bool   `json:"isOnPremise"`
			} `json:"scmProviderType"`
			Endpoint string `json:"endpoint"`
		} `json:"attributes"`
	} `json:"scmProvider"`
	ScmRepository struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			HTTPCloneURL   string `json:"httpCloneUrl"`
			SSHCloneURL    string `json:"sshCloneUrl"`
			OwnerName      string `json:"ownerName"`
			RepositoryName string `json:"repositoryName"`
		} `json:"attributes"`
	} `json:"scmRepository"`
	ScmGitReference struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Name          string `json:"name"`
			CanonicalName string `json:"canonicalName"`
			IsDeleted     bool   `json:"isDeleted"`
			Kind          string `json:"kind"`
		} `json:"attributes"`
	} `json:"scmGitReference"`
}

type Discord struct {
	Username  string   `json:"username"`
	AvatarURL string   `json:"avatar_url"`
	Content   string   `json:"content"`
	Embeds    []Embeds `json:"embeds"`
}

type Field struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Embeds struct {
	Title       string  `json:"title"`
	URL         string  `json:"url"`
	Description string  `json:"description"`
	Color       int     `json:"color"`
	Fields      []Field `json:"fields"`
	Thumbnail   struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
	Footer struct {
		Text    string `json:"text"`
		IconURL string `json:"icon_url"`
	} `json:"footer"`
}

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
		var xcloud Xcloud
		err := json.Unmarshal(reqBody, &xcloud)
		if err != nil {
			fmt.Println(err)
			return
		}

		commitField := Field{
			Name:  "Commit Sha",
			Value: xcloud.CiBuildRun.Attributes.SourceCommit.CommitSha + "[See Changes](" + xcloud.CiBuildRun.Attributes.SourceCommit.HTMLURL + ")",
		}

		fields := []Field{commitField}

		embed := Embeds{
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

		embeds := []Embeds{embed}

		discord := Discord{
			Username:  "Xcode Cloud",
			AvatarURL: "https://developer.apple.com/assets/elements/icons/xcode-cloud/xcode-cloud-128x128_2x.png",
			Content:   "",
			Embeds:    embeds}

		// Convert discord into JSON
		// data, err := json.Marshal(discord)
		// if err != nil {
		// 	log.Fatal(err)
		// }
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
