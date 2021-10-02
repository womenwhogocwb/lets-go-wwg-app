# 💎 Pedra 📜 Papel ✂️ Tesoura

## Objetivo

Disponibilizar, por meio uma API HTTP Rest, um jogo de "Pedra, Papel, Tesoura".

## Rotas

### Jogar um jogo de "Pedra, Papel, Tesoura"

#### Request

- Path: `/games`
- Method: `POST`
- Content-Type: `application/json`
- Body:
  ```json
  {
    "name": "Leslie Knope",
    "move": "tesoura"
  }
  ```

#### Response

##### Success

- Status code: `200`
- Content-Type: `application/json`
- Body:
  ```json
  {
    "id": "270e2928-74b6-4d2e-b333-15f77ddc2b61",
    "player_name": "Amy Santiago",
    "player_move": "pedra",
    "house_move": "tesoura",
    "result": "victory",
    "created_at": "2021-07-31T16:38:20Z"
  }
  ```

##### Failure

- Status code: `400`, `500`
- Content-Type: `application/json`
- Body (example):
  ```json
  {
    "reason": "invalid move"
  }
  ```

### Listar todos os jogos

#### Request

- Path: `/games`
- Method: `GET`

#### Response

##### Success

- Status code: `200`
- Content-Type: `application/json`
- Body:
  ```json
  [
    {
      "id": "647e1f05-5ceb-45ab-8f4a-b67510f6deb3",
      "player_name": "Amy Santiago",
      "player_move": "pedra",
      "house_move": "tesoura",
      "result": "victory",
      "created_at": "2021-08-06T10:57:43Z"
    },
    {
      "id": "8359fa5c-6901-47b1-93ff-6bcbf465e5aa",
      "player_name": "Rosa Diaz",
      "player_move": "tesoura",
      "house_move": "pedra",
      "result": "defeat",
      "created_at": "2021-08-06T10:57:51Z"
    }
  ]
  ```

##### Failure

- Status code: `500`
- Content-Type: `application/json`
- Body (example):
  ```json
  {
    "reason": "internal server error"
  }
  ```

### Listar um jogo pelo identificador

#### Request

- Path: `/games/{id}`
- Method: `GET`

#### Response

##### Success

- Status code: `200`
- Content-Type: `application/json`
- Body:
  ```json
  {
    "id": "647e1f05-5ceb-45ab-8f4a-b67510f6deb3",
    "player_name": "Amy Santiago",
    "player_move": "pedra",
    "house_move": "tesoura",
    "result": "victory",
    "created_at": "2021-08-06T10:57:43Z"
  }
  ```

##### Failure

- Status code: `500`
- Content-Type: `application/json`
- Body (example):
  ```json
  {
    "reason": "internal server error"
  }
  ```

## Exercício extra do Let's Go!

1. Faça um _fork_ desse repositório. Se não souber como, [aqui tem uma explicação](https://docs.github.com/pt/github/getting-started-with-github/quickstart/fork-a-repo#prerequisties)
2. Faça um clone local do seu _fork_. Se não souber como, [aqui tem uma explicação](https://docs.github.com/pt/github/getting-started-with-github/quickstart/fork-a-repo#configuring-git-to-sync-your-fork-with-the-original-repository)
3. Com seu _fork_ disponível localmente, realize as seguintes adições ao projeto:
   - Adicionar mais um endpoint à aplicação. Exemplos que podem ser explorados: consultar um jogo usando seu identificador, listar todos os jogos de uma mesma pessoa
   - Adicionar testes à camada HTTP
