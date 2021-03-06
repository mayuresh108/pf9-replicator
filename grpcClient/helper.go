//package grpcClient
package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	//"errors"
	_ "embed"
	"strings"
)

var (
	//go:embed command.conf.json
	tmpCmdConf string
	cmdJsonData string = strings.TrimSpace(tmpCmdConf)
	cmdList []cmd
	err error = json.Unmarshal([]byte(cmdJsonData), &cmdList)
)


func readCmdConfig() ([]byte, error) {
	byteData, readError := ioutil.ReadFile("./commad.conf.json")
    if readError != nil {
		fmt.Printf("ERROR: %s\n", readError.Error())
        return []byte{}, readError
    }

    return byteData, nil
}


func generateCmdData() error {
	/* cmdJsonData, err := readCmdConfig()
	if err != nil {
		fmt.Printf("read cmd config file error: %s\n", err.Error())
		return errors.New("Error in reading cmd configuration.")
	}

	cmdList := []cmd{}
	if err := json.Unmarshal(cmdJsonData, &cmdList); err != nil {
		fmt.Printf("Malformed cmd config: %s\n", err.Error())
		return errors.New("Malformed cmd configuration.")
	} */

	for _, rec := range cmdList {
		cmdData[rec.Name] = rec
	}

	return nil
}
