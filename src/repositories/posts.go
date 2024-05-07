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

func (repository postsRepository) GetPostByID(postID uint64) (models.Post, error) {

	rows, error := repository.db.Query(`
		SELECT p.*, u.username from posts p JOIN users u ON p.author_id = u.id
		WHERE p.id = ?
	`, postID)

	if error != nil {
		return models.Post{}, error
	}

	
	if rows.Next() {
		var post models.Post

		if error := rows.Scan(&post.ID, &post.Title, &post.Content, &post.AuthorID, &post.Likes, &post.CreatedAt, &post.AuthorUsername); error != nil {
			return models.Post{}, error
		}

		return post, error
	}

	return models.Post{}, error
}

func (repository postsRepository) GetPosts(userID uint64) ([]models.Post, error) {

	rows, error := repository.db.Query(`
	SELECT DISTINCT p.*, u.username 
	FROM posts p
	INNER JOIN users u ON u.id = p.author_id 
	INNER JOIN followers f ON p.author_id = f.user_id
	WHERE u.id = ?
		OR f.follower_id = ?
	ORDER BY 1 DESC
	`, userID, userID)

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

	return posts, nil
}

func (repository postsRepository) UpdatePost(postID uint64, post models.Post) error {

	statement, error := repository.db.Prepare("UPDATE posts set title = ?, content = ? where id = ?")


	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error := statement.Exec(post.Title, post.Content, postID); error != nil {
		return error
	}

	return nil
}

func (repository postsRepository) DeletePost(postID uint64) error {


	statement, error := repository.db.Prepare("DELETE FROM posts WHERE id = ?")

	if error != nil {
		return error
	}

	defer statement.Close()

	if _, error := statement.Exec(postID); error != nil {
		return error
	}

	return nil
}