package domain

import (
	"errors"
	"math"
	"testing"
)

type testeProduto struct {
	name         string
	p            *Produto
	deveFalhar   bool
	erroEsperado error
}

type testeMovimentacaoEstoque struct {
	name            string
	p               *Produto
	quantidade      int
	estoqueEsperado int
	erroEsperado    error
}

func TestProduto_AdicionarEstoqueOverflow(t *testing.T) {
	p := &Produto{
		Nome:          "Notebook",
		PrecoCentavos: 350000,
		Estoque:       math.MaxInt - 5,
	}
	err := p.AdicionarEstoque(10)

	if !errors.Is(err, ErrEstoqueOverflow) {
		t.Errorf("Erro esperado: %v, obtido: %v", ErrEstoqueOverflow, err)
	}
	if p.Estoque != math.MaxInt-5 {
		t.Errorf(
			"Estoque esperado: %d, obtido: %d",
			math.MaxInt-5,
			p.Estoque,
		)
	}
}

func TestProduto_AdicionarEstoqueQuantidadeNegativa(t *testing.T) {
	p := &Produto{
		Nome:          "Notebook",
		PrecoCentavos: 350000,
		Estoque:       10,
	}
	err := p.AdicionarEstoque(-5)

	if !errors.Is(err, ErrQuantidadeInvalida) {
		t.Errorf("Erro esperado: %v, obtido: %v", ErrQuantidadeInvalida, err)
	}
	if p.Estoque != 10 {
		t.Errorf("Estoque esperado: 10, obtido: %d", p.Estoque)
	}
}
func TestProduto_AdicionarEstoque(t *testing.T) {
	p := &Produto{
		Nome:          "Notebook",
		PrecoCentavos: 350000,
		Estoque:       5,
	}
	err := p.AdicionarEstoque(5)

	if err != nil {
		t.Errorf("Erro inesperado: %v", err)
	}

	if p.Estoque != 10 {
		t.Errorf("Estoque esperado: 10, obtido: %d", p.Estoque)
	}
}
func TestProduto_RemoverEstoque(t *testing.T) {

	var ValidarRemoverEstoque3 = testeMovimentacaoEstoque{
		name: "Remover estoque 3",
		p: &Produto{
			Nome:          "Notebook",
			PrecoCentavos: 350000,
			Estoque:       10,
		},
		quantidade:      3,
		estoqueEsperado: 7,
		erroEsperado:    nil,
	}

	var ValidarRemoverEstoque10 = testeMovimentacaoEstoque{
		name: "Remover estoque 10",
		p: &Produto{
			Nome:          "Notebook",
			PrecoCentavos: 350000,
			Estoque:       10,
		},
		quantidade:      10,
		estoqueEsperado: 0,
		erroEsperado:    nil,
	}

	var ValidarRemoverEstoque15 = testeMovimentacaoEstoque{
		name: "Remover estoque 15",
		p: &Produto{
			Nome:          "Notebook",
			PrecoCentavos: 350000,
			Estoque:       10,
		},
		quantidade:      15,
		estoqueEsperado: 10,
		erroEsperado:    ErrEstoqueInsuficiente,
	}

	var ValidarRemoverEstoqueNegativo = testeMovimentacaoEstoque{
		name: "Remover estoque -5",
		p: &Produto{
			Nome:          "Notebook",
			PrecoCentavos: 350000,
			Estoque:       10,
		},
		quantidade:      -5,
		estoqueEsperado: 10,
		erroEsperado:    ErrQuantidadeInvalida,
	}
	tests := []testeMovimentacaoEstoque{
		ValidarRemoverEstoque3,
		ValidarRemoverEstoque10,
		ValidarRemoverEstoque15,
		ValidarRemoverEstoqueNegativo,
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.p.RemoverEstoque(test.quantidade)

			if !errors.Is(err, test.erroEsperado) {
				t.Fatalf("expected error %v for %s, got: %v", test.erroEsperado, test.name, err)
			}
			if test.p.Estoque != test.estoqueEsperado {
				t.Errorf(
					"estoque esperado: %d, obtido: %d",
					test.estoqueEsperado,
					test.p.Estoque,
				)
			}

		})
	}
}

func TestProduto_ValidarTabela(t *testing.T) {

	var ValidarProdutoValido = testeProduto{
		name: "Produto válido",
		p: &Produto{
			Nome:          "Notebook",
			PrecoCentavos: 350000,
			Estoque:       5,
		},
		deveFalhar:   false,
		erroEsperado: nil,
	}

	var ValidarProdutoNomeVazio = testeProduto{
		name: "Produto com nome Vazio",
		p: &Produto{
			Nome:          "",
			PrecoCentavos: 350000,
			Estoque:       5,
		},
		deveFalhar:   true,
		erroEsperado: ErrNomeObrigatorio,
	}
	var ValidarProdutoEstoqueZero = testeProduto{
		name: "Produto com estoque zero",
		p: &Produto{
			Nome:          "Teclado",
			PrecoCentavos: 15000,
			Estoque:       0,
		},
		deveFalhar:   false,
		erroEsperado: nil,
	}
	var ValidarProdutoEstoqueNegativo = testeProduto{
		name: "Produto com estoque negativo",
		p: &Produto{
			Nome:          "Mouse",
			PrecoCentavos: 15000,
			Estoque:       -1,
		},
		deveFalhar:   true,
		erroEsperado: ErrEstoqueNegativo,
	}
	var ValidarProdutoPrecoNegativo = testeProduto{
		name: "Produto com preco negativo",
		p: &Produto{
			Nome:          "Mouse",
			PrecoCentavos: -15000,
			Estoque:       1,
		},
		deveFalhar:   true,
		erroEsperado: ErrPrecoInvalido,
	}

	var ValidarProdutoPrecoZero = testeProduto{
		name: "Produto com preco zero",
		p: &Produto{
			Nome:          "Mouse",
			PrecoCentavos: 0,
			Estoque:       1,
		},
		deveFalhar:   true,
		erroEsperado: ErrPrecoInvalido,
	}

	var ValidarProdutoNil = testeProduto{
		name:         "Produto nil",
		p:            nil,
		deveFalhar:   true,
		erroEsperado: ErrProdutoNil,
	}

	tests := []testeProduto{
		ValidarProdutoValido,
		ValidarProdutoNomeVazio,
		ValidarProdutoEstoqueZero,
		ValidarProdutoEstoqueNegativo,
		ValidarProdutoPrecoNegativo,
		ValidarProdutoPrecoZero,
		ValidarProdutoNil,
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.p.Validar()
			if test.deveFalhar {
				if err == nil {
					t.Fatalf("expected error for %s, got:", test.name)
				}

				if !errors.Is(err, test.erroEsperado) {
					t.Fatalf("expected error %v for %s, got: %v", test.erroEsperado, test.name, err)
				}

			} else {
				if err != nil {
					t.Fatalf("expected no error for %s, got: %v", test.name, err)
				}
			}
		})
	}
}
