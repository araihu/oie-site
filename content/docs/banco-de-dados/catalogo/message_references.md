---
title: "message_references"
description: "Definição SQL, colunas e relações de message_references no banco de dados do Oiê."
language: pt-BR
---

# `message_references`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_references (
 message_id TEXT NOT NULL REFERENCES messages(id),
 scope TEXT NOT NULL, remote_id TEXT NOT NULL, author TEXT NOT NULL,
 source_hash TEXT NOT NULL,
 PRIMARY KEY(message_id,scope,remote_id,author,source_hash)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| author | TEXT |  | Não |  |
| message_id | TEXT |  | Não | [messages](messages.md) |
| remote_id | TEXT |  | Não |  |
| scope | TEXT |  | Não |  |
| source_hash | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| author | PRIMARY KEY | PRIMARY KEY (author) |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| remote_id | PRIMARY KEY | PRIMARY KEY (remote_id) |
| scope | PRIMARY KEY | PRIMARY KEY (scope) |
| source_hash | PRIMARY KEY | PRIMARY KEY (source_hash) |
| sqlite_autoindex_message_references_1 | PRIMARY KEY | PRIMARY KEY (message_id, scope, remote_id, author, source_hash) |


## Índices

- `message_references_lookup`

  ```sql
  CREATE INDEX message_references_lookup ON message_references(remote_id,scope,author,message_id)
  ```

- `sqlite_autoindex_message_references_1`

  ```sql
  PRIMARY KEY (message_id, scope, remote_id, author, source_hash)
  ```
