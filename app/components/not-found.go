package components

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"github.com/vugu/vgrouter"
	"github.com/vugu/vugu"
)

type PageNotFound struct {
	vgrouter.NavigatorRef
	Image   string `vugu:"data"`
	Loading bool   `vugu:"data"`
}

type Response struct {
	Msg string `vugu:"data"`
}

func (c *PageNotFound) Init(ctx vugu.InitCtx) {
	// kick of loading data from an endpoint

	c.Loading = true
	go func() {
		defer ctx.EventEnv().UnlockRender()

		resp, err := http.Get("http://localhost:4001/main/Generate404")
		if err != nil {
			log.Printf("Error fetching: %v", err)
			return
		}
		log.Print(resp.Body)
		defer resp.Body.Close()
		var rsp map[string]string
		err = json.NewDecoder(resp.Body).Decode(&rsp)
		if err != nil {
			log.Printf("Error decoding response: %v", err)
			return
		}

		ctx.EventEnv().Lock()
		c.Loading = false
		rsp["Message"], _ = filepath.Rel("/home/diogenesoft/goserver/wasmdude", rsp["Message"])
		c.Image = "/" + rsp["Message"]
		log.Print(c.Image)
	}()
}
