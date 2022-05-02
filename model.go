package main

import (
	"database/sql"
	// "errors"
)

// CREATE TABLE menor_preco(
// 	uf VARCHAR,
// 	municipio VARCHAR,
// 	bairro	VARCHAR,
// 	produto VARCHAR,
// 	menor DECIMAL
// );


type product struct {
	Uf  string  `json:"uf"`
	City  string  `json:"city"`
	District  string  `json:"district"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
    return db.QueryRow("SELECT produto, menor FROM menor_preco WHERE uf=$1 AND municipio=$2 AND bairro=$3 AND produto LIKE $4",
    p.Uf, p.City, p.District, "%"+p.Name+"%").Scan(&p.Name, &p.Price)
}
