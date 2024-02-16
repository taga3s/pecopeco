package userinfo

import (
	"context"
	"time"

	"github.com/ayanami77/pecopeco-cli/auth"
	"github.com/ayanami77/pecopeco-cli/auth/api/model"
	api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func Get(ctx context.Context, oauth auth.OAuth) (model.Userinfo, error) {
	client := oauth.Config.Client(ctx, oauth.Token)
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	service, err := api.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return model.Userinfo{}, err
	}
	res, err := service.Userinfo.Get().Context(ctx).Do()
	if err != nil {
		return model.Userinfo{}, err
	}
	userinfo := model.Userinfo{
		ID:    res.Id,
		Name:  res.Name,
		Email: res.Email,
	}
	return userinfo, nil
}
