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

`http://\<host\>:20501/hvm/replicas`

Additionally, the executable features support for Windows Service Management, allowing it to be installed as a service with `sc.exe`.



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
