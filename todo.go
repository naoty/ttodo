package main

import "time"

type todo struct {
	title       string
	description string
	assignee    string
	deadline    time.Time
	done        bool
}
