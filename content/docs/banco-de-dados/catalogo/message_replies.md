---
title: "message_replies"
description: "Definição SQL, colunas e relações de message_replies no banco de dados do Oiê."
language: pt-BR
---

# `message_replies`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_replies (
 message_id TEXT NOT NULL REFERENCES messages(id), body_hash TEXT NOT NULL,
 payload TEXT NOT NULL, PRIMARY KEY(message_id,body_hash)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| body_hash | TEXT |  | Não |  |
| message_id | TEXT |  | Não | [messages](messages.md) |
| payload | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| body_hash | PRIMARY KEY | PRIMARY KEY (body_hash) |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| sqlite_autoindex_message_replies_1 | PRIMARY KEY | PRIMARY KEY (message_id, body_hash) |


## Índices

- `sqlite_autoindex_message_replies_1`

  ```sql
  PRIMARY KEY (message_id, body_hash)
  ```
