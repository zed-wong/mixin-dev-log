// OAuth handler

package auth

import(
	"fmt"

	"gopkg.in/resty.v1"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
)

const (
	MixinOauthURL = "https://api.mixin.one/oauth/token"
)

func NewAuthWorker(clientID, appSecret string){
	r := gin.Default()
	r.GET("/oauth", oauthHandler(clientID, appSecret))
	r.Run()
}

func oauthHandler(clientID, appSecret string) gin.HandlerFunc{
        fn := func(c *gin.Context){
                code := c.Query("code")
                if len(code) == 64{
                        client := resty.New()
                        body := fmt.Sprintf(`{"client_id":"%s","code":"%s","client_secret":"%s"}`, clientID, code, appSecret)

                        resp, err := client.R().
				SetHeader("Content-Type", "application/json").
				SetBody(body).
				Post(MixinOauthURL)
			if err != nil{
                                fmt.Println(err)
                        }
                        accessToken := gjson.Get(resp.String(), `data.access_token`).String()
			c.IndentedJSON(200, gin.H{
				"token": accessToken,
			})
		}
	}
	return fn
}
