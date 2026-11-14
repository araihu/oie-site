---
title: "message_content_heads"
description: "Definição SQL, colunas e relações de message_content_heads no banco de dados do Oiê."
language: pt-BR
---

# `message_content_heads`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_content_heads (
 canonical_id TEXT PRIMARY KEY REFERENCES messages(id), source_id TEXT NOT NULL REFERENCES messages(id),
 revision_id TEXT NOT NULL REFERENCES message_revisions(id), payload TEXT NOT NULL, edited_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| canonical_id | TEXT |  | Sim | [messages](messages.md) |
| edited_at | TEXT |  | Não |  |
| payload | TEXT |  | Não |  |
| revision_id | TEXT |  | Não | [message_revisions](message_revisions.md) |
| source_id | TEXT |  | Não | [messages](messages.md) |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (revision_id) REFERENCES message_revisions (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (source_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 2) | FOREIGN KEY | FOREIGN KEY (canonical_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| canonical_id | PRIMARY KEY | PRIMARY KEY (canonical_id) |
| sqlite_autoindex_message_content_heads_1 | PRIMARY KEY | PRIMARY KEY (canonical_id) |


## Índices

- `sqlite_autoindex_message_content_heads_1`

  ```sql
  PRIMARY KEY (canonical_id)
  ```
