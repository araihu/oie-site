---
title: "message_revisions"
description: "Definição SQL, colunas e relações de message_revisions no banco de dados do Oiê."
language: pt-BR
---

# `message_revisions`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_revisions (
 id TEXT PRIMARY KEY, entity_id TEXT NOT NULL REFERENCES messages(id) DEFERRABLE INITIALLY DEFERRED,
 parents TEXT NOT NULL, payload TEXT NOT NULL, created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| created_at | TEXT |  | Não |  |  |
| entity_id | TEXT |  | Não |  | [messages](messages.md) |
| id | TEXT |  | Sim | [message_content_heads](message_content_heads.md) [message_correction_facts](message_correction_facts.md) |  |
| parents | TEXT |  | Não |  |  |
| payload | TEXT |  | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (entity_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_message_revisions_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `message_revisions_entity`

  ```sql
  CREATE INDEX message_revisions_entity ON message_revisions(entity_id)
  ```

- `sqlite_autoindex_message_revisions_1`

  ```sql
  PRIMARY KEY (id)
  ```
