package system

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/docker"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

type Info struct {
	CPU              []cpu.InfoStat
	DiskUsage        disk.UsageStat
	DiskPartitions   []disk.PartitionStat
	DiskIOCounters   map[string]disk.IOCountersStat
	Docker           []docker.CgroupDockerStat
	Host             host.InfoStat
	HostUsers        []host.UserStat
	HostTemperature  []host.TemperatureStat
	LoadAvg          load.AvgStat
	LoadMisc         load.MiscStat
	VMemory          mem.VirtualMemoryStat
	SMemory          mem.SwapMemoryStat
	NetIOCounters    []net.IOCountersStat
	NetConnections   []net.ConnectionStat
	NetInterfaces    []net.InterfaceStat
	NetProtoCounters []net.ProtoCountersStat
}

func AllInfo() Info {
	cpuStat, _ := CPU()
	diskUsage, _ := DiskUsage()
	diskPartitions, _ := DiskPartitions()
	diskIOCounters, _ := DiskIOCounters()
	dockerStats, _ := Docker()
	hostStat, _ := Host()
	hostUsers, _ := HostUsers()
	hostTemperature, _ := HostTemperature()
	loadAvg := LoadAvg()
	loadMisc := LoadMisc()
	vMemory, _ := VMemory()
	sMemory, _ := SMemory()
	netIOCounters, _ := NetIOCounters()
	netConnections, _ := NetConnections()
	netInterfaces, _ := NetInterfaces()
	netProtoCounters, _ := NetProtoCounters()
	return Info{
		CPU:              cpuStat,
		DiskUsage:        diskUsage,
		DiskPartitions:   diskPartitions,
		DiskIOCounters:   diskIOCounters,
		Docker:           dockerStats,
		Host:             hostStat,
		HostUsers:        hostUsers,
		HostTemperature:  hostTemperature,
		LoadAvg:          loadAvg,
		LoadMisc:         loadMisc,
		VMemory:          vMemory,
		SMemory:          sMemory,
		NetIOCounters:    netIOCounters,
		NetConnections:   netConnections,
		NetInterfaces:    netInterfaces,
		NetProtoCounters: netProtoCounters,
	}
}

func CPU() ([]cpu.InfoStat, error) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	return cpuInfo, nil
}

func DiskUsage() (disk.UsageStat, error) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return disk.UsageStat{}, err
	}
	return *diskUsage, nil
}

func DiskPartitions() ([]disk.PartitionStat, error) {
	diskPartitions, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	return diskPartitions, nil
}

func DiskIOCounters() (map[string]disk.IOCountersStat, error) {
	diskIOCounters, err := disk.IOCounters()
	if err != nil {
		return nil, err
	}
	return diskIOCounters, nil
}

func Docker() ([]docker.CgroupDockerStat, error) {
	dockerStat, err := docker.GetDockerStat()
	if err != nil {
		return nil, err
	}
	return dockerStat, nil
}

func Host() (host.InfoStat, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return host.InfoStat{}, err
	}
	return *hostInfo, nil
}

func HostUsers() ([]host.UserStat, error) {
	hostUsers, err := host.Users()
	if err != nil {
		return nil, err
	}
	return hostUsers, nil
}

func HostTemperature() ([]host.TemperatureStat, error) {
	hostTemperature, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}
	return hostTemperature, nil
}

func LoadAvg() load.AvgStat {
	loadAvg, err := load.Avg()
	if err != nil {
		return load.AvgStat{}
	}
	return *loadAvg
}

func LoadMisc() load.MiscStat {
	loadMisc, err := load.Misc()
	if err != nil {
		return load.MiscStat{}
	}
	return *loadMisc
}

func VMemory() (mem.VirtualMemoryStat, error) {
	vMemory, err := mem.VirtualMemory()
	if err != nil {
		return mem.VirtualMemoryStat{}, err
	}
	return *vMemory, nil
}

func SMemory() (mem.SwapMemoryStat, error) {
	swapMemory, err := mem.SwapMemory()
	if err != nil {
		return mem.SwapMemoryStat{}, err
	}
	return *swapMemory, nil
}

func NetIOCounters() ([]net.IOCountersStat, error) {
	netIOCounters, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	return netIOCounters, nil
}

func NetConnections() ([]net.ConnectionStat, error) {
	netConnections, err := net.Connections("all")
	if err != nil {
		return nil, err
	}
	return netConnections, nil
}

func NetInterfaces() ([]net.InterfaceStat, error) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	return netInterfaces, nil
}

func NetProtoCounters() ([]net.ProtoCountersStat, error) {
	netProtoCounters, err := net.ProtoCounters([]string{"tcp", "udp"})
	if err != nil {
		return nil, err
	}
	return netProtoCounters, nil
}
