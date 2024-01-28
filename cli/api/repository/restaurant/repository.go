package restaurant

import (
	"fmt"

	"github.com/Seiya-Tagami/pecopeco-cli/api/client/hotpepper"
	"github.com/Seiya-Tagami/pecopeco-cli/api/client/linenotify"
	"github.com/Seiya-Tagami/pecopeco-cli/api/client/util"
)

type Repository interface {
	List(request ListRequest) (ListResponse, error)
	NotifyToLINE(request NotifyToLINERequest) error
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) List(request ListRequest) (ListResponse, error) {
	listResponse := ListResponse{}
	if err := hotpepper.HttpClient("GET", fmt.Sprintf("&keyword=%s,%s&count=100&format=json", request.City, request.Food), &listResponse); err != nil {
		return ListResponse{}, err
	}
	return listResponse, nil
}

var body map[string]string

func (r *repository) NotifyToLINE(request NotifyToLINERequest) error {
	notifyToLINEResponse := NotifyToLINEResponse{}
	if err := linenotify.HttpClient("POST", fmt.Sprintf(
		"\n---------------------\n[店名] %s\n[住所] %s\n[最寄り駅] %s\n[ジャンル] %s\n[URL] %s\n---------------------\n",
		request.Name,
		request.Address,
		request.StationName,
		request.GenreName,
		request.URL,
	), &notifyToLINEResponse); err != nil {
		return err
	}
	if err := util.CheckStatus(notifyToLINEResponse.Status); err != nil {
		return err
	}
	return nil
}
