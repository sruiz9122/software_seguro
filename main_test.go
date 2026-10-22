package main

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github-tracker/models"
	"github-tracker/repository"
	"github-tracker/repository/entity"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestDummy(t *testing.T) {
	assertions := require.New(t)

	result := 22

	assertions.Equal(22, result)
}

func TestInsert(t *testing.T) {
	assertions := require.New(t)
	timestamp := time.Date(2026, time.October, 21, 12, 0, 0, 0, time.UTC)
	webhook := models.PushEvent{
		Repository: models.Repository{
			FullName: "camilalenis/secure-dev",
		},
		HeadCommit: models.Commit{
			ID:        "9da3ed5d64d14d46dd1401d0768bc9dde90e86e1cb",
			Message:   "Add sample code for handle-github-webhook",
			Timestamp: timestamp,
			Author: models.CommitPerson{
				Email:    "mlmariacami2@gmail.com",
				Username: "camilalenis",
			},
		},
	}
	payload, err := json.Marshal(webhook)
	assertions.NoError(err)

	expectedCommit := &entity.Commit{
		RepoName:       webhook.Repository.FullName,
		CommitID:       webhook.HeadCommit.ID,
		CommitMessage:  webhook.HeadCommit.Message,
		AuthorUsername: webhook.HeadCommit.Author.Username,
		AuthorEmail:    webhook.HeadCommit.Author.Email,
		Payload:        string(payload),
		CreatedAt:      timestamp,
		UpdatedAt:      timestamp,
	}

	commitRepository := new(repository.MockCommit)
	commitRepository.On("Insert", mock.Anything, expectedCommit).Return(nil).Once()

	err = insertGitHubWebhook(context.Background(), commitRepository, webhook)

	assertions.NoError(err)
	commitRepository.AssertExpectations(t)
}
