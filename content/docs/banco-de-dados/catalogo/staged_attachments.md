---
title: "staged_attachments"
description: "Definição SQL, colunas e relações de staged_attachments no banco de dados do Oiê."
language: pt-BR
---

# `staged_attachments`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE staged_attachments (
    id TEXT PRIMARY KEY,
    conversation_id TEXT NOT NULL REFERENCES conversations(id),
    file_name TEXT NOT NULL,
    media_type TEXT NOT NULL,
    size INTEGER NOT NULL CHECK(size >= 0),
    sha256 TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| conversation_id | TEXT |  | Não | [conversations](conversations.md) |
| file_name | TEXT |  | Não |  |
| id | TEXT |  | Sim |  |
| media_type | TEXT |  | Não |  |
| sha256 | TEXT |  | Não |  |
| size | INTEGER |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(size >= 0) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_staged_attachments_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `sqlite_autoindex_staged_attachments_1`

  ```sql
  PRIMARY KEY (id)
  ```
