package manager

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Manager struct {}

func (Manager) PlatformMessage(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
