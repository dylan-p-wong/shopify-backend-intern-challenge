package controller

import (
	"errors"
	"main/error"
	"main/input"
	"main/repository"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ItemController struct {
}

func (c *ItemController) Create(context *gin.Context) {
  body:=input.ItemInput{};
  
  if err:=context.BindJSON(&body); err != nil {
    var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]string, len(ve))
            for i, fe := range ve {
                out[i] = error.GetErrorMsg(fe);
            }
            context.AbortWithStatusJSON(400, gin.H{"errors": out});
        }
    return;
  }

  repository := repository.ItemRepository{}
  
  item, err := repository.Create(body);

  if err != nil {
    context.AbortWithStatusJSON(400, gin.H{ "error" : err.Error() });
    return;
  }
  
	context.JSON(200, item);
}

func (c *ItemController) Edit(context *gin.Context) {
  body:=input.ItemInput{};
  
  if err:=context.BindJSON(&body); err != nil {
    var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]string, len(ve))
            for i, fe := range ve {
                out[i] = error.GetErrorMsg(fe);
            }
            context.AbortWithStatusJSON(400, gin.H{"errors": out});
        }
    return;
  }
  
  repository := repository.ItemRepository{};
  
  item, err := repository.Edit(context.Param("id"), body);

  if err != nil {
    context.AbortWithStatusJSON(400, gin.H{"errors" : [1]string{ err.Error()}});
    return;
  }
  
	context.JSON(200, item);
}

func (c *ItemController) GetAll(context *gin.Context) {
  repository := repository.ItemRepository{};

  deleted := context.Query("deleted");

  if deleted != "" {
    items := repository.GetAll(true);
    context.JSON(200, items);
    return;
  } 
  
  items := repository.GetAll(false);
	context.JSON(200, items);
}

func (c *ItemController) Delete(context *gin.Context) {
  body:=input.DeleteInput{};
  
  if err:=context.BindJSON(&body); err != nil {
    var ve validator.ValidationErrors
        if errors.As(err, &ve) {
            out := make([]string, len(ve))
            for i, fe := range ve {
                out[i] = error.GetErrorMsg(fe);
            }
            context.AbortWithStatusJSON(400, gin.H{"errors": out});
        }
    return;
  }
  
  repository := repository.ItemRepository{};
  
  err := repository.Delete(context.Param("id"), body.DeletionComments);

  if err != nil {
    context.AbortWithStatusJSON(400, gin.H{"errors" : [1]string{ err.Error()}});
    return;
  }
  
	context.Status(204);
}

func (c *ItemController) UnDelete(context *gin.Context) {
  repository := repository.ItemRepository{};
  
  item, err := repository.UnDelete(context.Param("id"));

  if err != nil {
    context.AbortWithStatusJSON(400, gin.H{"errors" : [1]string{ err.Error()}});
    return;
  }
  
	context.JSON(200, item);
}
