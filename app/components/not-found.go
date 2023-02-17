package components

import (
	"encoding/json"
	"log"
	"net/http"
	"path/filepath"

	"github.com/vugu/vugu"
)

type NotFound struct {
	Image  string `vugu:"data"`
	Loaded bool   `vugu:"data"`
}

type Response struct {
	Msg string `vugu:"data"`
}

func (c *NotFound) Init(ctx vugu.InitCtx) {
	// kick of loading data from an endpoint
	c.Loaded = false
	log.Print("Init")
	log.Printf("Context: %v", ctx.EventEnv())
	go c.getRandom404(ctx.EventEnv())
}

// This does not work currently because of the way the context is handled
// I need to figure out how to get the context to work properly
func (c *NotFound) getRandom404(ctx vugu.EventEnv) {
	// Currently this breaks, try to fix it
	log.Print("Fetching...")
	//TODO make this configurable using viper or something similar, could also use dotenv
	resp, err := http.Get("http://localhost:4001/main/Generate404")
	if err != nil {
		log.Printf("Error fetching: %v", err)
		return
	}
	defer resp.Body.Close()

	log.Print(resp.Body)

	var rsp map[string]string
	err = json.NewDecoder(resp.Body).Decode(&rsp)
	if err != nil {
		log.Printf("Error decoding response: %v", err)
		return
	}
	rsp["Message"], _ = filepath.Rel("/home/diogenesoft/goserver/wasmdude", rsp["Message"])
	log.Print(rsp["Message"])
	log.Print("Checking Context...")
	if ctx != nil {
		log.Printf("context is not nil %s", ctx)
		log.Print("Locking...")
		ctx.Lock()
		c.Image = "/" + rsp["Message"]
		c.Loaded = true
		log.Print(c.Image)
		log.Print("Unlocking...")
		ctx.UnlockRender()
	} else {
		log.Print("context is nil")
	}
}
