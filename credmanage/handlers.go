package credmanage

import (
	"feather/environment"
	redisclient "feather/redis-client"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var Clt *redis.Client
var signkey string

func init(){
	environment.ViperExport()
	connstring := viper.GetString("redisServer.connectionstring")
	redispass := viper.GetString("redisServer.password")
	redisdb := viper.GetInt("redisServer.database")
	Clt = redisclient.CreateClient(connstring,redispass,redisdb)
	signkey = viper.GetString("jwt.signkey")
}

type Credential struct{
	Username string	`json:"username"`
	Password string	`json:"password"`
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}

func Authenticate(c echo.Context)error{
	cred := &Credential{}
	if err := c.Bind(cred); err != nil {
		return err
	}

	hash,err := redisclient.RetrieveHashSalt(Clt,cred.Username)

	if err != nil {
		return err
	}

	if !ValidateHash(hash,cred.Password){
		return c.String(http.StatusForbidden,"Wrong Password")
	}
	
	t,err := CreateSignedToken(cred.Username,false,time.Now().Add(time.Minute*3),signkey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}

func SignUp(c echo.Context)error{
	cred := &Credential{}
	if err := c.Bind(cred); err != nil {
		return err
	}

	err := redisclient.StoreHashSalt(Clt,cred.Username,HashGen(cred.Password))
	if err!=nil{
		return err
	}
	t,err := CreateSignedToken(cred.Username,false,time.Now().Add(time.Minute*3),signkey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": t,
	})
}
