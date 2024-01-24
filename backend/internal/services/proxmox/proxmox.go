package proxmox

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"

	encryption "github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/service"
	"github.com/luthermonson/go-proxmox"
)

func ListNodes(service *service.Service) (proxmox.NodeStatuses, error) {

	tokenID := service.Username
	secret, err := encryption.Decrypt(service.ApiKey, service.User.Password, service.User.Salt)
	if err != nil {
		return nil, err
	}

	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client := proxmox.NewClient("https://"+service.Hostname+":"+fmt.Sprintf("%d", service.Port)+"/api2/json",
		proxmox.WithHTTPClient(&insecureHTTPClient),
		proxmox.WithAPIToken(tokenID, secret))

	nodes, err := client.Nodes(context.Background())

	if err != nil {
		return nil, err
	}

	return nodes, nil
}

func ListVMs(service *service.Service, node string) (proxmox.VirtualMachines, error) {

	tokenID := service.Username
	secret, err := encryption.Decrypt(service.ApiKey, service.User.Password, service.User.Salt)
	if err != nil {
		return nil, err
	}

	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client := proxmox.NewClient("https://"+service.Hostname+":"+fmt.Sprintf("%d", service.Port)+"/api2/json",
		proxmox.WithHTTPClient(&insecureHTTPClient),
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
	secret, err := encryption.Decrypt(service.ApiKey, service.User.Password, service.User.Salt)
	if err != nil {
		return nil, err
	}

	insecureHTTPClient := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client := proxmox.NewClient("https://"+service.Hostname+":"+fmt.Sprintf("%d", service.Port)+"/api2/json",
		proxmox.WithHTTPClient(&insecureHTTPClient),
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