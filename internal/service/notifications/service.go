package notifications

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"

	"github.com/mig8at/notifications/internal/model"
	"github.com/mig8at/notifications/pkg"
	"github.com/spf13/viper"
)

type NotificationsServiceInterface interface {
	Send(message string) error
}

type NotificationsService struct {
	cache  *pkg.UserCache
	limits []model.Limit
}

func Service() *NotificationsService {

	viper.SetConfigName("config")

	_, filename, _, _ := runtime.Caller(0)

	configPath := filepath.Join(filepath.Dir(filename), "..", "..", "..", "config")
	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading config file: %s\n", err)
	}

	var limits []model.Limit
	if err := viper.UnmarshalKey("limits", &limits); err != nil {
		fmt.Printf("Error decoding config into struct: %s\n", err)
	}

	cache := pkg.NewUserCache()
	return &NotificationsService{
		cache:  cache,
		limits: limits,
	}
}

func (s NotificationsService) Message(nType model.NotificationType, userId int) (NotificationsServiceInterface, error) {

	var limit *model.Limit
	for _, l := range s.limits {
		if l.Type == nType {
			limit = &l
			break
		}
	}

	if limit == nil {
		return nil, fmt.Errorf("notification type not supported")
	}

	getUser := func() (*model.User, error) {
		seconds, err := pkg.ConvertToSeconds(limit.Duration)
		if err != nil {
			return nil, err
		}

		secondsToAdd := time.Duration(seconds) * time.Second
		now := time.Now()

		user, exist := s.cache.Get(limit.Type, userId)
		if !exist {
			user = model.User{
				ID:   userId,
				Date: now.Add(secondsToAdd),
			}
			s.cache.Set(limit.Type, userId, user)
		}
		if user.Count >= limit.Count {
			if user.Date.After(now) {
				return nil, fmt.Errorf("user %d has reached the limit", userId)
			}
			user.Count = 0
			user.Date = now.Add(secondsToAdd)
			s.cache.Set(limit.Type, userId, user)
		} else {
			user.Count++
			s.cache.Set(limit.Type, userId, user)
		}
		return &user, nil
	}

	switch nType {
	case model.Status:
		return &StatusNotification{getUser: getUser}, nil
	case model.News:
		return &NewsNotification{getUser: getUser}, nil
	case model.Marketing:
		return &MarketingNotification{getUser: getUser}, nil
	default:
		return nil, fmt.Errorf("notification type not supported")
	}
}
