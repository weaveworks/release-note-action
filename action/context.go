package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/caarlos0/env"
)

// Context represents the trigger event context
type Context struct {
	Repository string `json:"repo" env:"GITHUB_REPOSITORY"`
	SHA        string `json:"sha" env:"GITHUB_SHA"`
	Ref        string `json:"ref" env:"GITHUB_REF"`
	EventName  string `json:"eventName" env:"GITHUB_EVENT_NAME"`
	Workflow   string `json:"workflow" env:"GITHUB_WORKFLOW"`
	Action     string `json:"action" env:"GITHUB_ACTION"`
	Actor      string `json:"actor" env:"GITHUB_ACTOR"`
	Job        string `json:"job" env:"GITHUB_JOB"`
	RunNumber  int    `json:"runNumber" env:"GITHUB_RUN_NUMBER"`
	RunID      int    `json:"runID" env:"GITHUB_RUN_ID"`
	APIURL     string `json:"apiUrl" env:"GITHUB_API_URL" envDefault:"ttps://api.github.com"`
	ServerURL  string `json:"serverUrl" env:"GITHUB_SERVER_URL" envDefault:"https://github.com"`
	GraphQLUrl string `json:"graphqlUrl" env:"GITHUB_GRAPHQL_URL " envDefault:"https://api.github.com/graphql"`
	Event      *Event `json:"payload"`
	EventPath  string `env:"GITHUB_EVENT_PATH"`
}

// NewContextFromEnv creates a new context from environment variables
func NewContextFromEnv() (*Context, error) {
	ghContext := &Context{}
	if err := env.Parse(ghContext); err != nil {
		return nil, fmt.Errorf("parsing environments variables: %w", err)
	}
	if err := payloadFromFile(ghContext); err != nil {
		return nil, fmt.Errorf("loading event from file: %w", err)
	}

	return ghContext, nil
}

func payloadFromFile(ghCtx *Context) error {
	if ghCtx.EventPath == "" {
		return nil
	}

	data, err := ioutil.ReadFile(ghCtx.EventPath)
	if err != nil {
		return fmt.Errorf("reading event file %s: %w", ghCtx.EventPath, err)
	}

	evt := &Event{}
	if err := json.Unmarshal(data, evt); err != nil {
		return fmt.Errorf("unmarshalling event file: %w", err)
	}
	ghCtx.Event = evt

	return nil
}
