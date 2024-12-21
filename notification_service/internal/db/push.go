package db

import (
	_ "embed"
	"fmt"
	"notification_service/internal/models"
)

//go:embed queries/push/getPushPattern.sql
var getPushPatternQuery string

//go:embed queries/push/savePush.sql
var savePushQuery string

func (r *Repository) Pattern(patternName string) (*models.PushPattern, error) {
	pushPattern := &models.PushPattern{}
	if err := r.db.Get(pushPattern, getPushPatternQuery, patternName); err != nil {
		return nil, err
	}
	return pushPattern, nil
}

func (r *Repository) SavePush(userIds []string, patternName string, params []any) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	pushPattern, err := r.Pattern(patternName)
	if err != nil {
		return err
	}

	for _, userId := range userIds {
		_, err = tx.Exec(savePushQuery, userId, pushPattern.Subject, fmt.Sprintf(pushPattern.Pushmessage, params...), fmt.Sprintf(pushPattern.FullMessage, params...))
		if err != nil {
			return err
		}
	}
	tx.Commit()
	return nil
}
