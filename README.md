# socket

Functional socket communication for Go.

## Installation

With Google's [Go](http://www.golang.org) installed on your machine:

    $ go get -u github.com/chuckpreslar/socket

## Usage



```go
package main

import (
  "bytes"
  "fmt"
)

import (
  "github.com/chuckpreslar/socket"
)

func main() {
  socket.Connect("localhost", 3000, socket.TCP, function(connection *socket.Socket) {
    connection.On(socket.Data, func(buffer *bytes.Buffer) {
      fmt.Printf("Received %s\n", buffer)
    }).On(socket.Error, func(err error) {
      fmt.Printf("Error %s\n", err)
    }).On(socket.Close, func() {
      fmt.Print("Socket connection closed\n")
    })
  })
}
```

## Documentation

View godoc or visit [godoc.org](http://godoc.org/github.com/chuckpreslar/socket).

    $ godoc socket

## License

> The MIT License (MIT)

> Copyright (c) 2013 Chuck Preslar

> Permission is hereby granted, free of charge, to any person obtaining a copy
> of this software and associated documentation files (the "Software"), to deal
> in the Software without restriction, including without limitation the rights
> to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
> copies of the Software, and to permit persons to whom the Software is
> furnished to do so, subject to the following conditions:

> The above copyright notice and this permission notice shall be included in
> all copies or substantial portions of the Software.

> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
> IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
> FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
> AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
> LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
> OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
> THE SOFTWARE.
