package form3

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws-master/form3/models"
)

func (svc *accountSvc) FetchAccount(fullPath string) (models.AccountData, error) {
	request, err := http.NewRequest(http.MethodGet, fullPath, nil)
	if err != nil {
		return models.AccountData{}, err
	}

	resp, err := doRequest(request)
	if err != nil {
		return models.AccountData{}, err
	}

	defer resp.Body.Close()

	resBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {

		return models.AccountData{}, err
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode > http.StatusMultipleChoices {
		// log.Println("resp.StatusCode", resp.StatusCode)
		return models.AccountData{}, fmt.Errorf(string(resBody))
	}

	data := struct {
		AccountData models.AccountData `json:"data"`
	}{}

	if err := json.Unmarshal(resBody, &data); err != nil {
		return models.AccountData{}, err
	}

	return data.AccountData, nil
}
