package main

import (
  "github.com/gin-gonic/gin"
  "main/controller"
  "main/db"
)

func init() {
  db.Init();
}

func main() {
	r := gin.Default();

  itemController := controller.ItemController{}

  itemRoutes := r.Group("/item")
  {
    itemRoutes.POST("/", itemController.Create);
    itemRoutes.PATCH("/:id", itemController.Edit);
    itemRoutes.DELETE("/:id", itemController.Delete);
    itemRoutes.PATCH("/:id/undelete", itemController.UnDelete);
    itemRoutes.GET("/", itemController.GetAll);
  }
  
  r.Run();
}
