// Code generated by github.com/firstcontributions/matro. DO NOT EDIT.

package storiesstore

import "time"

type Comment struct {
	StoryID         string    `bson:"story_id"`
	AbstractContent string    `bson:"abstract_content,omitempty"`
	ContentJson     string    `bson:"content_json,omitempty"`
	CreatedBy       string    `bson:"created_by,omitempty"`
	Id              string    `bson:"_id"`
	TimeCreated     time.Time `bson:"time_created,omitempty"`
	TimeUpdated     time.Time `bson:"time_updated,omitempty"`
}

func NewComment() *Comment {
	return &Comment{}
}

type CommentUpdate struct {
	TimeUpdated *time.Time `bson:"time_updated,omitempty"`
}

type CommentFilters struct {
	Ids []string

	Story *Story
}
