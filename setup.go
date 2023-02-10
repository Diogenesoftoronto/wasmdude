package main

import (
	"log"
	"path"
	"strings"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
	"github.com/vugu/vugu/js"
)

// nolint: unused
func vuguSetup(buildEnv *vugu.BuildEnv, eventEnv vugu.EventEnv) vugu.Builder {

	router := vgrouter.New(eventEnv)

	// if there is a fragment when the page is loaded we go into fragment mode
	if strings.HasPrefix(js.Global().Get("window").Get("location").Get("hash").String(), "#") {
		router.SetUseFragment(true)
	} else {
		// otherwise we set the path prefix
		browserPath := path.Clean("/" + js.Global().Get("window").Get("location").Get("pathname").String())
		pathPrefix := "/" + strings.Split(strings.TrimPrefix(browserPath, "/"), "/")[0]
		log.Print(pathPrefix)
	}

	buildEnv.SetWireFunc(func(b vugu.Builder) {
		if c, ok := b.(vgrouter.NavigatorSetter); ok {
			c.NavigatorSet(router)
		}
	})

	root := &Root{}
	buildEnv.WireComponent(root)

	router.MustAddRouteExact("/", vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
		root.Body = &Home{}
	}))
	router.MustAddRouteExact("contact-me", vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
		root.Body = &ContactMe{}
	}))
	router.MustAddRouteExact("about", vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
		root.Body = &AboutMe{}
	}))
	router.MustAddRouteExact("my-projects", vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
		root.Body = &MyProjects{}
	}))
	router.MustAddRouteExact("my-articles", vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
		root.Body = &MyArticles{}
	}))
	router.SetNotFound(vgrouter.RouteHandlerFunc(func(rm *vgrouter.RouteMatch) {
		root.Body = &PageNotFound{}
	}))

	err := router.ListenForPopState()
	if err != nil {
		panic(err)
	}

	err = router.Pull()
	if err != nil {
		panic(err)
	}

	return root
}

func linkFor(p string) string {
	return p
}
