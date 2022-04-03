package api

import (
	"net/http"
	"time"

	mw "viewee-service/middleware"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v4"
)

type LoginInfo struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	// var loginReq model.LoginReq
	user := LoginInfo{
		Id:       "tester",
		Username: "tester",
		Password: "password",
	}
	// if c.BindJSON(&loginReq) == nil {
	//     isPass, user, err := model.LoginCheck(loginReq)
	//     if isPass {
	token, err := generateToken(c, user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "Login success",
		"data":   data,
	})

	//     } else {
	//         c.JSON(http.StatusOK, gin.H{
	//             "status": -1,
	//             "msg":    "User not authenticated, " + err.Error(),
	//         })
	//     }
	// } else {
	//     c.JSON(http.StatusOK, gin.H{
	//         "status": -1,
	//         "msg":    "json parse error",
	//     })
	// }
}

// LoginResult
type LoginResult struct {
	Token string `json:"token"`
	User  LoginInfo
}

func generateToken(c *gin.Context, user LoginInfo) (string, error) {
	j := &mw.JWT{
		SigningKey: []byte("xxx"),
	}

	claims := mw.JWTClaims{
		user.Id,
		user.Username,
		user.Password,
		jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),
			ExpiresAt: int64(time.Now().Unix() + 3600),
			Issuer:    "newtrekWang",
		},
	}

	return j.CreateToken(claims)

	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": -1,
	// 		"msg":    err.Error(),
	// 	})
	// 	return
	// }

	// log.Println(token)

	// data := LoginResult{
	// 	User:  user,
	// 	Token: token,
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"status": 0,
	// 	"msg":    "登录成功！",
	// 	"data":   data,
	// })
	// return
}

// func GetDataByTime(c *gin.Context) {
// 	claims := c.MustGet("claims").(*mw.JWTClaims)
// 	if claims != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": 0,
// 			"msg":    "token有效",
// 			"data":   claims,
// 		})
// 	}
// }
