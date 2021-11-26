package app

import (
	"go.m3o.com/client"
)

func NewAppService(token string) *AppService {
	return &AppService{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type AppService struct {
	client *client.Client
}

// Reserve your app name
func (t *AppService) Reserve(request *ReserveRequest) (*ReserveResponse, error) {

	rsp := &ReserveResponse{}
	return rsp, t.client.Call("app", "Reserve", request, rsp)

}

type Reservation struct {
	// time of reservation
	Created string `json:"created"`
	// time reservation expires
	Expires string `json:"expires"`
	// name of the app
	Name string `json:"name"`
	// owner id
	Owner string `json:"owner"`
	// associated token
	Token string `json:"token"`
}

type ReserveRequest struct {
	// name of your app e.g helloworld
	Name string `json:"name"`
}

type ReserveResponse struct {
	Reservation *Reservation `json:"reservation"`
}
