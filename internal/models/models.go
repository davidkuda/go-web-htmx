package models

import "database/sql"

type Models struct {
	Pages PageModel
}

func New(db *sql.DB) Models {
	return Models{
		Pages: PageModel{DB: db},
	}
}
