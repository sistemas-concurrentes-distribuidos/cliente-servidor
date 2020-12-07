package provider

import (
	"client/dto"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const urlBase string = "http://localhost:8000/"

//ReturnProcess teste
func ReturnProcess(IDprocess int) (dto.ProcResponse, error) {
	bytes, code, err := call("process/collector/" + strconv.Itoa(IDprocess))

	if err != nil {
		return dto.ProcResponse{}, err
	}

	var response dto.ProcResponse
	json.Unmarshal(bytes, &response)
	response.StatusCode = code
	return response, nil
}

//GetProcess test
func GetProcess() (dto.ProcResponse, error) {
	bytes, code, err := call("process/dispatcher")

	if err != nil {
		return dto.ProcResponse{}, err
	}

	var response dto.ProcResponse
	json.Unmarshal(bytes, &response)
	response.StatusCode = code
	return response, nil
}

//Call test
func call(endpoint string) ([]byte, int, error) {
	//Get proces from server
	response, err := http.Get(urlBase + endpoint)
	if err != nil {
		return nil, 0, err
	}

	//We Read the response body on the line below.
	body, errRead := ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Print(err)
		return nil, 0, err
	}
	return body, response.StatusCode, nil
}
