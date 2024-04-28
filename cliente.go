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

	payload := PostPayload{
		ID:          input.RunID,
		Name:        input.Name,
		RunType:     input.RunType,
		StartTime:   time.Now().UTC(),
		Inputs:      input.Inputs,
		SessionName: input.SessionName,
		Tags:        input.Tags,
		ParentId:    input.ParentID,
		Extras:      input.Extras,
		Metadata:    input.Metadata,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	err = c.Do(BASE_URL, http.MethodPost, jsonData)

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
