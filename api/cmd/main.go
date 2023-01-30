package main

import (
	"github.com/Johannes-Krabbe/notes-app/api/pkg/common/db"
	"github.com/Johannes-Krabbe/notes-app/api/pkg/notes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)
    clientUrl := viper.Get("CLIENT_URL").(string)

    router := gin.Default()
    h := db.Init(dbUrl)


    corsConfig := cors.DefaultConfig()
    corsConfig.AllowOrigins = []string{clientUrl}
    // config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
    corsConfig.AllowCredentials = true
  
    router.Use(cors.New(corsConfig))


    notes.RegisterRoutes(router, h)
    // register more routes here

    router.Run(port)
}