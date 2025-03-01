package comment

import (
	"fmt"
	"go-simple-rest/src/v1/comment/model"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func New(c *gin.Context) {
	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	var comment model.Comment
	if err := c.BindJSON(&comment); err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": "Please provide valid credntials"})
		return
	}
	err, _ = NewComment(comment, articleId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Comment saved"})
}

func Show(c *gin.Context) {
	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	commentId, _ := primitive.ObjectIDFromHex(c.Param("cid"))

	err, res := GetComment(articleId, commentId)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "comment", "data": res})
}

// Comment godoc
// @Tags Articles
// @Summary Get article comments
// @Description Retrieves comments for a specific article.
// @Param id path string true "Article ID"
// @Param limit query int true "Limit"
// @Param prev query int true "Prev"
// @Param next query int true "Next"
// @Produce json
// @Success 200 {object} string "Response"
// @Failure 400 {object} string "Error"
// @Failure 500 {object} string "Error"
// @Router /articles/{id}/comments [get]
func Index(c *gin.Context) {
	var limit int
	articleId, _ := primitive.ObjectIDFromHex(c.Param("id"))
	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 0
	}

	prev := c.Query("prev")
	next := c.Query("next")

	res, err := GetComments(articleId, limit, prev, next)
	if err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "comments", "data": res})
}
