package pihole

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/CrYptOz007/Fusion/internal/models/service"
)

func sendRequest(service *service.Service, query string) (map[string]interface{}, error) {
	url := fmt.Sprintf("http://%s:%d/admin/api.php?%s&auth=%s", service.Hostname, service.Port, query, service.ApiKey)
	resp, error := http.Get(url)
	if error != nil {
		return map[string]interface{}{}, error
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]interface{}{}, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return map[string]interface{}{}, err
	}

	return data, nil
}

func Enable(service *service.Service) error {
	_, err := sendRequest(service, "enable")
	if err != nil {
		return fmt.Errorf("failed to enable pihole: %w", err)
	}

	return nil
}

func Disable(service *service.Service) error {
	_, err := sendRequest(service, "disable")
	if err != nil {
		return fmt.Errorf("failed to disable pihole: %w", err)
	}

	return nil
}

func Status(service *service.Service) (map[string]interface{}, error) {
	resp, err := sendRequest(service, "status")
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to fetch pihole status: %w", err)
	}

	return resp, nil
}

func Summary(service *service.Service) (map[string]interface{}, error) {

	resp, err := sendRequest(service, "summary")
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("failed to fetch pihole summary: %w", err)
	}

	return resp, nil
}
