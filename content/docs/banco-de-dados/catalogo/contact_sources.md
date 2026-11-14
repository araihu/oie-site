---
title: "contact_sources"
description: "Definição SQL, colunas e relações de contact_sources no banco de dados do Oiê."
language: pt-BR
---

# `contact_sources`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE contact_sources (
 account_id TEXT PRIMARY KEY REFERENCES accounts(id),
 state TEXT NOT NULL CHECK(state IN ('known','unavailable')),
 checked_at TEXT NOT NULL,
 updated_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Sim | [accounts](accounts.md) |
| checked_at | TEXT |  | Não |  |
| state | TEXT |  | Não |  |
| updated_at | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(state IN ('known','unavailable')) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| account_id | PRIMARY KEY | PRIMARY KEY (account_id) |
| sqlite_autoindex_contact_sources_1 | PRIMARY KEY | PRIMARY KEY (account_id) |


## Índices

- `sqlite_autoindex_contact_sources_1`

  ```sql
  PRIMARY KEY (account_id)
  ```
