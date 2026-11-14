---
title: "message_reaction_commands"
description: "Definição SQL, colunas e relações de message_reaction_commands no banco de dados do Oiê."
language: pt-BR
---

# `message_reaction_commands`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_reaction_commands (
 id TEXT PRIMARY KEY,
 revision_id TEXT NOT NULL,
 message_id TEXT NOT NULL REFERENCES messages(id),
 account_id TEXT NOT NULL REFERENCES accounts(id),
 conversation_id TEXT NOT NULL REFERENCES conversations(id),
 state TEXT NOT NULL CHECK(state IN ('queued','dispatching','submitted','unknown','failed')),
 payload TEXT NOT NULL,
 next_attempt_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Não | [accounts](accounts.md) |
| conversation_id | TEXT |  | Não | [conversations](conversations.md) |
| id | TEXT |  | Sim |  |
| message_id | TEXT |  | Não | [messages](messages.md) |
| next_attempt_at | TEXT |  | Não |  |
| payload | TEXT |  | Não |  |
| revision_id | TEXT |  | Não |  |
| state | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(state IN ('queued','dispatching','submitted','unknown','failed')) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 2) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_message_reaction_commands_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `message_reaction_active`

  ```sql
  CREATE UNIQUE INDEX message_reaction_active ON message_reaction_commands(message_id) WHERE state IN ('queued','dispatching','unknown')
  ```

- `message_reaction_pending`

  ```sql
  CREATE INDEX message_reaction_pending ON message_reaction_commands(state,next_attempt_at,id)
  ```

- `reaction_command_message`

  ```sql
  CREATE INDEX reaction_command_message ON message_reaction_commands(message_id)
  ```

- `sqlite_autoindex_message_reaction_commands_1`

  ```sql
  PRIMARY KEY (id)
  ```
