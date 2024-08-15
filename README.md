# Hyper V Monitoring

The *Hyper V Monitoring* package is a simple, lightweight executable to be run on Microsoft Windows machines with the Hyper-V role. It returns a JSON array of all Hyper-V Replicas and the below properties for each (if applicable): 

 - ID
 - Name
 - InstanceID
 - AllocatedGPU
 - Shielded
 - State
 - OtherEnabledState
 - GuestOperatingSystem
 - HealthState
 - Heartbeat
 - MemoryUsage
 - MemoryAvailable
 - AvailableMemoryBuffer
 - OperationalStatus
 - NumberOfProcessors
 - ProcessorLoad
 - UpTime
 - ReplicationState
 - ReplicationHealth
 - ReplicationMode
 - ApplicationHealth
 - IntegrationServicesVersionState
 - HostComputerSystemName

A GET request to the below endpoint returns this data.

`http://<host>:20501/hvm/replicas`

Additionally, the executable features support for Windows Service Management, allowing it to be installed as a service with `sc.exe`.


### Example Output
```
[
  {        
    "id": "98873FAE-33FF-508B-429B-90102EFB1CE7",
    "name": "Arch Linux VM",
    "state": "running",
    "instance_id": "Microsoft:98873FAE-33FF-508B-429B-90102EFB1CE7",
    "allocated_gpu": "",
    "shielded": true,
    "other_enabled_state": "",
    "guest_operating_system": "Arch Linux",
    "health_state": 5,
    "heartbeat": 5,
    "memory_usage": 1420,
    "memory_available": 17,
    "available_memory_buffer": 107,
    "number_of_processors": 1,
    "operational_status": [2],
    "processor_load": 1,
    "uptime": 1234578,
    "replication_state": 3,
    "replication_health": "Critical",
    "replication_mode": 2,
    "application_health": 0,
    "integration_services_version_state": 2,
    "host_computer_system_name": "SJROLLINS-PC"
  }
]
```


### Future Features

- `config.json` file for configuration of things like
  - HTTP/S Port
  - TLS Certificates for HTTPS
  - X-API-KEY Authentication
  - Whitelisted IPs
  - Numerical vs. Friendly Values
- API Authentication
- Non-replica VM monitoring data
- Hyper-V Host monitoring data
- HTTPS Support
- Windows Installer
