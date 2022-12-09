package model

import "time"

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
			StartedDate  string `json:"startedDate"`
			FinishedDate string `json:"finishedDate"`
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