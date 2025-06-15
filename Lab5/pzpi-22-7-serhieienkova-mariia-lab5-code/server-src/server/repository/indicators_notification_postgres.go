package repository

import (
	"clinic/server/structures"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type IndicatorsNotificationPostgres struct {
	db *sqlx.DB
}

func NewIndicatorsNotificationPostgres(db *sqlx.DB) *IndicatorsNotificationPostgres {
	return &IndicatorsNotificationPostgres{db: db}
}

func (r *IndicatorsNotificationPostgres) Create(indicatorsNotification structures.IndicatorsNotification) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (indicator_stamp_id, message, timestamp) values ($1, $2, $3) RETURNING id", indicatorsNotificationsTable)
	row := r.db.QueryRow(query, indicatorsNotification.IndicatorStampId, indicatorsNotification.Message, indicatorsNotification.Timestamp)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *IndicatorsNotificationPostgres) GetAll() ([]structures.IndicatorsNotification, error) {
	var indicatorsNotifications []structures.IndicatorsNotification
	query := fmt.Sprintf("SELECT * FROM %s", indicatorsNotificationsTable)
	err := r.db.Select(&indicatorsNotifications, query)
	if err != nil {
		return nil, err
	}
	return indicatorsNotifications, nil
}

func (r *IndicatorsNotificationPostgres) Get(id int) (structures.IndicatorsNotification, error) {
	var indicatorsNotification structures.IndicatorsNotification
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", indicatorsNotificationsTable)
	err := r.db.Get(&indicatorsNotification, query, id)
	if err != nil {
		return indicatorsNotification, err
	}
	return indicatorsNotification, nil
}

func (r *IndicatorsNotificationPostgres) GetAllByPatientID(patientID int) ([]structures.IndicatorsNotification, error) {
	var indicatorsNotifications []structures.IndicatorsNotification
	query := fmt.Sprintf(`
        SELECT n.id, n.indicator_stamp_id, n.message, n.timestamp
        FROM %s n
        JOIN %s i ON n.indicator_stamp_id = i.id
        WHERE i.device_id IN (SELECT id FROM %s WHERE patient_id = $1)
    `, indicatorsNotificationsTable, indicatorsStampsTable, devicesTable)
	err := r.db.Select(&indicatorsNotifications, query, patientID)
	if err != nil {
		return nil, err
	}
	return indicatorsNotifications, nil
}
