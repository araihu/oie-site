---
title: "media_files"
description: "Definição SQL, colunas e relações de media_files no banco de dados do Oiê."
language: pt-BR
---

# `media_files`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE media_files (
 hash TEXT PRIMARY KEY,mime_type TEXT NOT NULL,size INTEGER NOT NULL,created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| created_at | TEXT |  | Não |
| hash | TEXT |  | Sim |
| mime_type | TEXT |  | Não |
| size | INTEGER |  | Não |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| hash | PRIMARY KEY | PRIMARY KEY (hash) |
| sqlite_autoindex_media_files_1 | PRIMARY KEY | PRIMARY KEY (hash) |


## Índices

- `sqlite_autoindex_media_files_1`

  ```sql
  PRIMARY KEY (hash)
  ```
