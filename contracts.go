package langsmithgo

import (
	"bytes"
	"time"
)

const (
	BASE_URL = "https://api.smith.langchain.com/runs"
)

type Response struct {
	Detail string `json:"detail"`
}
type Event struct {
	EventName string `json:"event_name"`
	Reason    string `json:"reason",omitempty`
	Value     string `json:"value",omitempty`
}

type RunPayload struct {
	RunID       string                 `json:"id"`
	Name        string                 `json:"name"`
	RunType     RunType                `json:"run_type"`
	StartTime   time.Time              `json:"start_time"`
	Inputs      map[string]interface{} `json:"inputs"`
	ParentID    string                 `json:"parent_run_id"`
	SessionName string                 `json:"session_name"`
	Tags        []string               `json:"tags"`
	Outputs     map[string]interface{} `json:"outputs"`
	EndTime     time.Time              `json:"end_time"`
	Extras      map[string]interface{} `json:"extras"`
	Events      []Event                `json:"events"`
}

type Client struct {
	APIKey string
}

type SimplePayload struct {
	PostPayload
	PatchPayload
}

type PostPayload struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	RunType     RunType                `json:"run_type"`
	StartTime   time.Time              `json:"start_time"`
	Inputs      map[string]interface{} `json:"inputs"`
	SessionName string                 `json:"session_name"`
	Tags        []string               `json:"tags,omitempty"`
	ParentId    string                 `json:"parent_run_id,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

type PatchPayload struct {
	Outputs map[string]interface{} `json:"outputs"`
	EndTime time.Time              `json:"end_time"`
	Events  []Event                `json:"events,omitempty"`
}

type RunType int

// Enum values using iota
const (
	Tool RunType = iota
	Chain
	LLM
	Retriever
	Embedding
	Prompt
	Parser
)

// runTypeNames maps RunType values to their string representations
var runTypeNames = []string{"tool", "chain", "llm", "retriever", "embedding", "prompt", "parser"}

// String returns the string representation of the RunType
func (r RunType) String() string {
	if int(r) < len(runTypeNames) {
		return runTypeNames[r]
	}
	return "unknown"
}

// MarshalJSON customizes the JSON output for RunType
func (r RunType) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(r.String())
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}
