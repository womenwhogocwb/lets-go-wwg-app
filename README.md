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
- Body:
  ```json
  {
    "reason": "invalid move"
  }
  ```

### Listar todos os jogos

#### Request

[comment]: <> (preencher)

#### Response

[comment]: <> (preencher)
