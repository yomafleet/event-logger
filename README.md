# Application Event Log

## Event Message

```golang
type EventMessage struct {
	Message string                 `json:"message"`
	Event   string                 `json:"event"`
	Context string                 `json:"context"`
	Data    map[string]interface{} `json:"data"`
}
```

### Message example

```golang
EventMessage {
	Message: "User identification updated"
	Event:   "profile.updated" // filterable in daashboard
	Context: "user" // use to populate loki lable (index)
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
