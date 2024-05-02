package langsmithgo

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func TestMain(m *testing.M) {
	m.Run()

}
func TestRun(t *testing.T) {
	t.Run("use with GenerateFromSinglePrompt", func(t *testing.T) {

		// Create a new client
		client, err := NewClient()
		if err != nil {
			t.Errorf("Error creating client: %v", err)
		}
		prompt := "The first man to walk on the moon"
		llm, err := openai.New()
		if err != nil {
			t.Errorf("Error creating LLM: %v", err)
		}
		runId := uuid.New().String()
		ctx := context.Background()
		err = client.Run(&RunPayload{
			RunID:       runId,
			Name:        "langsmithgo-chain",
			SessionName: "langsmithgo",
			RunType:     Chain,
			Inputs: map[string]interface{}{
				"prompt": prompt,
			},
		})

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

		completion, err := llms.GenerateFromSinglePrompt(ctx,
			llm,
			prompt,
			llms.WithTemperature(0.8),
			llms.WithStopWords([]string{"Armstrong"}),
		)
		if err != nil {
			log.Fatalf("error generating completion: %v", err)
		}

		err = client.Run(&RunPayload{
			RunID: runId,
			Outputs: map[string]interface{}{
				"output": completion,
			},
		})

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

		fmt.Println(completion)
	})

	t.Run("use with Chain", func(t *testing.T) {
		runId := uuid.New().String()
		// Create a new client
		client, err := NewClient()
		if err != nil {
			t.Errorf("Error creating client: %v", err)
		}

		opts := []openai.Option{
			openai.WithModel("gpt-3.5-turbo"),
		}
		llm, err := openai.New(opts...)
		if err != nil {
			t.Errorf("Error creating LLM: %v", err)
		}
		ctx := context.Background()

		content := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, "You are a company branding design wizard."),
			llms.TextParts(llms.ChatMessageTypeHuman, "What would be a good company name a company that makes colorful socks?"),
		}

		err = client.Run(&RunPayload{
			RunID:       runId,
			Name:        "langsmithgo-llm",
			SessionName: "langsmithgo",
			RunType:     LLM,
			Tags:        []string{"llm"},
			Inputs: map[string]interface{}{
				"prompt":      content, // Ensure 'output' is properly defined and is of type that has a String method
				"model":       "gpt-3.5-turbo",
				"temperature": 0.7, // Assuming 'temperature' should be a float, not a string
			},
		})

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

		out, err := llm.GenerateContent(ctx, content)
		if err != nil {
			t.Errorf("Error running: %v", err)
		}
		err = client.Run(&RunPayload{
			RunID: runId,
			Outputs: map[string]interface{}{
				"output": out,
			},
		})

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

	})

	// use 2 chains
	t.Run("use with 2 traces", func(t *testing.T) {
		runId := uuid.New().String()
		// Create a new client
		client, err := NewClient()
		if err != nil {
			t.Errorf("Error creating client: %v", err)
		}

		opts := []openai.Option{
			openai.WithModel("gpt-3.5-turbo-0125"),
			openai.WithEmbeddingModel("text-embedding-3-large"),
		}
		llm, err := openai.New(opts...)
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		content := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, "You are a company branding design wizard."),
			llms.TextParts(llms.ChatMessageTypeHuman, "What would be a good company name a company that makes colorful socks?"),
		}

		err = client.Run(&RunPayload{
			RunID:       runId,
			Name:        "langsmithgo-llm",
			SessionName: "langsmithgo",
			RunType:     LLM,
			Tags:        []string{"llm"},
			Inputs: map[string]interface{}{
				"prompt":      content, // Ensure 'output' is properly defined and is of type that has a String method
				"model":       "gpt-3.5-turbo",
				"temperature": 0.7, // Assuming 'temperature' should be a float, not a string
			},
		})

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

		out, err := llm.GenerateContent(ctx, content)
		if err != nil {
			t.Errorf("Error running: %v", err)

		}

		err = client.Run(&RunPayload{
			RunID: runId,
			Outputs: map[string]interface{}{
				"output": out,
			},
		})

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

		runEmbedId := uuid.New().String()
		// create embedding
		err = client.Run(&RunPayload{
			RunID:       runEmbedId,
			ParentID:    runId,
			Name:        "langsmithgo-llm",
			SessionName: "langsmithgo",
			RunType:     Embedding,
			Tags:        []string{"llm"},
			Inputs: map[string]interface{}{
				"prompt":      out.Choices[0].Content, // Ensure 'output' is properly defined and is of type that has a String method
				"model":       "gpt-3.5-turbo",
				"temperature": 0.7, // Assuming 'temperature' should be a float, not a string
			},
		})

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

		embedings, err := llm.CreateEmbedding(ctx, []string{"ola", "mundo"})
		if err != nil {
			t.Errorf("Error running: %v", err)
		}

		err = client.Run(&RunPayload{
			RunID: runEmbedId,
			Outputs: map[string]interface{}{
				"output": embedings,
			},
		})

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

	})

	t.Run("use with 2 traces and SimpleRun", func(t *testing.T) {
		// Create a new client
		runId := uuid.New().String()

		client, err := NewClient()
		opts := []openai.Option{
			openai.WithModel("gpt-3.5-turbo-0125"),
			openai.WithEmbeddingModel("text-embedding-3-large"),
		}
		llm, err := openai.New(opts...)
		if err != nil {
			log.Fatal(err)
		}

		ctx := context.Background()

		content := []llms.MessageContent{
			llms.TextParts(llms.ChatMessageTypeSystem, "You are a company branding design wizard."),
			llms.TextParts(llms.ChatMessageTypeHuman, "What would be a good company name a company that makes colorful socks?"),
		}

		if err != nil {
			t.Errorf("Error running: %v", err)
		}
		startTime := time.Now().UTC()
		out, err := llm.GenerateContent(ctx, content)
		if err != nil {
			t.Errorf("Error running: %v", err)

		}
		runPayload := &RunPayload{
			RunID:       runId,
			Name:        "langsmithgo-llm",
			SessionName: "langsmithgo",
			RunType:     LLM,
			Tags:        []string{"llm"},
			StartTime:   startTime,
			Inputs: map[string]interface{}{
				"prompt":      content, // Ensure 'output' is properly defined and is of type that has a String method
				"model":       "gpt-3.5-turbo",
				"temperature": 0.7, // Assuming 'temperature' should be a float, not a string
			},
			Outputs: map[string]interface{}{
				"output": out,
			},
			EndTime: time.Now().UTC(),
		}

		err = client.RunSingle(runPayload)

		if err != nil {
			t.Errorf("Error running: %v", err)
		}

		// create embedding
		startTime = time.Now().UTC()

		embedings, err := llm.CreateEmbedding(ctx, []string{"ola", "mundo"})
		if err != nil {
			log.Fatal(err)
		}

		runEmbedId := uuid.New().String()
		err = client.RunSingle(&RunPayload{
			RunID:       runEmbedId,
			Name:        "langsmithgo-llm",
			SessionName: "langsmithgo",
			RunType:     Embedding,
			StartTime:   startTime,
			Tags:        []string{"llm"},
			Inputs: map[string]interface{}{
				"prompt":      out.Choices[0].Content, // Ensure 'output' is properly defined and is of type that has a String method
				"model":       "gpt-3.5-turbo",
				"temperature": 0.7, // Assuming 'temperature' should be a float, not a string
			},
			Outputs: map[string]interface{}{
				"output": embedings,
			},
			EndTime: time.Now().UTC(),
		})

		if err != nil {
			log.Fatal(err)
		}

	})
}
