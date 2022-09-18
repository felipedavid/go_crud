package models

import "database/sql"

type Gopher struct {
	id    int64
	name  string
	color string
}

type GopherModel struct {
	DB *sql.DB
}

func (m *GopherModel) Get(id int64) (*Gopher, error) {
	stmt := `SELECT * FROM gopher WHERE id = $1`

	row := m.DB.QueryRow(stmt, id)

	g := &Gopher{}

	if err := row.Scan(&g.id, &g.name, &g.color); err != nil {
		return nil, err
	}

	return g, nil
}

func (m *GopherModel) Insert(name, color string) (int64, error) {
	stmt := `INSERT INTO `
}

func (m *GopherModel) Update(id int64, name, color string) error {
	return nil
}

func (m *GopherModel) Delete(id int64) error {
	return nil
}
