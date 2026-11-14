---
title: "archive_checkpoints"
description: "Definição SQL, colunas e relações de archive_checkpoints no banco de dados do Oiê."
language: pt-BR
---

# `archive_checkpoints`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE archive_checkpoints (
 account_id TEXT NOT NULL REFERENCES accounts(id), scope TEXT NOT NULL, after_id TEXT NOT NULL,
 PRIMARY KEY(account_id,scope)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Não | [accounts](accounts.md) |
| after_id | TEXT |  | Não |  |
| scope | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| account_id | PRIMARY KEY | PRIMARY KEY (account_id) |
| scope | PRIMARY KEY | PRIMARY KEY (scope) |
| sqlite_autoindex_archive_checkpoints_1 | PRIMARY KEY | PRIMARY KEY (account_id, scope) |


## Índices

- `sqlite_autoindex_archive_checkpoints_1`

  ```sql
  PRIMARY KEY (account_id, scope)
  ```
