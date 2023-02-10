package main

import (
	"time"

	"github.com/vugu/vugu"
)

type Article struct {
	Title       string
	Thumbnail   string
	Text        string
	PublishDate time.Time
	Comments
	Loading   bool
	Engagement int
	Likes     int
}

func (c *Article) CalculateEngagement() {
	eng := len(c.Comments.Items) + len(c.Comments.Items)*2
	eng += c.Likes
	c.Engagement = eng
}

func (c *Article) Init(ctx vugu.InitCtx) {
	c.Loading = true
	go func() {
		ctx.EventEnv().Lock()
		defer ctx.EventEnv().UnlockRender()
		c.Loading = false
		c.CalculateEngagement()
	}()
}

func (c *Article) AddComment(text string) {
	c.Comments.Items = append(c.Comments.Items, &Comment{Content: text})
}

func (c *Article) AddLike() {
	c.Likes++
}

func (c *Article) RemoveLike() {
	c.Likes--
}

func (c *Article) Compute() {
	c.CalculateEngagement()
}




