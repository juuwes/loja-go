package models

import (
	"github.com/loja/db"
)

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaProdutos() []Produto {
	db := db.ConectaBD()
	selectProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer db.Close()
	return produtos

}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBD()

	insereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaBD()

	deletar, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletar.Exec(id)
	defer db.Close()
}

func EditaProduto(id string) Produto {
	db := db.ConectaBD()

	produtoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		produtoAtualizar.Id = id
		produtoAtualizar.Nome = nome
		produtoAtualizar.Descricao = descricao
		produtoAtualizar.Preco = preco
		produtoAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaBD()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
