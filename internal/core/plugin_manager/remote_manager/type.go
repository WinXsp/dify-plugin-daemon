package remote_manager

import (
	"sync"
	"time"

	"github.com/langgenius/dify-plugin-daemon/internal/types/entities"
	"github.com/langgenius/dify-plugin-daemon/internal/utils/stream"
	"github.com/panjf2000/gnet/v2"
)

type RemotePluginRuntime struct {
	entities.PluginRuntime

	// connection
	conn gnet.Conn

	// response entity to accept new events
	response *stream.StreamResponse[[]byte]

	// callbacks for each session
	callbacks      map[string][]func([]byte)
	callbacks_lock *sync.RWMutex

	// channel to notify all waiting routines
	shutdown_chan chan bool

	// heartbeat
	last_active_at time.Time

	// hand shake process completed
	handshake bool

	alive bool
}

// Listen creates a new listener for the given session_id
// session id is an unique identifier for a request
func (r *RemotePluginRuntime) addCallback(session_id string, fn func([]byte)) {
	r.callbacks_lock.Lock()
	if _, ok := r.callbacks[session_id]; !ok {
		r.callbacks[session_id] = make([]func([]byte), 0)
	}
	r.callbacks[session_id] = append(r.callbacks[session_id], fn)
	r.callbacks_lock.Unlock()
}

// removeCallback removes the listener for the given session_id
func (r *RemotePluginRuntime) removeCallback(session_id string) {
	r.callbacks_lock.Lock()
	delete(r.callbacks, session_id)
	r.callbacks_lock.Unlock()
}

func (r *RemotePluginRuntime) close() {
	// close shutdown channel to notify all waiting routines
	close(r.shutdown_chan)

	// close response to stop current plugin
	r.response.Close()
}