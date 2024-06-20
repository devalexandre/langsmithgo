---
slug: langsmithgo-integrating-llms-and-ai-tools-in-go-applications
title: "LangsmithGo: Integrating LLMs and AI Tools in Go Applications"
authors: [alexandre]
tags: [news]
---

# LangsmithGo: Integrating LLMs and AI Tools in Go Applications

## Introduction

Welcome to the official LangsmithGo project blog! Today, we will explore what LangsmithGo is, its benefits, and how it can be used to integrate language models (LLMs) and artificial intelligence (AI) tools into your Go applications. If you are a Go developer interested in AI, this article is for you.

## What is Langsmith?

Langsmith is a powerful platform designed to facilitate the integration of language models and AI tools into various applications. It offers a robust API that allows developers to create, manage, and monitor runs of language models, as well as support the addition of metadata, tools, and events. With Langsmith, you can leverage the potential of LLMs such as GPT-3 and GPT-3.5-turbo, among others, in your software solutions.

## LangsmithGo: The Go Library for Langsmith

LangsmithGo is the Go library that allows you to interact with the Langsmith API in a simple and efficient manner. It abstracts the complexities of communicating with the API, offering easy-to-use methods for creating and managing runs of LLMs and AI tools. With LangsmithGo, Go developers can quickly integrate advanced natural language processing capabilities into their applications.

## Benefits of LangsmithGo

1. Simple Integration: LangsmithGo offers an easy-to-use interface to integrate LLMs into Go applications.
2. Flexibility: Supports various types of runs, including LLMs, tools, execution chains, and more.
3. Monitoring and Management: Allows adding metadata, tools, and events to runs, facilitating monitoring and management.
4. OpenAI Compatibility: Native support for OpenAI models, such as GPT-3 and GPT-3.5-turbo.

## How to Use LangsmithGo

Let's explore some practical examples of how to use LangsmithGo to create and manage runs of LLMs and AI tools.

### Installation

```sh 
go get github.com/tmc/langsmithgo
```

### Configuration

```go 

import (
    "os"
)

func init() {
    os.Setenv("LANGSMITH_API_KEY", "your_api_key_here")
    os.Setenv("LANGSMITH_PROJECT_NAME", "your_project_name_here")
}

```

### Creating an LLM Run

Here is an example of how to create an LLM run with LangsmithGo:

```go

package main

import (
    "context"
    "fmt"
    "log"

    "github.com/google/uuid"
    "github.com/tmc/langchaingo/llms"
    "github.com/tmc/langchaingo/llms/openai"
    "github.com/tmc/langsmithgo"
)

func main() {
    client, err := langsmithgo.NewClient()
    if err != nil {
        log.Fatalf("Error creating client: %v", err)
    }

    llm, err := openai.New()
    if err != nil {
        log.Fatalf("Error creating LLM: %v", err)
    }

    runId := uuid.New().String()
    prompt := "The first man to walk on the moon"
    ctx := context.Background()

    err = client.Run(&langsmithgo.RunPayload{
        RunID:       runId,
        Name:        "example-llm-run",
        SessionName: "example-session",
        RunType:     langsmithgo.LLM,
        Inputs: map[string]interface{}{
            "prompt": prompt,
        },
    })

    if err != nil {
        log.Fatalf("Error running: %v", err)
    }

    completion, err := llms.GenerateFromSinglePrompt(ctx, llm, prompt)
    if err != nil {
        log.Fatalf("Error generating completion: %v", err)
    }

    err = client.Run(&langsmithgo.RunPayload{
        RunID: runId,
        Outputs: map[string]interface{}{
            "output": completion,
        },
    })

    if err != nil {
        log.Fatalf("Error running: %v", err)
    }

    fmt.Println(completion)
}

```

## Conclusion

LangsmithGo is a powerful tool for Go developers who want to integrate language models and AI tools into their applications. With its simple and flexible interface, you can quickly start creating, managing, and monitoring runs of LLMs, enriched with metadata, tools, and events. Try LangsmithGo today and elevate your Go applications to the next level!

Follow our blog for more tutorials, examples, and updates on LangsmithGo. If you have any questions or suggestions, feel free to contact us. Happy coding!