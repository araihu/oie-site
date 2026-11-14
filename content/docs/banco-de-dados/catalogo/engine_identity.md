---
title: "engine_identity"
description: "Definição SQL, colunas e relações de engine_identity no banco de dados do Oiê."
language: pt-BR
---

# `engine_identity`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE engine_identity (
    singleton INTEGER PRIMARY KEY CHECK(singleton = 1),
    id TEXT NOT NULL UNIQUE
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| id | TEXT |  | Não |
| singleton | INTEGER |  | Sim |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(singleton = 1) |
| singleton | PRIMARY KEY | PRIMARY KEY (singleton) |
| sqlite_autoindex_engine_identity_1 | UNIQUE | UNIQUE (id) |


## Índices

- `sqlite_autoindex_engine_identity_1`

  ```sql
  UNIQUE (id)
  ```
