package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
)


func render(c *gin.Context, data gin.H, templateName string) {
  loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)


  switch c.Request.Header.Get("Accept") {
  case "application/json":
    // response with json
    c.JSON(http.StatusOK, data["payload"])
  case "application/xml":
    // Respond with XML
    c.XML(http.StatusOK, data["payload"])
  default:
    // Respond with HTML
    c.HTML(http.StatusOK, templateName, data)
  }
}

var router *gin.Engine

func main() {
  router = gin.Default(); // set the router to be used in the rest of the app

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("templates/*"); // load all files contained in templates folder

  // initializing the routes
  initializeRoutes()

  // start the application
  router.Run()

}
