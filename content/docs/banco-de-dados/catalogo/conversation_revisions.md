---
title: "conversation_revisions"
description: "Definição SQL, colunas e relações de conversation_revisions no banco de dados do Oiê."
language: pt-BR
---

# `conversation_revisions`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE conversation_revisions (
 id TEXT PRIMARY KEY, entity_id TEXT NOT NULL REFERENCES conversations(id) DEFERRABLE INITIALLY DEFERRED,
 parents TEXT NOT NULL, payload TEXT NOT NULL, created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| created_at | TEXT |  | Não |  |
| entity_id | TEXT |  | Não | [conversations](conversations.md) |
| id | TEXT |  | Sim |  |
| parents | TEXT |  | Não |  |
| payload | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (entity_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_conversation_revisions_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `sqlite_autoindex_conversation_revisions_1`

  ```sql
  PRIMARY KEY (id)
  ```
