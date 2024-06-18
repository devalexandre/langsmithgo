---
sidebar_position: 3
sidebar_label:  Single Request
---

## Single Request
The `RunSingle` function appears to offer a way to handle a run operation where both initial data and completion data (like outputs and final status) are known upfront and can be sent together. This method might be useful for cases where the operation's duration is short or its outcome can be immediately determined, allowing for more efficient data processing and fewer API calls.
```go
singleRunInput := &langsmithgo.RunPayload{
    RunID:   "unique-run-id",
    Name:    "Single Request Run",
    RunType: langsmithgo.Tool,
    Inputs:  map[string]interface{}{"input_key": "input_value"},
    Outputs: map[string]interface{}{"output_key": "output_value"},
    Tags:    []string{"single", "test"},
}

err := client.RunSingle(singleRunInput)
if err != nil {
    log.Fatalf("Failed to run single request: %v", err)
}

```