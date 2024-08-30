package restaurant

import (
	urlpkg "net/url"
	"unicode/utf8"

	errDomain "github.com/taga3s/pecopeco-service/internal/domain/error"
	ulid "github.com/taga3s/pecopeco-service/pkg/uild"
)

type Restaurant struct {
	ID             string
	Name           string
	Genre          string
	NearestStation string
	Address        string
	URL            string
	UserID         string
}

func NewRestaurant(
	name string,
	genre string,
	nearestStation string,
	address string,
	url string,
	userID string,
) (*Restaurant, error) {
	// 店舗名のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("店舗名の値が不正です。")
	}
	// ジャンルのバリデーション
	if utf8.RuneCountInString(genre) < genreLengthMin || utf8.RuneCountInString(genre) > genreLengthMax {
		return nil, errDomain.NewError("ジャンルの値が不正です。")
	}
	// 最寄り駅のバリデーション
	if utf8.RuneCountInString(nearestStation) < nearestStationLengthMin || utf8.RuneCountInString(nearestStation) > nearestStationLengthMax {
		return nil, errDomain.NewError("最寄り駅の値が不正です。")
	}
	// 住所のバリデーション
	if utf8.RuneCountInString(address) < addressLengthMin || utf8.RuneCountInString(address) > addressLengthMax {
		return nil, errDomain.NewError("住所の値が不正です。")
	}
	// URLのフォーマットのチェック
	if _, err := urlpkg.ParseRequestURI(url); err != nil {
		return nil, errDomain.NewError("URLの値が不正です。")
	}
	return &Restaurant{
		ID:             ulid.NewULID(),
		Name:           name,
		Genre:          genre,
		NearestStation: nearestStation,
		Address:        address,
		URL:            url,
		UserID:         userID,
	}, nil
}

const (
	nameLengthMax = 100
	nameLengthMin = 1
)

const (
	genreLengthMax = 30
	genreLengthMin = 1
)

const (
	nearestStationLengthMax = 30
	nearestStationLengthMin = 1
)

const (
	addressLengthMax = 100
	addressLengthMin = 1
)
