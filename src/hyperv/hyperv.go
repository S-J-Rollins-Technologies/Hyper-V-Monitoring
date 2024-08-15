package hyperv

import (
	"errors"
	"fmt"
	"time"

	"github.com/bi-zone/wmi"
)

const hyperVNamespace = `root\virtualization\v2`

var (
	ErrQuery = errors.New("failed to query WMI interface")
)

type VMList []VM

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

func (vms *VMList) ToWMIQuery() string {
	return wmi.CreateQueryFrom(vms, "Msvm_SummaryInformation", "")
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

func GetVMList() (*VMList, error) {
	var vms VMList

	q := vms.ToWMIQuery()
	if err := wmi.QueryNamespace(q, &vms, hyperVNamespace); err != nil {
		return nil, fmt.Errorf("%s: %w\n\t- query:\t%s\n\t- namespace:\t%s", ErrQuery, err, q, hyperVNamespace)
	}

	return &vms, nil
}
