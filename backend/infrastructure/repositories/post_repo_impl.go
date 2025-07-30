package repositories

import (
	"context"
	"database/sql"

	entities "social-network/core/entities"
	repositories "social-network/core/repositories"
)

type PostRepo struct {
	db *sql.DB
}

func NewPostRepo(db *sql.DB) repositories.PostRepo {
	return &PostRepo{db: db}
}

func (p *PostRepo) Save(ctx context.Context, post *entities.Post) error {
	_, err := p.db.ExecContext(ctx, `INSERT INTO posts (user_id,title,content, privacy,img)
	VALUES (?,?,?,?,?)`, post.UserId, post.Title, post.Content, post.Privacy, post.Image)
	return err
}

func (p *PostRepo) Load(ctx context.Context, id string) ([]*entities.Post, string, error) {
	var currentUserPic string
	userPicQuery := `SELECT COALESCE(profile_pic, '') FROM users WHERE id = ?`
	err := p.db.QueryRowContext(ctx, userPicQuery, id).Scan(&currentUserPic)
	if err != nil {
		return nil, "", err
	}
	q := `
	SELECT 
		p.id,		
		p.user_id, 
		u.username, 
		COALESCE(p.title, '') as title,
		COALESCE(u.profile_pic, '') as user_pic,
		COALESCE(p.content, '') as content,
		p.privacy,
		COALESCE(p.img, '') as img,
		p.createdAt
	FROM posts p
	INNER JOIN users u ON p.user_id = u.id  
	WHERE p.privacy = 'public'
	ORDER BY p.createdAt DESC
`
	rows, err := p.db.QueryContext(ctx, q)
	if err != nil {
		return nil, "", err
	}
	defer rows.Close()

	var posts []*entities.Post
	for rows.Next() {
		var post entities.Post
		if err := rows.Scan(&post.Id,&post.UserId,&post.Username,&post.Title,&post.UserPic,&post.Content,&post.Privacy,&post.Image,&post.Date,
		); err != nil {
			return nil, "", err
		}
		posts = append(posts, &post)
	}

	return posts, currentUserPic, nil
}
