package main

import (
	"fmt"
	"net/http"
	"web/ex"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	id := ex.Show{ID: 1}
	showid := id.GetShow()

	ids := ex.Hsuge{ID: 1, Title: "hyf"}
	hhh := ids.Eshw()
	fmt.Println(hhh)
	fmt.Print(showid)
	router := gin.Default()
	router.GET("/albums", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, albums)

	})

	router.POST("/albums", func(c *gin.Context) {
		var newAlbum album

		if err := c.BindJSON(&newAlbum); err != nil {
			return
		}
		albums = append(albums, newAlbum)
		c.IndentedJSON(http.StatusCreated, newAlbum)
	})

	router.GET("/albums/:id", func(c *gin.Context) {
		id := c.Param("id")

		for _, album := range albums {
			if album.ID == id {
				c.IndentedJSON(http.StatusOK, album)
				return
			}
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})

	})
	router.Run(":8888")
}

func ReverseRunes(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	// 12345
	return string(r)
}
