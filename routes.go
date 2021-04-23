package main

import (
	"embed"
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
)

// embed assets resources for ui
//go:embed ui/dist/assets/*
var assetsDir embed.FS

//go:embed ui/dist/index.html
var indexPage []byte

const version string = "0.1.0"

// Routes API object
type Routes struct {
	Mux     *mux.Router
	Negroni *negroni.Negroni
	tunnels TunnelsMap
}

// NewRoutes create API backend for the program
func NewRoutes(td TunnelsMap) Routes {
	r := Routes{}
	r.tunnels = td
	return r
}

/* MIDDLEWARE */
func defaultHeaderMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, Authorization")
	next(rw, r)
}

/* API */
func message(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Here we GO!\n"))
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(version))
}

func (t *TunnelsMap) getTunnels(w http.ResponseWriter, r *http.Request) {
	data, err := json.Marshal(t)
	if err != nil {
		log.Println(err)
	}
	w.Write(data)
}

func (t *TunnelsMap) openTunnel(w http.ResponseWriter, r *http.Request) {}

func (t *TunnelsMap) closeTunnel(w http.ResponseWriter, r *http.Request) {}

/* WEBSOCKET */
func wsInteract(conn *websocket.Conn) {

	c1 := make(chan string)

	go func() {
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}
			in := string(p)
			// print out that message for clarity
			log.Println("received:" + in)
			c1 <- in
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			c1 <- "* WAKE-UP 2SEC *"
		}
	}()

	for {
		select {
		case msg := <-c1:
			if err := conn.WriteMessage(1, []byte(strings.ToUpper(msg))); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	//log.Println("Client Connected")
	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}
	// listen indefinitely for new messages coming
	// through on our WebSocket connection
	wsInteract(ws)
}

// Run API function
func (r *Routes) Run(addr string) {

	// Define mux
	r.Mux = mux.NewRouter()

	// API handler
	r.Mux.HandleFunc("/api/welcome", message).Methods("GET")
	r.Mux.HandleFunc("/api/version", getVersion).Methods("GET")

	// API tunnel handler
	r.Mux.HandleFunc("/api/tunnel", r.tunnels.getTunnels).Methods("GET")
	r.Mux.HandleFunc("/api/opentunnel", r.tunnels.openTunnel).Methods("POST")
	r.Mux.HandleFunc("/api/closetunnel", r.tunnels.closeTunnel).Methods("POST")

	// WebSocket handler
	r.Mux.HandleFunc("/api/ws", wsEndpoint)

	// UI handler
	assetsSubFs, _ := fs.Sub(fs.FS(assetsDir), "ui/dist")
	assetsFileServe := http.FileServer(http.FS(assetsSubFs))

	r.Mux.PathPrefix("/assets/").Handler(assetsFileServe)

	r.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write(indexPage)
	})

	// Define negroni middleware
	r.Negroni = negroni.New()
	r.Negroni.Use(negroni.HandlerFunc(defaultHeaderMiddleware))
	r.Negroni.UseHandler(r.Mux)

	log.Fatal(http.ListenAndServe(addr, r.Negroni))
}
