package main

import (
	pq "github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/emirpasic/gods/utils"
)

type Articles struct {
	Items []*Article
	Topic  string
	Order *pq.Queue
}

func byEngagment(a, b interface{}) int {
	engagementA := a.(Article).Engagement
	engagementB := b.(Article).Engagement
	return -utils.IntComparator(engagementA, engagementB) // "-" descending order
}

func NewArticles() *Articles {
	return &Articles{
		Order: pq.NewWith(byEngagment),
	}
}

func (a *Articles) Init() {
	a.Order = pq.NewWith(byEngagment)
	a.SortByEngagement()
}

func (a *Articles) SortByEngagement() {
	for _, item := range a.Items {
		a.Order.Enqueue(item)
	}
}

func (a *Articles) AddArticle(art *Article) {
	a.Order.Enqueue(art)
}
