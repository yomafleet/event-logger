# Application Event Log

### Todo:

- [x] add json client
- [ ] add protobuf client
- [ ] add promtail client
- [ ] utilize concurrency mode for better performance

## Event Message

```golang
type EventMessage struct {
	Message string                 `json:"message"`
	Event   string                 `json:"event"`
	Type string                 `json:"type"`
	Data    map[string]interface{} `json:"data"`
}
```

### Message example

```golang
EventMessage {
	Message: "User identification updated"
	Event:   "profile.updated" // filterable in dashboard
	Type: "user" // to populate log lable or index
	Data:    map[string]interface{
        "id": 1,
        "name": "Admin",
        "email": "admin@gmail.com",
        "profile": {
            "id": 1,
            "user_id": 1,
            "image": "profile.png"
            "address": "Example",
        },
    }
}
```

### Config

```yaml
client: "json" # default client to use

service: "plus" # service name as lable or index

# @todo: add more clients
settings:
  json:
    url: "http://localhost:3100/loki/api/v1/push"
```

### Example

```golang
import "github.com/yomafleet/eventlogger"

// instantiate with config file
logger := elog.New("path/to/config.yaml")

// or get a new logger with a desire client after,
// only support `json` client for now
logger = logger.NewWithClient("json")

// add event message
logger.AddMessage(&{
	Message: "User identification updated"
	Event:   "profile.updated" // filterable in dashboard
	Type: "user" // to populate log lable or index
	Data:    map[string]interface{
        "id": 1,
        "name": "Admin",
        "email": "admin@gmail.com",
        "profile": {
            "id": 1,
            "user_id": 1,
            "image": "profile.png"
            "address": "Admin Address",
        },
    }
})

// adding message with same Type and Event will be grouped
logger.AddMessage(&{
	Message: "User identification updated"
	Event:   "profile.updated" // filterable in dashboard
	Type: "user" // to populate log lable or index
	Data:    map[string]interface{
        "id": 2,
        "name": "Editor",
        "email": "editor@gmail.com",
        "profile": {
            "id": 2,
            "user_id": 2,
            "image": "profile.png"
            "address": "Editro Address",
        },
    }
})

// send all messages
logger.Send(nil)
```

Alternatively, you can send a single message in one go

```golang
import "github.com/yomafleet/eventlogger"

// instantiate with config file
logger := elog.New("path/to/config.yaml")
logger.Send(&{
	Message: "User identification updated"
	Event:   "profile.updated" // filterable in dashboard
	Type: "user" // to populate log lable or index
	Data:    map[string]interface{
        "id": 2,
        "name": "Editor",
        "email": "editor@gmail.com",
        "profile": {
            "id": 2,
            "user_id": 2,
            "image": "profile.png"
            "address": "Editro Address",
        },
    }
})
```

**NOTE**: after `Send` has been called, all the messages in streams data will be flushed.
