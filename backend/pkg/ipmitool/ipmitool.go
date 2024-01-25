package ipmitool

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

const ipmiToolCommand = "ipmitool"

type Client struct {
	address  string
	port     uint16
	username string
	password string
	inter    string

	Power  *Power
	Sensor *Sensor
	FRU    *FRU
}

type Device struct {
	ProductManufacturer string     `json:"product_manufacturer"`
	ProductModel        string     `json:"product_model"`
	ProductSerial       string     `json:"product_serial"`
	PowerState          PowerState `json:"power_state"`
	PowerConsumption    int        `json:"power_consumption"`
	CPUTemperature      int        `json:"cpu_temperature"`
}

func NewClient(address string, port uint16, username, password string) (*Client, error) {
	if port == 0 {
		port = 623
	}

	cl := &Client{
		address:  address,
		port:     port,
		username: username,
		password: password,
		inter:    "lanplus",
	}

	cl.Power = NewPower(cl)
	cl.Sensor = NewSensor(cl)
	cl.FRU = NewFRU(cl)

	return cl, nil
}

func (cl *Client) getBaseParam() []string {
	params := []string{
		"-I", cl.inter,
	}

	if cl.address != "" {
		params = append(params, "-H", cl.address)
	}

	if cl.port != 0 {
		params = append(params, "-p", strconv.Itoa(int(cl.port)))
	}

	if cl.username != "" {
		params = append(params, "-U", cl.username)
	}

	if cl.password != "" {
		params = append(params, "-P", cl.password)
	}

	return params
}

func (cl *Client) execute(args []string) (string, error) {
	cmd := exec.Command(ipmiToolCommand, args...)
	var outBuf, errBuf bytes.Buffer
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %v", err)
	}

	return outBuf.String(), nil
}

func (cl *Client) GetInfo() Device {
	errorCh := make(chan error, 1)
	statusCh := make(chan PowerState, 1)
	powerReadingCh := make(chan string, 1)
	cpuTemperatureCh := make(chan string, 1)
	fruCh := make(chan Inventory, 1)

	go cl.Power.Status(statusCh, errorCh)
	go cl.Sensor.PowerReading(powerReadingCh, errorCh)
	go cl.Sensor.CPUTemperature(cpuTemperatureCh, errorCh)
	go cl.FRU.FRUInfo(fruCh, errorCh)

	status := <-statusCh
	powerReading := <-powerReadingCh
	cpuTemperature := <-cpuTemperatureCh
	fru := <-fruCh

	powerReadingInt, _ := strconv.Atoi(powerReading)
	cpuTemperatureInt, _ := strconv.Atoi(cpuTemperature)

	defer close(errorCh)
	defer close(statusCh)
	defer close(powerReadingCh)
	defer close(cpuTemperatureCh)
	defer close(fruCh)

	return Device{
		ProductManufacturer: fru.Manufacturer,
		ProductModel:        fru.Product,
		ProductSerial:       fru.Serial,
		PowerState:          status,
		PowerConsumption:    powerReadingInt,
		CPUTemperature:      cpuTemperatureInt,
	}
}
