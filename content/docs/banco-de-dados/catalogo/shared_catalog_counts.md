---
title: "shared_catalog_counts"
description: "Definição SQL, colunas e relações de shared_catalog_counts no banco de dados do Oiê."
language: pt-BR
---

# `shared_catalog_counts`

[Catálogo de tabelas e views](index.md)

Disponível a partir da migração `000027`, posterior à `v0.0.1-rc.9`.

## Definição SQL

```sql
CREATE TABLE shared_catalog_counts (
 conversation_id TEXT NOT NULL,
 parser_version INTEGER NOT NULL,
 messages INTEGER NOT NULL,
 PRIMARY KEY(conversation_id,parser_version)
) WITHOUT ROWID
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| conversation_id | TEXT |  | Não |
| messages | INTEGER |  | Não |
| parser_version | INTEGER |  | Não |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| conversation_id | PRIMARY KEY | PRIMARY KEY (conversation_id) |
| parser_version | PRIMARY KEY | PRIMARY KEY (parser_version) |
| sqlite_autoindex_shared_catalog_counts_1 | PRIMARY KEY | PRIMARY KEY (conversation_id, parser_version) |


## Índices

- `sqlite_autoindex_shared_catalog_counts_1`

  ```sql
  PRIMARY KEY (conversation_id, parser_version)
  ```
