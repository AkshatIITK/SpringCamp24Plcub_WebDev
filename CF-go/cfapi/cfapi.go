// codeforces_api.go
package cfapi

import (
	model "cfapiapp/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	apiURL      = "https://codeforces.com/api/recentActions?maxCount=100"
	timeout     = 2
	resultIndex = 2
)

func Cfapicall() ([]model.Result, error) {
	Client := http.Client{
		Timeout: time.Second * timeout,
	}

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return nil, err
	}

	res, getErr := Client.Do(req)
	if getErr != nil {
		log.Printf("Error making request: %v", getErr)
		return nil, getErr
	}
	defer res.Body.Close()

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Printf("Error reading response body: %v", readErr)
		return nil, readErr
	}
	wrapper := struct {
		Status string
		Result []model.Result
	}{}

	jsonErr := json.Unmarshal(body, &wrapper)
	if jsonErr != nil {
		log.Printf("Error unmarshalling JSON: %v", jsonErr)
		return nil, jsonErr
	}

	if wrapper.Status != "OK" {
		log.Printf("API returned non-OK status: %v", wrapper.Status)
		return nil, fmt.Errorf("API returned non-OK status: %v", wrapper.Status)
	}

	if err := saveToFile(wrapper.Result, "file.json"); err != nil {
		return nil, err
	}

	return wrapper.Result, nil

}

func saveToFile(data interface{}, filename string) error {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Error marshalling JSON: %v", err)
		return err
	}

	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Printf("Error writing to file: %v", err)
		return err
	}

	return nil
}
