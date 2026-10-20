package repository

import (
	"context"

	"example.com/labgo/internal/domain"
)

type ProdutoRepository interface {
	Salvar(ctx context.Context, produto *domain.Produto) error
	BuscarPorID(ctx context.Context, id int64) (*domain.Produto, error)
}
