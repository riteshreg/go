package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

type SystemInfo struct {
	Hostname    string `json:"hostname"`
	OS          string `json:"os"`
	Platform    string `json:"platform"`
	PlatformVer string `json:"platform_version"`
	CPUModel    string `json:"cpu_model"`
	CPUCount    int    `json:"cpu_count"`
	TotalMemory uint64 `json:"total_memory"`
	FreeMemory  uint64 `json:"free_memory"`
	DiskUsage   uint64 `json:"disk_usage"`
	DiskFree    uint64 `json:"disk_free"`
}

func getSystemInfo() (*SystemInfo, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	hostInfo, err := host.Info()
	if err != nil {
		return nil, err
	}

	diskUsage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	sysInfo := SystemInfo{
		FreeMemory:  vmStat.Free,
		TotalMemory: vmStat.Total,
		CPUModel:    strings.TrimSpace(cpuInfo[0].ModelName),
		CPUCount:    len(cpuInfo),
		Hostname:    hostInfo.Hostname,
		OS:          hostInfo.OS,
		Platform:    hostInfo.Platform,
		PlatformVer: hostInfo.PlatformVersion,
		DiskUsage:   diskUsage.Used,
		DiskFree:    diskUsage.Free,
	}

	return &sysInfo, nil

}

func systemInfoHandler(w http.ResponseWriter, r *http.Request) {
	systemInfo, err := getSystemInfo()
	if err != nil {
		http.Error(w, "Failed to get Sys Info", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(systemInfo)

}

func main() {

	http.HandleFunc("/", systemInfoHandler)
	fmt.Println("Server running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
