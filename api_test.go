package gogpt

import (
	"context"
	"io/ioutil"
	"testing"
)

func TestAPI(t *testing.T) {
	tokenBytes, err := ioutil.ReadFile(".openai-token")
	if err != nil {
		t.Fatalf("Could not load auth token from .openai-token file")
	}

	c := NewClient(string(tokenBytes))
	ctx := context.Background()
	_, err = c.ListEngines(ctx)
	if err != nil {
		t.Fatalf("ListEngines error: %v", err)
	}

	_, err = c.GetEngine(ctx, "davinci")
	if err != nil {
		t.Fatalf("GetEngine error: %v", err)
	}

	req := CompletionRequest{MaxTokens: 5}
	req.Prompt = "Lorem ipsum"
	_, err = c.CreateCompletion(ctx, "ada", req)
	if err != nil {
		t.Fatalf("CreateCompletion error: %v", err)
	}

	searchReq := SearchRequest{
		Documents: []string{"White House", "hospital", "school"},
		Query:     "the president",
	}
	_, err = c.Search(ctx, "ada", searchReq)
	if err != nil {
		t.Fatalf("Search error: %v", err)
	}
}
