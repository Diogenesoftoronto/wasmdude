package components

import (
	"time"

	"github.com/vugu/vugu"
)

type Comment struct {
	Content   string
	Author    string
	Published time.Time
	Likes     int
	Replies   []Comment
	Loading   bool
	vugu.Builder
}

func (c *Comment) AddLike() {
	c.Likes++
}

func (c *Comment) RemoveLike() {
	c.Likes--
}
