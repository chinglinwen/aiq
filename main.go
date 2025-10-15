package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	modelFlag    = flag.String("m", "gemini-2.5-flash", "model version (e.g., gemini-pro, gemini-2.5-flash, gemini-2.5-pro)")
	providerFlag = flag.String("p", "gemini", "AI provider (currently only gemini is supported)")
	systemFlag   = flag.String("s", "", "system instruction to guide the AI's behavior")
	shellFlag    = flag.Bool("shell", false, "output only shell commands (for piping to shell)")
	jsonFlag     = flag.Bool("json", false, "output response in JSON format")
)

func main() {
	// Load .env file if it exists (ignore errors if file doesn't exist)
	_ = godotenv.Load()

	flag.Parse()

	// Get the query prompt from arguments
	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: query prompt is required")
		fmt.Fprintln(os.Stderr, "Usage: aiq [flags] <query prompt>")
		flag.PrintDefaults()
		os.Exit(1)
	}
	queryPrompt := strings.Join(args, " ")

	// Read data from stdin if available
	var stdinData string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Data is being piped to stdin
		stdinBytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading stdin: %v\n", err)
			os.Exit(1)
		}
		stdinData = string(stdinBytes)
	}

	// Execute the query based on provider
	switch *providerFlag {
	case "gemini":
		if err := queryGemini(queryPrompt, stdinData, *modelFlag, *systemFlag, *shellFlag, *jsonFlag); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Error: unsupported provider '%s'\n", *providerFlag)
		os.Exit(1)
	}
}

func queryGemini(prompt, data, model, systemInstruction string, shellMode, jsonMode bool) error {
	ctx := context.Background()

	// Get API key from environment
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_API_KEY")
	}
	if apiKey == "" {
		return fmt.Errorf("GEMINI_API_KEY or GOOGLE_API_KEY environment variable not set")
	}

	// Create client
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return fmt.Errorf("failed to create client: %w", err)
	}
	defer client.Close()

	// Get the model
	genModel := client.GenerativeModel(model)

	// Build combined system instruction
	var instructions []string
	if systemInstruction != "" {
		instructions = append(instructions, systemInstruction)
	}
	if shellMode {
		shellInstruction := "You are a shell command generator. Output ONLY the shell command(s) needed, with no explanations, no markdown, no code blocks, no additional text. Just the raw command(s) that can be directly piped to a shell."
		instructions = append(instructions, shellInstruction)
	}
	if jsonMode {
		jsonInstruction := "You must respond with valid JSON only. Do not include markdown code blocks, explanations, or any text outside the JSON structure. Output raw JSON that can be directly parsed."
		instructions = append(instructions, jsonInstruction)
	}

	// Set combined system instruction if any instructions exist
	if len(instructions) > 0 {
		combinedInstruction := strings.Join(instructions, " ")
		genModel.SystemInstruction = &genai.Content{
			Parts: []genai.Part{genai.Text(combinedInstruction)},
		}
	}

	// Build the full prompt
	fullPrompt := prompt
	if data != "" {
		fullPrompt = fmt.Sprintf("%s\n\nData:\n%s", prompt, data)
	}

	// Generate content
	resp, err := genModel.GenerateContent(ctx, genai.Text(fullPrompt))
	if err != nil {
		return fmt.Errorf("failed to generate content: %w", err)
	}

	// Print the response
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Printf("%v", part)
			}
		}
	}
	fmt.Println()

	return nil
}
