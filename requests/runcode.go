package requests

import (
	"fmt"
	"io"
	"malluscript/types"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func generateFileName() string {
	uuid := uuid.New()
	return uuid.String()+".ms"
}

func RunCode(c *gin.Context) {
	var json types.CodeRequest
	
	if err := c.ShouldBind(&json); err != nil {
		fmt.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "bad request! " + err.Error(),
		})
		return
	}
	
	filename := generateFileName()

	f, err := os.Create("scripts/"+filename)
	if err != nil {
		os.Mkdir("scripts",os.ModePerm)
		f, err = os.Create("scripts/"+filename)
	}
	
	f.Write([]byte(json.Code))

	subProcess := exec.Command("./malluscript", "scripts/"+filename)
	
	stdin, err := subProcess.StdinPipe()
	if err != nil {
		f.Close()
		os.Remove("scripts/"+filename)
		panic(err)
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, json.Input)
	}()

	data, _ := subProcess.CombinedOutput()
	
	f.Close()
	err = os.Remove("scripts/"+filename)
	if err!=nil {
		panic(err)
	}
	
	formatted := string(data)
	strings.Replace(formatted,"\n","\\n",-1)
	
	c.JSON(http.StatusOK, gin.H{
		"output": string(formatted),
	})

}
