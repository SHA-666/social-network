package repositories

import (
	"database/sql"
	"fmt"

	entities "social-network/core/entities"
	repositories "social-network/core/repositories"
)

type MessageRepo struct {
	db *sql.DB
}

func NewMessageRepo(db *sql.DB) repositories.MessageRepo {
	return &MessageRepo{db: db}
}

func (m *MessageRepo) Load(user1 string, user2 string) ([]entities.Message, error) {
	query := `
		SELECT id_message, id_sender, id_receiver, content_message, date
		FROM messages
		WHERE (id_sender = ? AND id_receiver = ?)
		   OR (id_sender = ? AND id_receiver = ?)
		ORDER BY date ASC;
	`

	rows, err := m.db.Query(query, user1, user2, user2, user1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entities.Message
	for rows.Next() {
		var msg entities.Message
		err := rows.Scan(
			&msg.ID,
			&msg.From,
			&msg.To,
			&msg.Content,
			&msg.SendAt,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println(messages)
	return messages, nil
}

func (m *MessageRepo) Save(From, To, Content string) error {
	_, err := m.db.Exec(`INSERT INTO messages (id_sender,id_receiver,content_message)
 	VALUES (?,?,?)`, From, To, Content)
	fmt.Println(err)
	return err
}
