package userinfo

import (
	"context"

	"github.com/Seiya-Tagami/pecopeco-cli/auth"
	"github.com/Seiya-Tagami/pecopeco-cli/auth/api/model"
	api "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func Get(ctx context.Context, oauth auth.OAuth) (model.Userinfo, error) {
	client := oauth.Config.Client(ctx, oauth.Token)
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
