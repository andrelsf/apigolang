# API Games in GO 

I am learning `Golang`.

## Modules

```
    "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
```

## Entity

```python
type Game struct {
    Id          int       `json:"id"`
	Name        string    `json:"name"`
	Platform    string    `json:"platform"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	CreateAt    time.Time `json:"createat"`
	UpdateAt    time.Time `json:"updateat"`
}
```

## Endpoints

| __Method__  | __Endpoint__ | __Response__ |
|---------|:--------:|:----------|
| __GET__ | `/api/ping` | {StatusCode: 200, Message: "pong"} |