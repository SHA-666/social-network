package repositories

import (
	"context"
	"database/sql"
	"fmt"

	"social-network/core/entities"
	repositories "social-network/core/repositories"
)

type FriendRepo struct {
	db *sql.DB
}

func (f *FriendRepo) Get(ctx context.Context, userId string) ([]entities.Friends, error) {
	q := `
	SELECT 
		u.id,
		u.first_name,
		u.last_name,
		COALESCE(u.profile_pic, '') as profile_pic
	FROM users u
	WHERE u.id IN (
		SELECT receiver_id FROM friends 
		WHERE sender_id = ? AND validated = 1
		UNION
		SELECT sender_id FROM friends 
		WHERE receiver_id = ? AND validated = 1
	)
	ORDER BY u.first_name
	`
	rows, err := f.db.Query(q, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var friends []entities.Friends
	for rows.Next() {
		var friend entities.Friends
		err := rows.Scan(
			&friend.Id,
			&friend.FirstName,
			&friend.LastName,
			&friend.ProfilePic,
		)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	return friends, nil
}


func NewFriendRepo(db *sql.DB) repositories.FriendsRepo {
	return &FriendRepo{db: db}
}

// Delete implements repositories.FriendsRepo.
func (f *FriendRepo) Delete(ctx context.Context, receiver, sender string) error {
	fmt.Println(sender, receiver, "ok")
	_, err := f.db.ExecContext(ctx, `DELETE FROM friends WHERE receiver_id = ? AND sender_id = ?`, receiver, sender)
	if err != nil {
		fmt.Println("Delete error:", err)
		return err
	}
	return nil
}

func (f *FriendRepo) Add(ctx context.Context, userId string, receiverId string) error {
	_, err := f.db.ExecContext(ctx, `INSERT INTO friends (sender_id,receiver_id)
 	VALUES (?,?)`, userId, receiverId)
	return err
}

func (f *FriendRepo) Accepte(ctx context.Context, userId string, receiverId string) error {
	fmt.Printf("%s sender, //%s receiver \n", userId, receiverId)

	_, err := f.db.ExecContext(ctx, `UPDATE friends 
		SET validated = 1 
		WHERE receiver_id = ? AND sender_id = ? AND validated = 0`,
		userId, receiverId)
	if err != nil {
		fmt.Println("Accept friend error:", err)
		return err
	}
	return nil
}

func (f *FriendRepo) Load(ctx context.Context, userId string) ([]entities.Friends, error) {
	q := `
	SELECT
		u.id,
		u.first_name,
		u.last_name,
		COALESCE(u.profile_pic, '') as profile_pic
	FROM users u
	WHERE u.id != ?
	AND u.id NOT IN (
		SELECT receiver_id FROM friends 
		WHERE sender_id = ? AND validated = 1
		UNION
		SELECT sender_id FROM friends 
		WHERE receiver_id = ? AND validated = 1
		UNION
		SELECT receiver_id FROM friends 
		WHERE sender_id = ? AND validated = 0
		UNION
		SELECT sender_id FROM friends 
		WHERE receiver_id = ? AND validated = 0
	)
	ORDER BY RAND()
	LIMIT 7
	`
	rows, err := f.db.Query(q, userId, userId, userId, userId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var friends []entities.Friends
	for rows.Next() {
		var friend entities.Friends
		err := rows.Scan(
			&friend.Id,
			&friend.FirstName,
			&friend.LastName,
			&friend.ProfilePic,
		)
		if err != nil {
			return nil, err
		}
		friends = append(friends, friend)
	}

	return friends, nil
}


// func (f *FriendRepo) Add(ctx context.Context,userId,receiverId string) error {
// 	_, err := f.db.ExecContext(ctx, `INSERT INTO friends (sender_id,receiver_id)
// 	VALUES (?,?)`, )
// 	return err
// }

// func (f *FriendRepo) Load(currentUserId string) ([]entities.Friends, error) {

// }

// CREATE TABLE friends (
//     id_friend INT PRIMARY KEY AUTO_INCREMENT,
//     sender_id INT NOT NULL,
//     receiver_id INT NOT NULL,
//     date_follow DATETIME DEFAULT CURRENT_TIMESTAMP,
//     validated BOOLEAN DEFAULT FALSE,
//     FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
//     FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
// );

// CREATE TABLE private_messages (
//     message_id INT PRIMARY KEY AUTO_INCREMENT,
//     sender_id INT NOT NULL,
//     receiver_id INT NOT NULL,
//     date DATETIME DEFAULT CURRENT_TIMESTAMP,
//     content TEXT NOT NULL,
//     image BLOB,
//     is_read BOOLEAN DEFAULT FALSE,
//     FOREIGN KEY (sender_id) REFERENCES users(id) ON DELETE CASCADE,
//     FOREIGN KEY (receiver_id) REFERENCES users(id) ON DELETE CASCADE
// );

// 	q := `
// 	SELECT
// 		p.id,
// 		p.user_id,
// 		u.username,
// 		COALESCE(p.title, '') as title,
// 		COALESCE(u.profile_pic, '') as user_pic,
// 		COALESCE(p.content, '') as content,
// 		p.privacy,
// 		COALESCE(p.img, '') as img,
// 		p.createdAt
// 	FROM posts p
// 	INNER JOIN users u ON p.user_id = u.id
// 	WHERE p.privacy = 'public'
// 	ORDER BY p.createdAt DESC
// `
// 	rows, err := p.db.QueryContext(ctx, q)
// 	if err != nil {
// 		return nil, "", err
// 	}
// 	defer rows.Close()

// 	var posts []*entities.Post
// 	for rows.Next() {
// 		var post entities.Post
// 		if err := rows.Scan(&post.Id,&post.UserId,&post.Username,&post.Title,&post.UserPic,&post.Content,&post.Privacy,&post.Image,&post.Date,
// 		); err != nil {
// 			return nil, "", err
// 		}
// 		posts = append(posts, &post)
// 	}
