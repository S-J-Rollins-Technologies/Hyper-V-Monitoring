package main

import (
	"fmt"
	"log"
	"os"

	"hyperv-monitoring/hyperv"

	"github.com/gin-gonic/gin"
	"golang.org/x/sys/windows/svc"
)

type myService struct{}

func (m *myService) Execute(args []string, req <-chan svc.ChangeRequest, status chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	status <- svc.Status{State: svc.StartPending}

	go runMain()

	status <- svc.Status{State: svc.Running, Accepts: svc.AcceptStop | svc.AcceptShutdown}

loop:
	for {
		select {
		case c := <-req:
			switch c.Cmd {
			case svc.Interrogate:
				status <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				status <- svc.Status{State: svc.StopPending}
				status <- svc.Status{State: svc.Stopped}
				break loop
			default:
				continue
			}
		}
	}

	return
}

func runMain() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/hv_rep_stat", func(c *gin.Context) {
		replicaHealthStatuses := getReplicaHealth()
		c.JSON(200, replicaHealthStatuses)
	})

	if err := r.Run(":20501"); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
}

type ReplicaHealth struct {
	Name  string `json:"name"`
	State string `json:"replica_health"`
}

func getReplicaHealth() []ReplicaHealth {
	vms, err := hyperv.GetVMList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	var AllReplicaHealth []ReplicaHealth

	for _, vm := range *vms {
		var health string

		switch vm.ReplicationHealth {
		case 0:
			health = "Not Applicable"
		case 1:
			health = "Normal"
		case 2:
			health = "Warning"
		case 3:
			health = "Critical"
		default:
			health = "Unknown"
		}

		replicaHealth := ReplicaHealth{
			Name:  vm.Name,
			State: health,
		}

		AllReplicaHealth = append(AllReplicaHealth, replicaHealth)
	}

	return AllReplicaHealth
}

func main() {
	isService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("Failed to determine if running as service: %v", err)
	}

	if isService {
		if err := svc.Run("HyperVReplicaHealthAPI", &myService{}); err != nil {
			log.Fatalf("Service failed: %v", err)
		}
	} else {
		runMain()
	}
}
