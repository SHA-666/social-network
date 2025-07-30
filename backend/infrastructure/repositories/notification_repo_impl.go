package repositories

import (
	"context"
	"database/sql"

	"social-network/core/entities"
	repositories "social-network/core/repositories"
)

type NotificationRepo struct {
	db *sql.DB
}

func NewNotificationRepo(db *sql.DB) repositories.NotificationRepo {
	return &NotificationRepo{db: db}
}

func (n *NotificationRepo) GetNotifications(ctx context.Context, id string) ([]entities.Notification, error) {
	rows, err := n.db.QueryContext(ctx, `
		SELECT id, to_user, from_user, type, message, created_at
		FROM notifications
		WHERE to_user = ?
		ORDER BY created_at DESC
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []entities.Notification

	for rows.Next() {
		var notif entities.Notification
		if err := rows.Scan(
			&notif.Id,
			&notif.To,
			&notif.From,
			&notif.Type,
			&notif.Message,
			&notif.At,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, notif)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (n *NotificationRepo) NotifUserFriends(ctx context.Context, from string, to string, types string, message string) error {
	_, err := n.db.ExecContext(ctx, `INSERT INTO notifications (to_user,from_user,type,message)
 	VALUES (?,?,?,?)`, to, from, types, message)
	return err
}

func (n *NotificationRepo) DeleteNotif(ctx context.Context, id string) error {
	_, err := n.db.ExecContext(ctx, `DELETE FROM notifications WHERE id = ?`, id)
	return err
}

// | Field      | Type        | Null | Key | Default           | Extra             |
// +------------+-------------+------+-----+-------------------+-------------------+
// | id         | int         | NO   | PRI | NULL              | auto_increment    |
// | to_user    | int         | NO   | MUL | NULL              |                   |
// | from_user  | int         | NO   | MUL | NULL              |                   |
// | type       | varchar(50) | NO   |     | NULL              |                   |
// | message    | text        | NO   |     | NULL              |                   |
// | created_at | timestamp   | YES  |     | CURRENT_TIMESTAMP | DEFAULT_GENERATED |
