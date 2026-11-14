---
title: "conversation_read_states"
description: "Definição SQL, colunas e relações de conversation_read_states no banco de dados do Oiê."
language: pt-BR
---

# `conversation_read_states`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE conversation_read_states (
 conversation_id TEXT PRIMARY KEY REFERENCES conversations(id),
 revision_id TEXT NOT NULL REFERENCES conversation_read_revisions(id), updated_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| conversation_id | TEXT |  | Sim | [conversations](conversations.md) |
| revision_id | TEXT |  | Não | [conversation_read_revisions](conversation_read_revisions.md) |
| updated_at | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (revision_id) REFERENCES conversation_read_revisions (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| conversation_id | PRIMARY KEY | PRIMARY KEY (conversation_id) |
| sqlite_autoindex_conversation_read_states_1 | PRIMARY KEY | PRIMARY KEY (conversation_id) |


## Índices

- `sqlite_autoindex_conversation_read_states_1`

  ```sql
  PRIMARY KEY (conversation_id)
  ```
