package main

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	// "html/template"
	// "io"
	"log"
	"net/http"
	"time"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/rest/v1/msisdn-lookup/:msisdn", func(c *gin.Context) {
		msisdn := c.Param("msisdn")
		// c.String(http.StatusOK, "Hello %s", msisdn)
		c.JSON(http.StatusOK, gin.H{
			"msisdn": msisdn,
			"status": http.StatusOK,
		})
	})
	router.GET("/soap/v1/msisdn-lookup/:msisdn", func(c *gin.Context) {
		msisdn := c.Param("msisdn")
		// c.String(http.StatusOK, "Hello %s", msisdn)
		c.XML(http.StatusOK, gin.H{
			"msisdn": msisdn,
			"status": http.StatusOK,
		})
	})

	router.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note than you are using the copied context "c_cp", IMPORTANT
			message := "Done! in path " + cCp.Request.URL.Path
			log.Println(message)
			c.String(http.StatusOK, message)
		}()

	})

	router.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(5 * time.Second)

		// since we are NOT using a goroutine, we do not have to copy the context
		message := "Done! in path " + c.Request.URL.Path
		log.Println(message)
		c.String(http.StatusOK, message)
	})

	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run()
	endless.ListenAndServe(":80", router)
	// router.Run(":3000") for a hard coded port
}
