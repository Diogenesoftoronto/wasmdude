package main

import "github.com/vugu/vgrouter"



type PageNotFound struct {
	vgrouter.NavigatorRef
	image string
}


func (c *PageNotFound) GetImage() string {
	


	return c.image
}