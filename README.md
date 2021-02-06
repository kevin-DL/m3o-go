# M3O Go Client

By default the client connects to api.m3o.com/client

```go
package main

import (
    "fmt"
    "os"

    "github.com/m3o/m3o-go/client"
)

type Request struct {
	Name string `json:"name"`
}

type Response struct {
	Msg string `json:"msg"`
}

var (
	token, _ = os.Getenv("TOKEN")
)

func main() {
	c := client.NewClient(nil)

	// set your api token
	c.SetToken(token)

   	req := &Request{
		Name: "John",
	}
	var rsp Response

	if err := c.Call("greeter", "Say.Hello", req, &rsp); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp)
}
```

If you want to access your local micro:

```go
    c := client.NewClient(client.Options{Local: true})
```

You can also set the api address explicitly:

```go
    c := client.NewClient(client.Options{Address: "https://api.yourdomain.com/client"})
```

## Streaming

The client supports streaming

```go
package main

import (
	"fmt"

	"github.com/m3o/m3o-go/client"
)

type Request struct {
	Count string `json:"count"`
}

type Response struct {
	Count string `json:"count"`
}

func main() {
	c := client.NewClient(&client.Options{Local: true})

	stream, err := c.Stream("example", "Streamer.ServerStream", Request{Count: "10"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		var rsp Response
		if err := stream.Recv(&rsp); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("got", rsp.Count)
	}
}
```
