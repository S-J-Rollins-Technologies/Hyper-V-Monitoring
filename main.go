package main

import (
	"fmt"
	"log"
	"os"

	"github.com/s-j-rollins-technologies/hyper-v-monitoring/hyperv"

	"github.com/gin-gonic/gin"
	"golang.org/x/sys/windows/svc"
)

type hvmonSvc struct{}

func (m *hvmonSvc) Execute(args []string, req <-chan svc.ChangeRequest, status chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
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

	r.GET("/hvm/replicas", func(c *gin.Context) {
		replicaStats := getReplicaStats()
		c.JSON(200, replicaStats)
	})

	if err := r.Run(":20501"); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}
}

type ReplicaStats struct {
	ID                              string   `json:"id"`
	Name                            string   `json:"name"`
	State                           string   `json:"state"`
	InstanceID                      string   `json:"instance_id"`
	AllocatedGPU                    string   `json:"allocated_gpu"`
	Shielded                        bool     `json:"shielded"`
	OtherEnabledState               string   `json:"other_enabled_state"`
	GuestOperatingSystem            string   `json:"guest_operating_system"`
	HealthState                     uint16   `json:"health_state"`
	Heartbeat                       uint16   `json:"heartbeat"`
	MemoryUsage                     uint64   `json:"memory_usage"`
	MemoryAvailable                 int      `json:"memory_available"`
	AvailableMemoryBuffer           int      `json:"available_memory_buffer"`
	NumberOfProcessors              uint16   `json:"number_of_processors"`
	OperationalStatus               []uint16 `json:"operational_status"`
	ProcessorLoad                   uint16   `json:"processor_load"`
	UpTime                          uint64   `json:"uptime"`
	ReplicationState                uint16   `json:"replication_state"`
	ReplicationHealth               string   `json:"replication_health"`
	ReplicationMode                 uint16   `json:"replication_mode"`
	ApplicationHealth               uint16   `json:"application_health"`
	IntegrationServicesVersionState uint16   `json:"integration_services_version_state"`
	HostComputerSystemName          string   `json:"host_computer_system_name"`
}

func getReplicaStats() []ReplicaStats {
	vms, err := hyperv.GetVMList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	fmt.Println(vms)
	var AllReplicaStats []ReplicaStats

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

		replicaStats := ReplicaStats{
			ID:                              vm.ID,
			Name:                            vm.Name,
			State:                           vm.State.String(),
			InstanceID:                      vm.InstanceID,
			AllocatedGPU:                    vm.AllocatedGPU,
			Shielded:                        vm.Shielded,
			OtherEnabledState:               vm.OtherEnabledState,
			GuestOperatingSystem:            vm.GuestOperatingSystem,
			HealthState:                     vm.HealthState,
			Heartbeat:                       vm.HealthState,
			MemoryUsage:                     vm.MemoryUsage,
			MemoryAvailable:                 vm.MemoryAvailable,
			AvailableMemoryBuffer:           vm.AvailableMemoryBuffer,
			NumberOfProcessors:              vm.NumberOfProcessors,
			OperationalStatus:               vm.OperationalStatus,
			ProcessorLoad:                   vm.ProcessorLoad,
			UpTime:                          vm.UpTime,
			ReplicationState:                vm.ReplicationState,
			ReplicationHealth:               health,
			ReplicationMode:                 vm.ReplicationMode,
			ApplicationHealth:               vm.ApplicationHealth,
			IntegrationServicesVersionState: vm.IntegrationServicesVersionState,
			HostComputerSystemName:          vm.HostComputerSystemName,
		}

		AllReplicaStats = append(AllReplicaStats, replicaStats)
	}

	return AllReplicaStats
}

func main() {
	isService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("Failed to determine if running as service: %v", err)
	}

	if isService {
		if err := svc.Run("HyperVReplicaHealthAPI", &hvmonSvc{}); err != nil {
			log.Fatalf("Service failed: %v", err)
		}
	} else {
		runMain()
	}
}
