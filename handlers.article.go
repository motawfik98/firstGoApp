package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
  articles := getAllArticles()

  // call the HTML method of the Context to render a template
  c.HTML (
    // set the status code to 200
    http.StatusOK,
    // use the index..html template
    "index.html",
    // pass the data that the page uses
    gin.H {
      "title": "Home Page",
      "payload": articles,
    },
  )
}
