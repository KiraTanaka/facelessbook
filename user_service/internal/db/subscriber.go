package db

import (
	_ "embed"
	"fmt"
)

//go:embed queries/subscriber/getListSubscribers.sql
var getListSubscribersQuery string

func (r *Repository) ListSubscribers(publisherId string) ([]string, error) {
	subscriberIds := []string{}
	if err := r.db.Select(&subscriberIds, getListSubscribersQuery, publisherId); err != nil {
		return nil, fmt.Errorf("get list of subscribers: %w", err)
	}
	return subscriberIds, nil
}
