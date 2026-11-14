---
title: "message_reads"
description: "Definição SQL, colunas e relações de message_reads no banco de dados do Oiê."
language: pt-BR
---

# `message_reads`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_reads (
 message_id TEXT PRIMARY KEY REFERENCES messages(id),
 revision_id TEXT NOT NULL REFERENCES conversation_read_revisions(id) DEFERRABLE INITIALLY DEFERRED
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| message_id | TEXT |  | Sim | [messages](messages.md) |
| revision_id | TEXT |  | Não | [conversation_read_revisions](conversation_read_revisions.md) |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (revision_id) REFERENCES conversation_read_revisions (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| sqlite_autoindex_message_reads_1 | PRIMARY KEY | PRIMARY KEY (message_id) |


## Índices

- `sqlite_autoindex_message_reads_1`

  ```sql
  PRIMARY KEY (message_id)
  ```
