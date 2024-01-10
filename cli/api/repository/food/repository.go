package food

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Repository interface {
	GetFoodList(request GetFoodListRequest) (GetFoodListResponse, error)
}

type repository struct{}

func New() Repository {
	return &repository{}
}

func (r *repository) GetFoodList(request GetFoodListRequest) (GetFoodListResponse, error) {
	uri := os.Getenv("HOTPEPPER_API_URL")
	key := os.Getenv("HOTPEPPER_API_KEY")

	fmt.Println(uri + "/?key=" + key + "&keyword=" + request.City + "," + request.Food + "&count=100&format=json")
	req, _ := http.NewRequest("GET", uri+"/?key="+key+"&keyword="+request.City+","+request.Food+"&count=100&format=json", nil)

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return GetFoodListResponse{}, err
	}
	defer res.Body.Close()

	byteArray, _ := io.ReadAll(res.Body)
	getFoodListResponse := GetFoodListResponse{}
	if err := json.Unmarshal(byteArray, &getFoodListResponse); err != nil {
		return GetFoodListResponse{}, err
	}
	return getFoodListResponse, nil
}
