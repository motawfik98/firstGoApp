package main

import "testing"

func TestGetAllArticles(t *Testing.T) {
  alist := getAllArticles()

  if len(alist) != len(articleList) {
    t.Fail()
  }

  for i, v := range alist {
    if v.Content != articleList[i].Content ||
       v.ID != articleList[i].ID ||
       v.Title != articleList[i].Title {
         t.Fail()
         break
       }
  }
}
