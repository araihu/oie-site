---
title: "room_reaction_privacy"
description: "Definição SQL, colunas e relações de room_reaction_privacy no banco de dados do Oiê."
language: pt-BR
---

# `room_reaction_privacy`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE room_reaction_privacy (
 account_id TEXT NOT NULL REFERENCES accounts(id),
 address TEXT NOT NULL,
 public INTEGER NOT NULL CHECK(public IN (0,1)), identity_mode TEXT NOT NULL DEFAULT '', epoch INTEGER NOT NULL DEFAULT 0,
 PRIMARY KEY(account_id,address)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Não | [accounts](accounts.md) |
| address | TEXT |  | Não |  |
| epoch | INTEGER | 0 | Não |  |
| identity_mode | TEXT | '' | Não |  |
| public | INTEGER |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(public IN (0,1)) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| account_id | PRIMARY KEY | PRIMARY KEY (account_id) |
| address | PRIMARY KEY | PRIMARY KEY (address) |
| sqlite_autoindex_room_reaction_privacy_1 | PRIMARY KEY | PRIMARY KEY (account_id, address) |


## Índices

- `sqlite_autoindex_room_reaction_privacy_1`

  ```sql
  PRIMARY KEY (account_id, address)
  ```
