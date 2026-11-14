---
title: "history_status"
description: "Definição SQL, colunas e relações de history_status no banco de dados do Oiê."
language: pt-BR
---

# `history_status`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE history_status (
 account_id TEXT PRIMARY KEY REFERENCES accounts(id), payload TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Sim | [accounts](accounts.md) |
| payload | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| account_id | PRIMARY KEY | PRIMARY KEY (account_id) |
| sqlite_autoindex_history_status_1 | PRIMARY KEY | PRIMARY KEY (account_id) |


## Índices

- `sqlite_autoindex_history_status_1`

  ```sql
  PRIMARY KEY (account_id)
  ```
