package models

import (
	"time"

	"github.com/bi-zone/wmi"
)

type VMList []VM

func (vms *VMList) ToWMIQuery() string {
	return wmi.CreateQueryFrom(vms, "Msvm_SummaryInformation", "")
}

type VM struct {
	ID                              string `wmi:"Name"`
	Name                            string `wmi:"ElementName"`
	InstanceID                      string
	AllocatedGPU                    string
	Shielded                        bool
	CreationTime                    time.Time
	State                           State `wmi:"EnabledState"`
	OtherEnabledState               string
	GuestOperatingSystem            string
	HealthState                     uint16
	Heartbeat                       uint16
	MemoryUsage                     uint64
	MemoryAvailable                 int
	AvailableMemoryBuffer           int
	SwapFilesInUse                  bool
	Notes                           string
	Version                         string
	NumberOfProcessors              uint16
	OperationalStatus               []uint16
	ProcessorLoad                   uint16
	ProcessorLoadHistory            []uint16
	StatusDescriptions              []string
	ThumbnailImage                  []uint8
	ThumbnailImageHeight            uint16
	ThumbnailImageWidth             uint16
	UpTime                          uint64
	ReplicationState                uint16
	ReplicationStateEx              []uint16
	ReplicationHealth               uint16
	ReplicationHealthEx             []uint16
	ReplicationMode                 uint16
	ApplicationHealth               uint16
	IntegrationServicesVersionState uint16
	MemorySpansPhysicalNumaNodes    bool
	ReplicationProviderId           []string
	EnhancedSessionModeState        uint16
	VirtualSwitchNames              []string
	VirtualSystemSubType            string
	HostComputerSystemName          string
}

type State uint16

const (
	StateUnknown State = iota
	StateOther
	StateRunning
	StateOff
	StateShuttingDown
	StateNotApplicable
	StateEnabledButOffline
	StateInTest
	StateDeferred
	StateQuiesce
	StateStarting
)

func (s State) String() string {
	switch s {
	case StateUnknown:
		return "unknown"
	case StateOther:
		return "other"
	case StateRunning:
		return "running"
	case StateOff:
		return "off"
	case StateShuttingDown:
		return "shutting down"
	case StateNotApplicable:
		return "not applicable"
	case StateEnabledButOffline:
		return "enabled but offline"
	case StateInTest:
		return "in test"
	case StateDeferred:
		return "deferred"
	case StateQuiesce:
		return "quiesce"
	case StateStarting:
		return "starting"
	default:
		if 11 <= s && 32767 <= s {
			return "DMTF reserved"
		}

		if 32768 <= s && 65535 <= s {
			return "vendor reserved"
		}

		return ""
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
