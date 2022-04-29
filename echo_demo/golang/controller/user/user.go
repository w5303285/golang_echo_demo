package user

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
    // 方式1：获取url上的path参数，url模式里面定义了参数:id
    id := c.Param("id")
    
    return c.String(http.StatusOK, "conterollr获取参数例子"+id)
}