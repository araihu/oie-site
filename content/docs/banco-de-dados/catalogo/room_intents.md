---
title: "room_intents"
description: "Definição SQL, colunas e relações de room_intents no banco de dados do Oiê."
language: pt-BR
---

# `room_intents`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE room_intents (
 conversation_id TEXT PRIMARY KEY REFERENCES conversations(id),
 account_id TEXT NOT NULL REFERENCES accounts(id),
 id TEXT NOT NULL UNIQUE,
 revision_id TEXT NOT NULL,
 nickname TEXT NOT NULL,
 active INTEGER NOT NULL CHECK(active IN (0,1)),
 last_error_code TEXT NOT NULL,
 observation_fingerprint TEXT NOT NULL DEFAULT '',
 updated_at TEXT NOT NULL
, creation_name TEXT)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Não | [accounts](accounts.md) |
| active | INTEGER |  | Não |  |
| conversation_id | TEXT |  | Sim | [conversations](conversations.md) |
| creation_name | TEXT |  | Sim |  |
| id | TEXT |  | Não |  |
| last_error_code | TEXT |  | Não |  |
| nickname | TEXT |  | Não |  |
| observation_fingerprint | TEXT | '' | Não |  |
| revision_id | TEXT |  | Não |  |
| updated_at | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(active IN (0,1)) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| conversation_id | PRIMARY KEY | PRIMARY KEY (conversation_id) |
| sqlite_autoindex_room_intents_1 | PRIMARY KEY | PRIMARY KEY (conversation_id) |
| sqlite_autoindex_room_intents_2 | UNIQUE | UNIQUE (id) |


## Índices

- `sqlite_autoindex_room_intents_1`

  ```sql
  PRIMARY KEY (conversation_id)
  ```

- `sqlite_autoindex_room_intents_2`

  ```sql
  UNIQUE (id)
  ```
