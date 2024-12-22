package db

import (
	_ "embed"
	"fmt"
)

//go:embed queries/subscriber/subscribe.sql
var subscribeQuery string

//go:embed queries/subscriber/unsubscribe.sql
var unsubscribeQuery string

//go:embed queries/subscriber/getListSubscribers.sql
var getListSubscribersQuery string

func (r *Repository) Subscribe(publisher_id, subscriber_id string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("begin a transaction for subscribe: %w", err)
	}
	defer tx.Rollback()
	if _, err = tx.Exec(subscribeQuery, publisher_id, subscriber_id); err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}
	tx.Commit()
	return nil
}

func (r *Repository) Unsubscribe(publisher_id, subscriber_id string) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("begin a transaction for unsubscribe: %w", err)
	}
	defer tx.Rollback()
	if _, err = tx.Exec(unsubscribeQuery, publisher_id, subscriber_id); err != nil {
		return fmt.Errorf("unsubscribe: %w", err)
	}
	tx.Commit()
	return nil
}

func (r *Repository) ListSubscribers(publisherId string) ([]string, error) {
	subscriberIds := []string{}
	if err := r.db.Select(&subscriberIds, getListSubscribersQuery, publisherId); err != nil {
		return nil, fmt.Errorf("get list of subscribers: %w", err)
	}
	return subscriberIds, nil
}
