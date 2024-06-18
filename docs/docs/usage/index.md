---
sidebar_position: 1
sidebar_label: Setting Up
---

# Setting Up


## Installation

```sh 
go get github.com/devalexandre/langsmithgo
```


First, create a new client. The client requires an API key to authenticate requests. You can get an API key by signing up for a LangSmith account at LangSmith. The API key can be passed as an argument to the function or set as an environment variable `LANGSMITH_API_KEY`.

```go
package main

import (
    "log"
    "github.com/devalexandre/langsmithgo"
)

func main() {
    client, err := langsmithgo.NewClient()
    if err != nil {
        log.Fatalf("Failed to create LangSmith client: %v", err)
    }

    // Use the client for various operations...
}
```