package langsmithgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

func NewClient(apiKey string) *Client {
	apikey := apiKey
	if apikey == "" {
		apikey = os.Getenv("LANGSMITH_API_KEY")
	}
	return &Client{
		APIKey: apiKey,
	}
}

func (c *Client) PostRun(input *RunPayload) error {

	SetRunId(uuid.New().String())

	payload := PostPayload{
		ID:          GetRunId(),
		Name:        input.Name,
		RunType:     input.RunType,
		StartTime:   time.Now().UTC(),
		Inputs:      input.Inputs,
		SessionName: input.SessionName,
		Tags:        input.Tags,
		ParentId:    GetParentId(),
		Extras:      input.Extras,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = c.Do(BASE_URL, http.MethodPost, jsonData)

	SetParentId(GetRunId())

	return err
}

func (c *Client) PatchRun(id string, input *RunPayload) error {
	payload := PatchPayload{
		Outputs: input.Outputs,
		EndTime: time.Now().UTC(),
		Events:  input.Events,
	}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = c.Do(BASE_URL+"/"+id, http.MethodPatch, jsonData)

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
	SetRunId(uuid.New().String())
	payload := SimplePayload{
		PostPayload: PostPayload{ // Especificando a struct PostPayload
			ID:          GetRunId(),
			Name:        input.Name,
			RunType:     input.RunType,
			StartTime:   input.StartTime,
			Inputs:      input.Inputs,
			SessionName: input.SessionName,
			Tags:        input.Tags,
			ParentId:    GetParentId(),
			Extras:      input.Extras,
		},
		PatchPayload: PatchPayload{ // Especificando a struct PatchPayload
			Outputs: input.Outputs,
			EndTime: input.EndTime,
			Events:  input.Events,
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = c.Do(BASE_URL, http.MethodPost, jsonData)

	SetParentId(GetRunId())
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
