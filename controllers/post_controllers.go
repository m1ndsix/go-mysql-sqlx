package controllers

import (
	"fmt"
	"go-mysql-sqlx/db_client"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID        int64     `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Content   string    `json:"content" db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func CreatePost(c *gin.Context) {

	var reqBody Post
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	res, err := db_client.DBClient.Exec("INSERT INTO `test_db`.`posts` (`title`, `content`) VALUES (?,?);",
		reqBody.Title,
		reqBody.Content,
	)
	if err != nil {
		panic(err.Error())
	}

	// insert, err := db_client.DBClient.Query("INSERT INTO `test_db`.`posts` (`id`, `title`, `content`, `created_at`) VALUES ('1', 'First Note', 'No idea', '2016-11-23');")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("id: ", id)

	c.JSON(http.StatusCreated, gin.H{
		"error": false,
		"id":    id,
	})
}

func GetPosts(c *gin.Context) {
	var posts []Post

	// db_client.DBClient.Get() // SINGLE ROW (queryRow)
	db_client.DBClient.Select(&posts, "SELECT id, title, content, created_at FROM posts") // SLICE OF ROWS (query)

	c.JSON(http.StatusOK, posts)

	// sql
	// rows, err := db_client.DBClient.Query("SELECT id, title, content, created_at FROM posts")
	// if err != nil {
	// 	c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 		"error": true,
	// 	})
	// 	return
	// }

	// for rows.Next() {
	// 	var singlePost Post
	// 	if err := rows.Scan(&singlePost.ID, &singlePost.Title, &singlePost.Content, &singlePost.CreatedAt); err != nil {
	// 		c.JSON(http.StatusUnprocessableEntity, gin.H{
	// 			"error": true,
	// 		})
	// 		return
	// 	}
	// 	posts = append(posts, singlePost)
	// }
	// c.JSON(http.StatusOK, posts)
}

func GetPost(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	var post Post

	db_client.DBClient.Get(&post, "SELECT id, title, content, created_at FROM posts WHERE id = ?;", id)

	c.JSON(http.StatusOK, post)

	// sql
	// row := db_client.DBClient.QueryRow("SELECT id, title, content, created_at FROM posts WHERE id = ?;", id)
	// var post Post
	// if err := row.Scan(&post.ID, &post.Title, &post.Content, &post.CreatedAt); err != nil {
	// 	if err == sql.ErrNoRows {
	// 		c.JSON(http.StatusNotFound, gin.H{
	// 			"error": true,
	// 		})
	// 		return
	// 	}
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error":   true,
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }
	// c.JSON(http.StatusOK, post)
}
