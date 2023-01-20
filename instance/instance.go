package instance

import (
	"sync"

	"github.com/r3labs/sse/v2"
)

type instance struct {
	sseServer *sse.Server
}

var singleton = &instance{}
var once sync.Once

func Init() {
	once.Do(func() {
		singleton.sseServer = sse.New()
	})
}

// Validator returns the validator
func SSEServer() *sse.Server {
	return singleton.sseServer
}
