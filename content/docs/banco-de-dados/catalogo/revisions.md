---
title: "revisions"
description: "Definição SQL, colunas e relações de revisions no banco de dados do Oiê."
language: pt-BR
---

# `revisions`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE revisions (
    id TEXT PRIMARY KEY,
    entity_id TEXT NOT NULL REFERENCES accounts(id),
    parents TEXT NOT NULL,
    payload TEXT NOT NULL,
    created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| created_at | TEXT |  | Não |  |  |
| entity_id | TEXT |  | Não |  | [accounts](accounts.md) |
| id | TEXT |  | Sim | [accounts](accounts.md) |  |
| parents | TEXT |  | Não |  |  |
| payload | TEXT |  | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (entity_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_revisions_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `sqlite_autoindex_revisions_1`

  ```sql
  PRIMARY KEY (id)
  ```
