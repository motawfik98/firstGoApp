package main

import (
  "github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
  router := gin.Default(); // set the router to be used in the rest of the app

  // Process the templates at the start so that they don't have to be loaded
  // from the disk again. This makes serving HTML pages very fast.
  router.LoadHTMLGlob("templates/*"); // load all files contained in templates folder

  // initializing the routes
  initializeRoutes()

  // start the application
  router.Run()

}
