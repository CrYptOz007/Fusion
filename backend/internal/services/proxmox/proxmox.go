package proxmox

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CrYptOz007/Fusion/internal/models/service"
	"github.com/luthermonson/go-proxmox"
)

func ListNodes(service *service.Service) (proxmox.NodeStatuses, error) {

	tokenID := service.Username
	secret := service.ApiKey

	httpClient := http.Client{
		Transport: &http.Transport{},
	}

	client := proxmox.NewClient("https://"+service.Hostname+":"+fmt.Sprintf("%d", service.Port)+"/api2/json",
		proxmox.WithHTTPClient(&httpClient),
		proxmox.WithAPIToken(tokenID, secret))

	nodes, err := client.Nodes(context.Background())

	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func ListVMs(service *service.Service, node string) (proxmox.VirtualMachines, error) {

	tokenID := service.Username
	secret := service.ApiKey

	httpClient := http.Client{
		Transport: &http.Transport{},
	}

	client := proxmox.NewClient("https://"+service.Hostname+":"+fmt.Sprintf("%d", service.Port)+"/api2/json",
		proxmox.WithHTTPClient(&httpClient),
		proxmox.WithAPIToken(tokenID, secret))

	nodeInfo, err := client.Node(context.Background(), node)

	if err != nil {
		return nil, err
	}

	vms, err := nodeInfo.VirtualMachines(context.Background())

	if err != nil {
		return nil, err
	}

	return vms, nil
}

func ListContainers(service *service.Service, node string) (proxmox.Containers, error) {
	tokenID := service.Username
	secret := service.ApiKey

	httpClient := http.Client{
		Transport: &http.Transport{},
	}

	client := proxmox.NewClient("https://"+service.Hostname+":"+fmt.Sprintf("%d", service.Port)+"/api2/json",
		proxmox.WithHTTPClient(&httpClient),
		proxmox.WithAPIToken(tokenID, secret))

	nodeInfo, err := client.Node(context.Background(), node)

	if err != nil {
		return nil, err
	}

	containers, err := nodeInfo.Containers(context.Background())

	if err != nil {
		return nil, err
	}

	return containers, nil

}
