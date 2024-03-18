package postgres

import "fmt"

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
	Image string `json:"image"`
}

type Storage interface {
	GetPosts() ([]*Post, error)
	GetPostByID(id int) (*Post, error)
	CreatePost(post *Post) (*Post, error)
	UpdatePost(id int, post *Post) (*Post, error)
	DeletePost(id int) error
}

func (s *PostgresStore) GetPosts() ([]*Post, error) {
	const op = "storage.postgres.GetPosts"
	posts := []*Post{}

	rows, err := s.db.Query("select * from post;")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var p Post

		if err := rows.Scan(&p.ID, &p.Title, &p.Text, &p.Image); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		posts = append(posts, &p)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return posts, nil
}

func (s *PostgresStore) GetPostByID(id int) (*Post, error) {
	const op = "storage.postgres.GetPostByID"

	p := &Post{}
	row := s.db.QueryRow("select * from post where id = $1;", id)

	if err := row.Scan(&p.ID, &p.Title, &p.Text, &p.Image); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return p, nil
}

func (s *PostgresStore) CreatePost(post *Post) (*Post, error) {
	const op = "storage.postgres.CreatePost"
	insertQuery := `insert into post (title, text, image) values ($1, $2, $3) returning id;`

	var id int
	err := s.db.QueryRow(insertQuery, post.Title, post.Text, post.Image).Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	post.ID = id

	return post, nil
}

func (s *PostgresStore) UpdatePost(id int, post *Post) (*Post, error) {
	const op = "storage.postgres.UpdatePost"

	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM post WHERE id = $1", id).Scan(&count)
	if err != nil {
		return nil, fmt.Errorf("failed to check post existence: %w", err)
	}
	if count == 0 {
		return nil, fmt.Errorf("post with id %d does not exist", id)
	}

	updateQuery := `update post set title = $1, text = $2, image = $3 where id = $4;`

	_, err = s.db.Exec(updateQuery, post.Title, post.Text, post.Image, id)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	post.ID = id

	return post, nil
}

func (s *PostgresStore) DeletePost(id int) error {
	const op = "storage.postgres.DeletePost"
	deleteQuery := "delete from post where id = $1"

	_, err := s.db.Exec(deleteQuery, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
