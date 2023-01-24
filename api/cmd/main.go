package main

import (
	"github.com/Johannes-Krabbe/notes-app/api/pkg/common/db"
	"github.com/Johannes-Krabbe/notes-app/api/pkg/notes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
    viper.SetConfigFile("./pkg/common/envs/.env")
    viper.ReadInConfig()

    port := viper.Get("PORT").(string)
    dbUrl := viper.Get("DB_URL").(string)

    r := gin.Default()
    h := db.Init(dbUrl)

    notes.RegisterRoutes(r, h)
    // register more routes here

    r.Run(port)
}