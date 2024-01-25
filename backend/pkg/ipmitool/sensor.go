package ipmitool

import (
	"fmt"
	"regexp"
	"strings"
)

// PowerConsumption represents an IPMI node/server power reading
type PowerConsumption string

// NewSensor returns a new instance of the ipmi dcmi sub command
func NewSensor(cl *Client) *Sensor {
	return &Sensor{
		cl: cl,
	}
}

// Sensor represents the ipmi dcmi command
type Sensor struct {
	cl *Client
}

// PowerReading fetches the current power reading of the ipmi server
func (p *Sensor) PowerReading(s chan string, e chan error) {
	params := p.cl.getBaseParam()
	params = append(params, "dcmi", "power", "reading")

	stdout, err := p.cl.execute(params)
	if err != nil {
		e <- fmt.Errorf("failed to fetch power reading: %w", err)
	}

	output := strings.Split(stdout, "\n")

	for _, line := range output {
		if strings.Contains(line, "power reading") {
			re := regexp.MustCompile(`(\d+)\sWatts`)
			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				s <- matches[1]
			}
		}
	}
}

// CPUTemperature fetches the current cpu temperature of the ipmi server
func (p *Sensor) CPUTemperature(s chan string, e chan error) {
	params := p.cl.getBaseParam()
	params = append(params, "dcmi", "get_temp_reading")

	stdout, err := p.cl.execute(params)
	if err != nil {
		e <- fmt.Errorf("failed to fetch cpu temperature: %w", err)
	}

	output := strings.Split(stdout, "\n")

	for _, line := range output {
		if strings.Contains(line, "CPU") {
			re := regexp.MustCompile(`\+(\d+) C`)
			matches := re.FindStringSubmatch(line)
			if len(matches) > 0 {
				s <- matches[1]
				break
			}
		}
	}
}
