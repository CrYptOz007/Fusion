package proxmox

import (
	"strconv"

	"github.com/CrYptOz007/Fusion/internal/helpers"
	"github.com/CrYptOz007/Fusion/internal/models/service"
	"github.com/CrYptOz007/Fusion/internal/server/types"
	"github.com/CrYptOz007/Fusion/internal/server/utils"
	"github.com/CrYptOz007/Fusion/internal/services/proxmox"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type NodeResponse struct {
	Status string  `json:"status"`
	Node   string  `json:"node"`
	MaxCPU int     `json:"max_cpu"`
	MaxMem int     `json:"max_memory"`
	Mem    int     `json:"memory,omitempty"`
	CPU    float64 `json:"cpu_usage,omitempty"`
}

type VMResponse struct {
	NodeResponse
	Name string `json:"name"`
}

func GetNodes(c echo.Context) error {
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	id := c.QueryParam("id")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get service from database
	service, err := service.FetchServiceWithUser(parsedId, database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get nodes from proxmox
	nodes, err := proxmox.ListNodes(service)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// map nodes to NodeResponse
	response := make([]NodeResponse, len(nodes))
	for i, node := range nodes {
		response[i] = NodeResponse{
			Status: node.Status,
			Node:   node.Node,
			MaxCPU: node.MaxCPU,
			MaxMem: int(node.MaxMem / 1024 / 1024 / 1024),
			Mem:    int(node.Mem / 1024 / 1024 / 1024),
			CPU:    node.CPU,
		}
	}

	return c.JSON(200, types.Response{Data: response})
}

func GetVirtualMachines(c echo.Context) error {
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	id := c.QueryParam("id")
	node := c.QueryParam("node")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get service from database
	service, err := service.FetchServiceWithUser(parsedId, database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get virtual machines from a node in proxmox
	vms, err := proxmox.ListVMs(service, node)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// map vms to VMResponse
	response := make([]VMResponse, len(vms))
	for i, vm := range vms {
		response[i] = VMResponse{
			Name: vm.Name,
			NodeResponse: NodeResponse{
				Status: vm.Status,
				Node:   vm.Node,
				MaxCPU: vm.CPUs,
				MaxMem: int(vm.MaxMem / 1024 / 1024 / 1024),
				Mem:    int(vm.Mem / 1024 / 1024 / 1024),
				CPU:    vm.CPU,
			},
		}
	}

	return c.JSON(200, types.Response{Data: response})
}

func GetContainers(c echo.Context) error {
	database := utils.GetLocal[*gorm.DB](c, "dbClient")

	id := c.QueryParam("id")
	node := c.QueryParam("node")

	parsedId, err := strconv.Atoi(id)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get service from database
	service, err := service.FetchServiceWithUser(parsedId, database)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// Get containers from a node in proxmox
	containers, err := proxmox.ListContainers(service, node)
	if err != nil {
		return helpers.ReturnUnexpectedError(c, []string{err.Error()})
	}

	// map containers to VMResponse
	response := make([]VMResponse, len(containers))
	for i, container := range containers {
		response[i] = VMResponse{
			Name: container.Name,

			NodeResponse: NodeResponse{
				Status: container.Status,
				Node:   container.Node,
				MaxCPU: container.CPUs,
				MaxMem: int(container.MaxMem / 1024 / 1024 / 1024),
			},
		}
	}

	return c.JSON(200, types.Response{Data: response})
}
