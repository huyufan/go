package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var wu = &websocket.Upgrader{ReadBufferSize: 512, WriteBufferSize: 512, CheckOrigin: func(r *http.Request) bool { return true }}

func myws(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{ws: ws, sc: make(chan []byte, 256), data: &Data{}}
	fmt.Println(c)
	h.r <- c

	go c.writer()
	c.reader()

	defer func() {
		c.data.Type = "loginout"
		user_list = del(user_list, c.data.User)
		c.data.UserList = user_list
		c.data.Content = c.data.User
		data_b, _ := json.Marshal(c.data)
		h.b <- data_b
		h.r <- c
	}()

}

func local(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "local.html")
}

func (c *connection) writer() {

	for message := range c.sc {
		fmt.Println("12")
		c.ws.WriteMessage(websocket.TextMessage, message)
	}
	c.ws.Close()
}

var user_list = []string{}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(message, &c.data)
		switch c.data.Type {
		case "login":
			c.data.User = c.data.Content
			c.data.From = c.data.User
			user_list = append(user_list, c.data.User)
			c.data.UserList = user_list
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "user":
			c.data.Type = "user"
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
		case "loginout":
			c.data.Type = "loginout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c

		default:
			fmt.Println("================================default================================")
		}

	}
}

func del(slice []string, user string) []string {
	count := len(slice)
	fmt.Printf("count%s", count)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var nslice = []string{}
	for i := range slice {
		fmt.Printf("i%s", i)
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			nslice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(nslice)
	return nslice
}
