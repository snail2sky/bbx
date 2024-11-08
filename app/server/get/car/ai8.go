package car

import (
	"github.com/gin-gonic/gin"
	"github.com/snail2sky/bbx/app/get/car"
	"net/http"
)

type AI8Config struct {
	Num string `form:"num" binding:"required"`
}

func Ai8Num(c *gin.Context) {
	var ai8Config AI8Config
	if err := c.Bind(&ai8Config); err == nil {
		password, err := car.GetArrizoPassword(ai8Config.Num)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
		} else {
			c.String(http.StatusOK, password)
		}
	} else {
		c.String(http.StatusBadRequest, err.Error())
	}
}
