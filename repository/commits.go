package repository

import (
	"context"
	"database/sql"
	"errors"

	"github-tracker/repository/entity"
)

type Commit interface {
	Insert(ctx context.Context, commit *entity.Commit) (err error)
	GetCommitByAuthorEmail(ctx context.Context, email string) (commits []entity.Commit, err error)
}

type SQLCommitRepository struct {
	Conn *sql.DB
}

func NewCommit(conn *sql.DB) *SQLCommitRepository {
	return &SQLCommitRepository{Conn: conn}
}

func (m *SQLCommitRepository) Insert(ctx context.Context, commit *entity.Commit) error {
	if m == nil || m.Conn == nil {
		return errors.New("commit repository database connection is nil")
	}
	if commit == nil {
		return errors.New("commit is nil")
	}

	const query = `
		INSERT INTO commits (
			repo_name,
			commit_id,
			commit_message,
			author_username,
			author_email,
			payload,
			created_at,
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx,
		commit.RepoName,
		commit.CommitID,
		commit.CommitMessage,
		commit.AuthorUsername,
		commit.AuthorEmail,
		commit.Payload,
		commit.CreatedAt,
		commit.UpdatedAt,
	)
	return err
}

func (m *SQLCommitRepository) GetCommitByAuthorEmail(ctx context.Context, email string) ([]entity.Commit, error) {
	if m == nil || m.Conn == nil {
		return nil, errors.New("commit repository database connection is nil")
	}
	if email == "" {
		return nil, errors.New("author email is empty")
	}

	const query = `
		SELECT
			id,
			repo_name,
			commit_id,
			commit_message,
			author_username,
			author_email,
			payload,
			created_at,
			updated_at
		FROM commits
		WHERE author_email = $1
		ORDER BY created_at DESC`

	rows, err := m.Conn.QueryContext(ctx, query, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	commits := make([]entity.Commit, 0)
	for rows.Next() {
		var commit entity.Commit
		if err := rows.Scan(
			&commit.ID,
			&commit.RepoName,
			&commit.CommitID,
			&commit.CommitMessage,
			&commit.AuthorUsername,
			&commit.AuthorEmail,
			&commit.Payload,
			&commit.CreatedAt,
			&commit.UpdatedAt,
		); err != nil {
			return nil, err
		}
		commits = append(commits, commit)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return commits, nil
}
