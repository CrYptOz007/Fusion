package ipmitool

import (
	"fmt"
	"regexp"
	"strings"
)

// Inventory represents an IPMI node/server inventory
type Inventory struct {
	Manufacturer string `json:"manufacturer"`
	Product      string `json:"product"`
	Serial       string `json:"serial"`
}

// NewFRU returns a new instance of the ipmi fru sub command
func NewFRU(cl *Client) *FRU {
	return &FRU{
		cl: cl,
	}
}

// Sensor represents the ipmi fru command
type FRU struct {
	cl *Client
}

// Status fetches the fru inventory of the ipmi server
func (p *FRU) FRUInfo(i chan Inventory, e chan error) {
	params := p.cl.getBaseParam()
	params = append(params, "fru", "print", "0")
	inventory := Inventory{}

	stdout, err := p.cl.execute(params)
	if err != nil {
		e <- fmt.Errorf("failed to fetch fru inventory: %w", err)
	}

	output := strings.Split(stdout, "\n")

	for _, line := range output {
		re := regexp.MustCompile(`Product Manufacturer\s*:\s*(.*)`)
		matches := re.FindStringSubmatch(line)
		if len(matches) > 0 {
			inventory.Manufacturer = matches[1]
		}

		re = regexp.MustCompile(`Product Name\s*:\s*(.*)`)
		matches = re.FindStringSubmatch(line)
		if len(matches) > 0 {
			inventory.Product = matches[1]
		}

		re = regexp.MustCompile(`Product Serial\s*:\s*(.*)`)
		matches = re.FindStringSubmatch(line)
		if len(matches) > 0 {
			inventory.Serial = matches[1]
		}
	}

	i <- inventory
}
