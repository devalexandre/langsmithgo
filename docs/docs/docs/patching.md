---
sidebar_position: 3
sidebar_label: Patching
---

## Patching an Existing Run

The `PatchRun` function allows updating an existing run, perhaps to add results, change status, or log additional data as the run progresses. For instance, if a run is completed and you want to update it with output data or change its status, you would use PATCH to modify only those specific fields without touching the rest of the run's data.
```go
patchInput := &langsmithgo.RunPayload{
    RunID:    "unique-run-id",
    Outputs:  map[string]interface{}{"output_key": "output_value"},
    Events:   []langsmithgo.Event{{EventName: "event1", Value: "value1"}},
}

err := client.PatchRun("unique-run-id", patchInput)
if err != nil {
    log.Fatalf("Failed to patch run: %v", err)
}
```