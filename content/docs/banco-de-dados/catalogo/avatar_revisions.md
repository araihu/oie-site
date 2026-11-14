---
title: "avatar_revisions"
description: "Definição SQL, colunas e relações de avatar_revisions no banco de dados do Oiê."
language: pt-BR
---

# `avatar_revisions`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE avatar_revisions (
 id TEXT PRIMARY KEY,entity_id TEXT NOT NULL REFERENCES avatars(id),parents TEXT NOT NULL,
 payload TEXT NOT NULL,created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| created_at | TEXT |  | Não |  |
| entity_id | TEXT |  | Não | [avatars](avatars.md) |
| id | TEXT |  | Sim |  |
| parents | TEXT |  | Não |  |
| payload | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (entity_id) REFERENCES avatars (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_avatar_revisions_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `sqlite_autoindex_avatar_revisions_1`

  ```sql
  PRIMARY KEY (id)
  ```
