// Code generated by github.com/firstcontributions/matro. DO NOT EDIT.

package usersstore

type CursorCheckpoints struct {
	PullRequests string `bson:"pull_requests,omitempty"`
}

func NewCursorCheckpoints() *CursorCheckpoints {
	return &CursorCheckpoints{}
}

type CursorCheckpointsFilters struct {
	Ids []string
}
