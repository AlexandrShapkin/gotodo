package main

import "time"

// task priority type
type Priority uint8

const (
	// highest priority of task
	HIGHEST Priority = iota
	// high priority of task
	HIGH
	// medium priority of task
	MEDIUM
	// low priority of task
	LOW
)

type Task struct {
	ID             int        `json:"id"`
	Name           string     `json:"name"`
	Description    string     `json:"description"`
	Status         string     `json:"status"`
	Priority       Priority   `json:"priority"`
	CreationTime   time.Time  `json:"creationTime"`
	Deadline       time.Time  `json:"deadline"`
	CompletionTime *time.Time `json:"completionTime,omitempty"`
}
