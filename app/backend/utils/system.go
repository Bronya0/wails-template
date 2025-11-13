package utils

import (
	"app/backend/common"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net"
	"os"
	"os/exec"
	"strings"
	"time"
)

// GetMachineID 返回设备的唯一标识
func GetMachineID() string {
	var components []string

	// 3. 获取操作系统特定的硬件信息
	macAddr, err := getFirstMacAddr()
	if err == nil {
		components = append(components, macAddr)
	}

	// 如果没有任何组件可用，使用默认值
	if len(components) == 0 {
		components = append(components, "default-machine-id")
	}

	components = append(components, common.SK)

	// 将所有组件组合并生成哈希
	h := sha256.New()
	for _, c := range components {
		h.Write([]byte(c))
	}
	return hex.EncodeToString(h.Sum(nil))
}

// GetSysHash 生成基于年月和固定字符串的 SHA-256 哈希
func GetSysHash() string {
	yearMonth := time.Now().Format("200601")
	combined := yearMonth + common.SK
	h := sha256.New()
	h.Write([]byte(combined))
	return hex.EncodeToString(h.Sum(nil))
}

func getFirstMacAddr() (string, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range interfaces {
		if len(iface.HardwareAddr) > 0 {
			return iface.HardwareAddr.String(), nil
		}
	}
	return "", fmt.Errorf("no valid MAC address found")
}

// Windows: 获取硬盘序列号
func getWindowsDiskSerial() string {
	cmd := exec.Command("wmic", "diskdrive", "get", "SerialNumber")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed != "" && trimmed != "SerialNumber" {
			return trimmed
		}
	}
	return ""
}

// Linux: 获取 /etc/machine-id 或主板信息
func getLinuxMachineID() string {
	// 优先使用 /etc/machine-id
	data, err := os.ReadFile("/etc/machine-id")
	if err == nil {
		return strings.TrimSpace(string(data))
	}

	// 备选：使用 dmidecode 获取主板序列号
	cmd := exec.Command("dmidecode", "-s", "baseboard-serial-number")
	output, err := cmd.Output()
	if err == nil {
		return strings.TrimSpace(string(output))
	}
	return ""
}

// macOS: 获取序列号
func getMacOSSerial() string {
	cmd := exec.Command("system_profiler", "SPHardwareDataType")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.Contains(line, "Serial Number") {
			parts := strings.Split(line, ":")
			if len(parts) > 1 {
				return strings.TrimSpace(parts[1])
			}
		}
	}
	return ""
}
