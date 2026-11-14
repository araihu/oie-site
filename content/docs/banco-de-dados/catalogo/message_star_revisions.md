---
title: "message_star_revisions"
description: "Definição SQL, colunas e relações de message_star_revisions no banco de dados do Oiê."
language: pt-BR
---

# `message_star_revisions`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_star_revisions (
 id TEXT PRIMARY KEY,
 message_id TEXT NOT NULL REFERENCES messages(id),
 parents TEXT NOT NULL,
 payload TEXT NOT NULL,
 created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| created_at | TEXT |  | Não |  |  |
| id | TEXT |  | Sim | [message_stars](message_stars.md) |  |
| message_id | TEXT |  | Não |  | [messages](messages.md) |
| parents | TEXT |  | Não |  |  |
| payload | TEXT |  | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_message_star_revisions_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `sqlite_autoindex_message_star_revisions_1`

  ```sql
  PRIMARY KEY (id)
  ```
