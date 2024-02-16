/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/minyou08042/kouyi_palm_treasure/internal/app"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/dig"
	"github.com/minyou08042/kouyi_palm_treasure/pkg/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kouyi palm treasure",
	Short: "kouyi palm treasure server",
	//Long: `A longer description that spans multiple lines and likely contains`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		// Run
		if err := run(); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			os.Exit(-1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kouyi_palm_treasure.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run() error {

	g := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowOrigins = []string{"http://172.104.98.231", "http://172.105.193.78"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "UPDATA"}
	corsConfig.AllowHeaders = []string{"Authorization", "Origin"}
	corsConfig.ExposeHeaders = []string{"Authorization", "Content-Type", "Origin",
		"Connection", "Accept-Encoding", "Accept-Language", "Host"}
	corsConfig.MaxAge = 12 * time.Hour
	g.Use(CORSMiddleware())
	//g.Use(cors.New(corsConfig))
	d := dig.BuildContainer()
	var l logger.LogInfoFormat
	dig.Invoke(func(log logger.LogInfoFormat) {
		l = log
	})

	svr := app.NewServer(g, d, l)
	svr.MapRoutes()
	if err := svr.SetupDB(); err != nil {
		return err
	}
	return svr.Start()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
