---
title: "message_stars"
description: "Definição SQL, colunas e relações de message_stars no banco de dados do Oiê."
language: pt-BR
---

# `message_stars`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_stars (
 message_id TEXT PRIMARY KEY REFERENCES messages(id),
 revision_id TEXT NOT NULL UNIQUE REFERENCES message_star_revisions(id) DEFERRABLE INITIALLY DEFERRED,
 starred INTEGER NOT NULL CHECK(starred IN (0,1)),
 payload TEXT NOT NULL,
 updated_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| message_id | TEXT |  | Sim | [messages](messages.md) |
| payload | TEXT |  | Não |  |
| revision_id | TEXT |  | Não | [message_star_revisions](message_star_revisions.md) |
| starred | INTEGER |  | Não |  |
| updated_at | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(starred IN (0,1)) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (revision_id) REFERENCES message_star_revisions (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| sqlite_autoindex_message_stars_1 | PRIMARY KEY | PRIMARY KEY (message_id) |
| sqlite_autoindex_message_stars_2 | UNIQUE | UNIQUE (revision_id) |


## Índices

- `message_stars_selected`

  ```sql
  CREATE INDEX message_stars_selected ON message_stars(starred,message_id)
  ```

- `sqlite_autoindex_message_stars_1`

  ```sql
  PRIMARY KEY (message_id)
  ```

- `sqlite_autoindex_message_stars_2`

  ```sql
  UNIQUE (revision_id)
  ```
