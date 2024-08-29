package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/s-j-rollins-technologies/hyper-v-monitoring/hyperv"
)

func StartWebServer() (err error) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/hvm/replicas", func(c *gin.Context) {
		replicaStats := hyperv.GetReplicaStats()
		c.JSON(200, replicaStats)
	})

	if err := r.Run(":20501"); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
	return err
}
