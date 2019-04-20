package main

import (
  "net/http"
  "strconv"
  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  articles := getAllArticles()

  // call the HTML method of the Context to render a template
  render(c, gin.H{
    "title": "Home Page",
    "payload": articles,
    },
    "index.html")
}

func getArticle(c *gin.Context) {
  // check if the article ID is valid
  if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
    // check if the article exists
    if article, err := getArticleByID(articleID); err == nil {
      // Call the HTML method of the Context to render a template
      c.HTML (
        // set the status code to OK
        http.StatusOK,
        // use the article.html
        "article.html",
        // passes the data that the pages use
        gin.H {
          "title": article.Title,
          "payload": article,
        },
      )
    } else {
      // If the article is not found, abort with an error
      c.AbortWithError(http.StatusNotFound, err)
    }
  } else {
    // If an invalid article ID is specified in the URL, abort with an error
    c.AbortWithStatus(http.StatusNotFound)
  }
}


func showArticleCreationPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Create New Article"}, "create-article.html")
}


func createArticle(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")

	if a, err := createNewArticle(title, content); err == nil {
		render(c, gin.H{
			"title":   "Submission Successful",
			"payload": a}, "submission-successful.html")
	} else {
		c.AbortWithStatus(http.StatusBadRequest)
	}
}
