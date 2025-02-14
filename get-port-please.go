package getportplease

import (
	"fmt"
	"net"
)

func GetFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", ":0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}

	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}
func IsPortAvailable(port int) (bool, error) {
	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return false, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return false, nil
	}
	defer l.Close()
	return true, nil
}

func GetFreePortInRange(minPort, maxPort int) (int, error) {
	if minPort < 1 || maxPort > 65535 || minPort > maxPort {
		return 0, fmt.Errorf("invalid port range")
	}

	for port := minPort; port <= maxPort; port++ {
		available, err := IsPortAvailable(port)
		if err != nil {
			return 0, err
		}
		if available {
			return port, nil
		}
	}
	return 0, fmt.Errorf("no available port found in the range")
}
