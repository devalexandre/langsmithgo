package langsmithgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/google/uuid"
)

// NewClient creates a new LangSmith client
// The client requires an API key to authenticate requests.
// You can get an API key by signing up for a LangSmith account at https://smith.langchain.com
// The API key can be passed as an argument to the function or set as an environment variable LANGSMITH_API_KEY
func NewClient() (*Client, error) {

	if os.Getenv("LANGSMITH_API_KEY") == "" {
		return nil, errors.New("langsmith api key is required")
	}

	url := os.Getenv("LANGSMITH_URL")
	if url == "" {
		url = BASE_URL

	}

	if os.Getenv("LANGSMITH_PROJECT_NAME") == "" {
		return nil, errors.New("langsmith project name is required")
	}

	return &Client{
		APIKey:      os.Getenv("LANGSMITH_API_KEY"),
		baseUrl:     fmt.Sprintf("%s/runs", url),
		projectName: os.Getenv("LANGSMITH_PROJECT_NAME"),
	}, nil
}

func (c *Client) PostRun(input *RunPayload) error {

	payload := PostPayload{
		ID:          input.RunID,
		Name:        input.Name,
		RunType:     input.RunType,
		StartTime:   time.Now().UTC(),
		Inputs:      input.Inputs,
		SessionName: c.projectName,
		Tags:        input.Tags,
		ParentId:    input.ParentID,
		Extras:      input.Extras,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = c.Do(c.baseUrl, http.MethodPost, jsonData)

	return err
}

func (c *Client) PatchRun(id string, input *RunPayload) error {
	payload := PatchPayload{
		Outputs: input.Outputs,
		EndTime: time.Now().UTC(),
		Events:  input.Events,
		Extras:  input.Extras,
		Error:   input.Error,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = c.Do(c.baseUrl+"/"+id, http.MethodPatch, jsonData)

	return err
}

func (c *Client) Run(input *RunPayload) error {

	uuidRegex := `^[a-f\d]{8}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{12}$`
	match, _ := regexp.MatchString(uuidRegex, input.RunID)
	if !match {
		return fmt.Errorf("invalid UUID format: %s", input.RunID)
	}

	if input.Outputs != nil {
		return c.PatchRun(input.RunID, input)
	}

	return c.PostRun(input)

}

/*
Single Requests
If you want to reduce the number of requests you make to LangSmith, you can log runs in a single request. Just be sure to include the outputs or error and fix the end_time all in the post request.
Below is an example that logs the completion LLM run from above in a single call.
*/

func (c *Client) RunSingle(input *RunPayload) error {

	payload := SimplePayload{
		ID:          input.RunID,
		Name:        input.Name,
		RunType:     input.RunType,
		StartTime:   input.StartTime,
		Inputs:      input.Inputs,
		SessionName: c.projectName,
		Tags:        input.Tags,
		ParentId:    input.ParentID,
		Extras:      input.Extras,
		Outputs:     input.Outputs,
		EndTime:     input.EndTime,
		Events:      input.Events,
	}

	if payload.ID == "" {
		payload.ID = uuid.New().String()
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	err = c.Do(c.baseUrl, http.MethodPost, jsonData)

	return err
}

// client for http requests
func (c *Client) Do(url string, method string, jsonData []byte) error {
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// print a response body for human readability

	// Check the response status
	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			return err
		}

		return errors.New(response.Detail)
	}

	return nil
}
