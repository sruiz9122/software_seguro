package repository

import (
	"context"

	"github-tracker/repository/entity"
	"github.com/stretchr/testify/mock"
)

type MockCommit struct {
	mock.Mock
}

var _ Commit = (*MockCommit)(nil)

func (m *MockCommit) Insert(ctx context.Context, commit *entity.Commit) error {
	results := m.Called(ctx, commit)
	return results.Error(0)
}

func (m *MockCommit) GetCommitByAuthorEmail(ctx context.Context, email string) ([]entity.Commit, error) {
	results := m.Called(ctx, email)
	commits, _ := results.Get(0).([]entity.Commit)
	return commits, results.Error(1)
}
