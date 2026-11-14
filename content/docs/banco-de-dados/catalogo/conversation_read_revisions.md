---
title: "conversation_read_revisions"
description: "Definição SQL, colunas e relações de conversation_read_revisions no banco de dados do Oiê."
language: pt-BR
---

# `conversation_read_revisions`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE conversation_read_revisions (
 id TEXT PRIMARY KEY, conversation_id TEXT NOT NULL REFERENCES conversations(id),
 parents TEXT NOT NULL, payload TEXT NOT NULL, message_ids TEXT NOT NULL, created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| conversation_id | TEXT |  | Não |  | [conversations](conversations.md) |
| created_at | TEXT |  | Não |  |  |
| id | TEXT |  | Sim | [conversation_read_states](conversation_read_states.md) [message_reads](message_reads.md) |  |
| message_ids | TEXT |  | Não |  |  |
| parents | TEXT |  | Não |  |  |
| payload | TEXT |  | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_conversation_read_revisions_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `sqlite_autoindex_conversation_read_revisions_1`

  ```sql
  PRIMARY KEY (id)
  ```
