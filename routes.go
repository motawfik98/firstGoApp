package main

func initializeRoutes() {
  // handle the index route
  router.GET("/", showIndexPage)

  router.GET("/article/view/:article_id", getArticle)
}
