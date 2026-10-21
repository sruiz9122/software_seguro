package models

import "time"

type PushEvent struct {
	Repository Repository `json:"repository"`
	HeadCommit Commit     `json:"head_commit"`
}

type Repository struct {
	FullName string `json:"full_name"`
}

type Commit struct {
	ID        string       `json:"id"`
	Message   string       `json:"message"`
	Timestamp time.Time    `json:"timestamp"`
	Author    CommitPerson `json:"author"`
}

type CommitPerson struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
