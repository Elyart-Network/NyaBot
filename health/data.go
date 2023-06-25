package health

import (
	"context"
	"errors"
	"github.com/Elyart-Network/NyaBot/config"
	"github.com/Elyart-Network/NyaBot/data/actions"
)

func Database(ctx context.Context) (string, error) {
	if actions.CONN.DB == nil {
		return "failed", errors.New("failed to get database connection from CONN")
	}
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
		if actions.CONN.Redis == nil {
			return "failed", errors.New("failed to get redis connection from CONN")
		}
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
		if actions.CONN.Mongo == nil {
			return "failed", errors.New("failed to get mongo connection from CONN")
		}
		err := actions.CONN.Mongo.Ping(ctx, nil)
		if err != nil {
			return "error", err
		}
	} else {
		return "disabled", nil
	}
	return "healthy", nil
}
