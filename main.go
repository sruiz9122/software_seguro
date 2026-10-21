package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github-tracker/models"
	"github-tracker/repository"
	"github-tracker/repository/entity"

	"github.com/gorilla/mux"
)

func insertGitHubWebhook(ctx context.Context, repo repository.Commit, webhook models.PushEvent) error {
	payload, err := json.Marshal(webhook)
	if err != nil {
		return fmt.Errorf("marshal GitHub webhook: %w", err)
	}

	commit := &entity.Commit{
		RepoName:       webhook.Repository.FullName,
		CommitID:       webhook.HeadCommit.ID,
		CommitMessage:  webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail:    webhook.HeadCommit.Author.Email,
		Payload:        string(payload),
		CreatedAt:      webhook.HeadCommit.Timestamp,
		UpdatedAt:      webhook.HeadCommit.Timestamp,
	}
	if err := repo.Insert(ctx, commit); err != nil {
		return fmt.Errorf("insert GitHub webhook commit: %w", err)
	}

	return nil
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received POST reques!")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	fmt.Println("Request Body:", string(body))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", postHandler).Methods("POST")
	fmt.Println("Server is running on http://localhost:8080")
	// Define your routes here
	log.Fatal(http.ListenAndServe(":8080", r))
}
