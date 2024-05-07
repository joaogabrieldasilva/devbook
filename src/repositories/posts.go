package repositories

import (
	"api/src/models"
	"database/sql"
)

type postsRepository struct {
	db *sql.DB
}


func NewPostsRepository(db *sql.DB) *postsRepository {
	return &postsRepository{
		db,
	}
}


func (repository postsRepository) CreatePost(authorId uint64, title string, content string) (uint64, error) {

	statement, error := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES (?, ?, ?)")

	if error != nil {
		return 0, error
	}

	defer statement.Close()

	result, error := statement.Exec(title, content, authorId)

	if error != nil {
		return 0, error
	}

	createdPostID, error := result.LastInsertId()

	if error != nil {
		return 0, error
	}

	return uint64(createdPostID), nil
}

func (repository postsRepository) GetPostByID(postID uint64) ([]models.Post, error) {

	rows, error := repository.db.Query(`
		SELECT p.*, u.username from posts p JOIN users u ON p.author_id = u.id
		WHERE p.id = ?
	`, postID)

	if error != nil {
		return nil, error
	}

	var posts []models.Post

	for rows.Next() {
		var post models.Post

		if error := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorUsername); error != nil {
			return nil, error
		}

		posts = append(posts, post)
	}
	
	return posts, error
}