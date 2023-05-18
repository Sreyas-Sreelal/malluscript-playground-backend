package requests

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func GetVersion(c *gin.Context) {
	subProcess := exec.Command("./malluscript", "--version")
	data, err := subProcess.CombinedOutput()

	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"version": string(data),
	})

}
