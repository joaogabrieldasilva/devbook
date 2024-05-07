package repositories

import "database/sql"

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