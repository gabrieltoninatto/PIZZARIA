# API de Pizzaria

Esta é uma API RESTful desenvolvida em Go (Golang) utilizando o framework Gin para gerenciar o cardápio de uma pizzaria e suas avaliações.

## Funcionalidades

- **Gerenciamento de Pizzas**: Criar, listar, atualizar, buscar por ID e remover pizzas.
- **Avaliações**: Adicionar avaliações (nota e comentário) para as pizzas.
- **Persistência**: Os dados são armazenados localmente em um arquivo JSON (`dados/pizza.json`).

## Tecnologias Utilizadas

- Go
- Gin Web Framework

## Instalação e Execução

1. Certifique-se de ter o Go instalado em sua máquina.
2. Clone este repositório.
3. Instale as dependências do projeto:
   ```bash
   go mod tidy
   ```
4. Execute a aplicação:
   ```bash
   go run cmd/main.go
   ```

O servidor iniciará (padrão na porta 8080).

## Documentação da API

### Rotas de Pizzas

| Método | Endpoint       | Descrição                   |
|--------|----------------|-----------------------------|
| GET    | `/pizzas`      | Lista todas as pizzas       |
| POST   | `/pizzas`      | Cria uma nova pizza         |
| GET    | `/pizzas/:id`  | Busca uma pizza pelo ID     |
| PUT    | `/pizzas/:id`  | Atualiza uma pizza existente|
| DELETE | `/pizzas/:id`  | Remove uma pizza            |

### Rotas de Avaliações

| Método | Endpoint                 | Descrição                          |
|--------|--------------------------|------------------------------------|
| POST   | `/pizzas/:id/reviews`    | Adiciona uma avaliação a uma pizza |

### Exemplos de JSON (Body)

**Criar/Atualizar Pizza:**
```json
{
  "Nome": "Calabresa",
  "Preco": 45.90
}
```

**Adicionar Review:**
```json
{
  "rating": 5,
  "comment": "Muito saborosa!"
}
```