package health

import (
	"context"
	"github.com/Elyart-Network/NyaBot/config"
)

func Check(ctx context.Context) (map[string]string, map[string]string) {
	status := make(map[string]string)
	errors := make(map[string]string)

	// Check database
	dbType := config.Get("database.type").(string)
	dbs, dbe := Database(ctx)
	status[dbType] = dbs
	if dbe != nil {
		errors[dbType] = dbe.Error()
	}

	// Check redis
	rds, rde := Redis(ctx)
	status["redis"] = rds
	if rde != nil {
		errors["redis"] = rde.Error()
	}

	// Check mongo
	mgs, mge := Mongo(ctx)
	status["mongo"] = mgs
	if mge != nil {
		errors["mongo"] = mge.Error()
	}

	return status, errors
}
