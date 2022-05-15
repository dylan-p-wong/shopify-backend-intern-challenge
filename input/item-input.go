package input

type ItemInput struct {
  Count *int `json:"count" binding:"required,gte=0,lte=2000000"`
  Title string `json:"title" binding:"required,min=3,max=255"`
  Description string `json:"description"`
}
