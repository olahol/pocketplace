package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"net/http"
	"strings"
)

//go:generate file2const -package main files/index.html:indexHtml index_html.go
//go:generate file2const -package main files/frontend.js:frontendJs frontend_js.go

var port = flag.Int("port", 5000, "port to listen on")
var size = flag.Int("size", 1000, "canvas pixel size")
var cooldown = flag.Int("cooldown", 2, "drawing cooldown in seconds")

func RequestToId(req *http.Request) string {
	proxy := req.Header.Get("X-Forwarded-For")

	if proxy != "" {
		return proxy
	}

	addr := req.RemoteAddr
	parts := strings.Split(addr, ":")
	return parts[0]
}

func Authorize(req *http.Request) bool {
	return true
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()

	r := gin.New()
	m := melody.New()
	l := NewLimiter(*cooldown)
	c := NewCanvas(*size)

	r.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.String(200, indexHtml)
	})

	r.GET("/frontend.js", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/js")
		c.String(200, frontendJs)
	})

	optionsJs := fmt.Sprintf("var SIZE = %d;\n var COOLDOWN = %d;", *size, *cooldown)
	r.GET("/options.js", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/js")
		c.String(200, optionsJs)
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleConnect(func(s *melody.Session) {
		if !Authorize(s.Request) {
			return
		}

		s.Set("id", RequestToId(s.Request))

		c.RWMutex.RLock()
		s.WriteBinary(c.Data)
		c.RWMutex.RUnlock()
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		if !Authorize(s.Request) {
			return
		}

		cmd, err := ParseCmd(*size, msg)

		if err != nil {
			fmt.Println(err)
			return
		}

		id := s.MustGet("id").(string)

		if !l.Check(id) {
			return
		}

		l.Add(id)

		c.RWMutex.Lock()
		c.Set(cmd.x, cmd.y, cmd.r, cmd.g, cmd.b)
		c.RWMutex.Unlock()

		m.Broadcast(msg)
	})

	fmt.Printf("Canvas pixel size %dx%d\n", *size, *size)
	fmt.Printf("Drawing cooldown %ds\n", *cooldown)
	fmt.Printf("Listening on port %d\n", *port)
	r.Run(fmt.Sprintf(":%d", *port))
}
