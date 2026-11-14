---
title: "preview_revisions"
description: "Definição SQL, colunas e relações de preview_revisions no banco de dados do Oiê."
language: pt-BR
---

# `preview_revisions`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE preview_revisions (
 id TEXT PRIMARY KEY,entity_id TEXT NOT NULL REFERENCES message_previews(id),
 parents TEXT NOT NULL,payload TEXT NOT NULL,created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| created_at | TEXT |  | Não |  |
| entity_id | TEXT |  | Não | [message_previews](message_previews.md) |
| id | TEXT |  | Sim |  |
| parents | TEXT |  | Não |  |
| payload | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (entity_id) REFERENCES message_previews (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_preview_revisions_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `sqlite_autoindex_preview_revisions_1`

  ```sql
  PRIMARY KEY (id)
  ```
