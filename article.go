package main

import "time"

type Article struct {
	Content   string
	Published time.Time
	Comments
	Loading   bool
	Engagment int
	Likes     int
}

// func (c Article) CalculateEngagment() {
// 	eng := len(Article.Comments.Items) + len(Article.Comments.Item)*2
// 	eng += Article.Likes
// 	Article.Engagment = eng
// }
