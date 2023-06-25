package health

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/data/actions"
)

func Database(ctx context.Context) (string, error) {
	DB, _ := actions.CONN.DB.WithContext(ctx).DB()
	err := DB.Ping()
	if err != nil {
		return "error", err
	}
	return "healthy", err
}

func Redis(ctx context.Context) (string, error) {
	externalCache := config.Get("cache.external").(bool)
	if externalCache {
		err := actions.CONN.Redis.Ping(ctx).Err()
		if err != nil {
			return "error", err
		}
	} else {
		return "disabled", nil
	}
	return "healthy", nil
}

func Mongo(ctx context.Context) (string, error) {
	externalLogging := config.Get("logging.external").(bool)
	if externalLogging {
		err := actions.CONN.Mongo.Ping(ctx, nil)
		if err != nil {
			return "error", err
		}
	} else {
		return "disabled", nil
	}
	return "healthy", nil
}
