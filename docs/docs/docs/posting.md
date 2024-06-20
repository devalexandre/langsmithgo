---
sidebar_position: 2
sidebar_label: Posting
---

## Posting a New Run

The `PostRun` function might be used to create a new "run" record in the LangSmith system.
A "run" could represent a single execution instance of a language model, and using POST, you send all necessary details to initiate and log this run in the LangSmith backend.

```go
input := &langsmithgo.RunPayload{
    RunID:   "unique-run-id",
    Name:    "Example Run",
    RunType: langsmithgo.Tool,
    Inputs:  map[string]interface{}{"input_key": "input_value"},
    Tags:    []string{"example", "test"},
}

err := client.PostRun(input)
if err != nil {
    log.Fatalf("Failed to post run: %v", err)
}

```
