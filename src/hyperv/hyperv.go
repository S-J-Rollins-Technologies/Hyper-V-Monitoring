package hyperv

import (
	"errors"
	"fmt"
	"os"

	"github.com/bi-zone/wmi"
	"github.com/s-j-rollins-technologies/hyper-v-monitoring/models"
)

const hyperVNamespace = `root\virtualization\v2`

var (
	ErrQuery = errors.New("failed to query WMI interface")
)

func GetVMList() (*models.VMList, error) {
	var vms models.VMList

	q := vms.ToWMIQuery()
	if err := wmi.QueryNamespace(q, &vms, hyperVNamespace); err != nil {
		return nil, fmt.Errorf("%s: %w\n\t- query:\t%s\n\t- namespace:\t%s", ErrQuery, err, q, hyperVNamespace)
	}

	return &vms, nil
}

func GetReplicaStats() []models.ReplicaStats {
	vms, err := GetVMList()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return nil
	}

	fmt.Println(vms)
	var AllReplicaStats []models.ReplicaStats

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

		replicaStats := models.ReplicaStats{
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
