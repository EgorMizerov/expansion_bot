package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/EgorMizerov/expansion_bot/internal/domain/entity"
	"github.com/EgorMizerov/expansion_bot/internal/interface/telebot/template"

	"github.com/go-redis/redis"
)

var (
	stateKey              = "%d:state"
	driverRegistrationKey = "driver_registration"
)

type StateMachine struct {
	client *redis.Client
}

func NewStateMachine(client *redis.Client) *StateMachine {
	return &StateMachine{client: client}
}

func (self *StateMachine) SetState(userID int64, state entity.State) {
	key := fmt.Sprintf(stateKey, userID)
	self.client.Set(key, int(state), 0)
}

func (self *StateMachine) GetState(userID int64) (entity.State, error) {
	data, err := self.client.Get(fmt.Sprintf(stateKey, userID)).Int()
	if err != nil {
		return 0, err
	}
	return entity.State(data), nil
}

func (self *StateMachine) Clear(userId int64) {
	self.client.Del(fmt.Sprintf(stateKey, userId))
}

func (self *StateMachine) SaveRegistrationData(ctx context.Context, data template.DriverRegistrationData) error {
	value, _ := json.Marshal(data)
	return self.client.Set(driverRegistrationKey, value, 0).Err()
}

func (self *StateMachine) GetRegistrationData(ctx context.Context) (result template.DriverRegistrationData, err error) {
	p := self.client.Get(driverRegistrationKey)
	if p.Err() != nil {
		return template.DriverRegistrationData{}, err
	}

	err = json.Unmarshal([]byte(p.Val()), &result)
	return result, err
}
