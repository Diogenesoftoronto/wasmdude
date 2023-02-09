package main

import "time"

type Comment struct {
	Content   string
	Author    string
	Published time.Time
	Likes     int
	Engagment int
	Replies []Comment
}

func (c Comment) CalculateEngagment() {
	eng := len(c.Replies) + len(c.Replies)*2
	eng += c.Likes
	c.Engagment = eng
}
