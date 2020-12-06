package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github-stars/pkg/model/github"

	"github.com/gin-gonic/gin"
)

const (
	githubAuthURL   = "https://github.com/login/oauth/access_token"
	githubPrefixURL = "https://api.github.com"
	jsonContentType = "application/json"
)

// Auth authentiacation
func Auth(c *gin.Context) {

	var (
		err          error
		param        []byte
		resp         *http.Response
		body         []byte
		authRespData *github.AuthResp
		code         string
	)

	code = c.DefaultQuery("code", "")
	param, err = json.Marshal(
		github.Auth{
			Code:         code,
			ClientID:     "",
			ClientSecret: "",
		},
	)
	if err != nil {
		log.Printf("unmarshal error %v", err)
		c.JSON(400, gin.H{
			"err_message": "unmarsharl error",
		})
		return
	}
	resp, err = http.Post(githubAuthURL, jsonContentType, strings.NewReader(string(param)))
	if err != nil {
		log.Printf("post github error %v", err)
		c.JSON(400, gin.H{
			"err_message": "request gitub error",
		})
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read github resp error %v", err)
		c.JSON(400, gin.H{
			"err_message": "request gitub error",
		})
		return
	}

	authRespData = &github.AuthResp{}

	err = json.Unmarshal(body, authRespData)
	if err != nil {
		log.Printf("unmarshal error %v", err)
		c.JSON(400, gin.H{
			"err_message": "unmarsharl error",
		})
		return
	}

	c.JSON(200, gin.H{
		"access_token": fmt.Sprintf("%s", authRespData.AccessToken),
		"token_type":   fmt.Sprintf("%s", authRespData.TokenType),
	})
}
