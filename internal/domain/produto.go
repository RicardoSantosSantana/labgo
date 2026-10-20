package domain

import (
	"errors"
	"math"
	"strings"
)

type Produto struct {
	ID            int64  `json:"id"`
	Nome          string `json:"nome"`
	PrecoCentavos int64  `json:"preco_centavos"`
	Estoque       int    `json:"estoque"`
}

var ErrNomeObrigatorio = errors.New("nome do produto é obrigatório")
var ErrPrecoInvalido = errors.New("preço do produto deve ser maior que zero")
var ErrEstoqueNegativo = errors.New("estoque do produto não pode ser negativo")
var ErrProdutoNil = errors.New("produto não pode ser nil")
var ErrQuantidadeInvalida = errors.New("quantidade deve ser maior que zero")
var ErrEstoqueOverflow = errors.New("quantidade excede a capacidade máxima do estoque")
var ErrEstoqueInsuficiente = errors.New("estoque insuficiente")

func (p *Produto) RemoverEstoque(quantidade int) error {
	if quantidade <= 0 {
		return ErrQuantidadeInvalida
	}

	err := p.Validar()

	if err != nil {
		return err
	}

	if p.Estoque < quantidade {
		return ErrEstoqueInsuficiente
	}

	p.Estoque -= quantidade

	return nil
}

func (p *Produto) Validar() error {

	if p == nil {
		return ErrProdutoNil
	}

	if strings.TrimSpace(p.Nome) == "" {
		return ErrNomeObrigatorio
	}

	if p.PrecoCentavos <= 0 {
		return ErrPrecoInvalido
	}

	if p.Estoque < 0 {
		return ErrEstoqueNegativo
	}

	return nil
}
func (p *Produto) AdicionarEstoque(quantidade int) error {

	if quantidade <= 0 {
		return ErrQuantidadeInvalida
	}

	err := p.Validar()

	if err != nil {
		return err
	}

	if p.Estoque > (math.MaxInt - quantidade) {
		return ErrEstoqueOverflow
	}

	p.Estoque += quantidade

	return nil
}
