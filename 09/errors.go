package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Health struct {
	Status     string `json:"status"`
	Version    string `json:"version"`
	ReleasedAt string `json:"releasedAt"`
}

func getHealth() (Health, error) {
	resp, err := http.Get("https://app.kenticocloud.com/api/health")
	if err != nil {
		return Health{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Health{}, err
	}
	var health Health
	if err := json.Unmarshal(body, &health); err != nil {
		return Health{}, err
	}
	return health, nil
}

func getVersion(health Health) (string, error) {
	if health.Status != "Ok" {
		return "", errors.New("Health is not ok")
	}
	return health.Version, nil
}

func main() {
	criticalWorkRecovery()

	health, err := getHealth()
	if err != nil {
		fmt.Println(err)
		return
	}

	version, err := getVersion(health)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(version)
}

func criticalWork() {
	somethingGoesTerriblyWrong := false
	if somethingGoesTerriblyWrong {
		panic("this should never happen")
	}
}

func criticalWorkRecovery() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("recovered from: %s\n", err)
		}
	}()
	criticalWork()
}
