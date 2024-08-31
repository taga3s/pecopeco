package notifytoline

import (
	"fmt"

	"github.com/taga3s/pecopeco-cli/api/client/notify"
	"github.com/taga3s/pecopeco-cli/api/client/util"
)

type Repository interface {
	NotifyToLINE(request NotifyToLINERequest) error
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) NotifyToLINE(request NotifyToLINERequest) error {
	notifyToLINEResponse := NotifyToLINEResponse{}
	if err := notify.HttpClient("POST", fmt.Sprintf(
		"\n---------------------\n[店名] %s\n[住所] %s\n[最寄り駅] %s\n[ジャンル] %s\n[URL] %s\n---------------------\n",
		request.Name,
		request.Address,
		request.NearestStation,
		request.Genre,
		request.URL,
	), &notifyToLINEResponse); err != nil {
		return err
	}
	if err := util.CheckStatus(notifyToLINEResponse.Status); err != nil {
		return err
	}
	return nil
}
