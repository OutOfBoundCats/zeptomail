package zeptomail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/OutOfBoundCats/zeptomail/utils"
	"github.com/OutOfBoundCats/zeptomail/zmodels"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

const zeptoBaseUrl = "https://api.zeptomail.in/v1.1"
const emailEndpoint = "/email"
const templateEmailEndpoint = "/email/template"
const templateBatchEmailEndpoint = "/email/template/batch"

type ZeptoClient struct {
	ZeptoMailerToken      string
	ZeptoMailerAgentAlias string
	Baseurl               string
}

func New(token string, agentAlias string) ZeptoClient {
	return ZeptoClient{
		ZeptoMailerToken:      token,
		ZeptoMailerAgentAlias: agentAlias,
		Baseurl:               zeptoBaseUrl,
	}
}

func makeRequest[T any](client *ZeptoClient, reqURl string, method string, data ...T) (*zmodels.SuccessResponse, *zmodels.FailureResponse, error) {
	finalReqUrl := client.Baseurl + reqURl
	httpClient := http.DefaultClient
	var result zmodels.SuccessResponse
	var errorResult zmodels.FailureResponse

	var body io.Reader

	if data != nil {
		bb, jsonErr := json.Marshal(data[0])
		if jsonErr != nil {
			return nil, nil, jsonErr
		}
		//fmt.Println(string(bb))
		body = bytes.NewBuffer(bb)
	}
	req, reqErr := http.NewRequest(method, finalReqUrl, body)
	if reqErr != nil {
		return nil, nil, reqErr
	} else {
		//add headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Zoho-enczapikey %v", client.ZeptoMailerToken))
	}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	respBody := res.Body
	bytesD, err := io.ReadAll(respBody)
	jsonDecodeErr := json.NewDecoder(bytes.NewReader(bytesD)).Decode(&result)
	if jsonDecodeErr != nil {
		fmt.Println("Error Decoding Zepto success response trying with err resp" + jsonDecodeErr.Error())
	}

	failureDecodeErr1 := json.NewDecoder(bytes.NewReader(bytesD)).Decode(&errorResult)
	if failureDecodeErr1 != nil {
		fmt.Println("!!!!Error Decoding Zepto error response !!!!" + failureDecodeErr1.Error() + jsonDecodeErr.Error())
	}
	if jsonDecodeErr != nil && failureDecodeErr1 != nil {
		return nil, nil, failureDecodeErr1
	}
	return &result, &errorResult, nil
}

func SendSimpleEmail(client *ZeptoClient, data zmodels.EmailData) (*zmodels.SuccessResponse, *zmodels.FailureResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if v := validate.Struct(data); v != nil {
		return nil, nil, v
	}
	resp, errResp, err := makeRequest[zmodels.EmailData](client, emailEndpoint, utils.MethodPOST, data)
	return resp, errResp, err
}

func SendTemplateEmail[T any](client *ZeptoClient, data zmodels.ZeptoTemplateEmail[T]) (*zmodels.SuccessResponse, *zmodels.FailureResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if v := validate.Struct(data); v != nil {
		return nil, nil, v
	}
	resp, errResp, err := makeRequest[zmodels.ZeptoTemplateEmail[T]](client, templateEmailEndpoint, utils.MethodPOST, data)
	return resp, errResp, err
}

func SendTemplateBatchEmail[T any](client *ZeptoClient, data zmodels.TemplateEmailBatch[T]) (*zmodels.SuccessResponse, *zmodels.FailureResponse, error) {
	validate := validator.New(validator.WithRequiredStructEnabled())
	if v := validate.Struct(data); v != nil {
		return nil, nil, v
	}
	resp, errResp, err := makeRequest[zmodels.TemplateEmailBatch[T]](client, templateBatchEmailEndpoint, utils.MethodPOST, data)
	return resp, errResp, err
}
