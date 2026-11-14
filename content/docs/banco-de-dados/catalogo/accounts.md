---
title: "accounts"
description: "Definição SQL, colunas e relações de accounts no banco de dados do Oiê."
language: pt-BR
---

# `accounts`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE accounts (
    id TEXT PRIMARY KEY,
    revision_id TEXT NOT NULL UNIQUE REFERENCES revisions(id) DEFERRABLE INITIALLY DEFERRED,
    protocol TEXT NOT NULL,
    address TEXT NOT NULL,
    label TEXT NOT NULL,
    created_at TEXT NOT NULL,
    UNIQUE(protocol, address)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| address | TEXT |  | Não |  |  |
| created_at | TEXT |  | Não |  |  |
| id | TEXT |  | Sim | [archive_aliases](archive_aliases.md) [archive_checkpoints](archive_checkpoints.md) [avatars](avatars.md) [contact_sources](contact_sources.md) [contacts](contacts.md) [conversations](conversations.md) [history_status](history_status.md) [message_edit_commands](message_edit_commands.md) [message_previews](message_previews.md) [message_reaction_commands](message_reaction_commands.md) [message_reaction_observations](message_reaction_observations.md) [message_sender_identities](message_sender_identities.md) [messages](messages.md) [revisions](revisions.md) [room_intents](room_intents.md) [room_reaction_privacy](room_reaction_privacy.md) |  |
| label | TEXT |  | Não |  |  |
| protocol | TEXT |  | Não |  |  |
| revision_id | TEXT |  | Não |  | [revisions](revisions.md) |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (revision_id) REFERENCES revisions (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_accounts_1 | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_accounts_2 | UNIQUE | UNIQUE (revision_id) |
| sqlite_autoindex_accounts_3 | UNIQUE | UNIQUE (protocol, address) |


## Índices

- `sqlite_autoindex_accounts_1`

  ```sql
  PRIMARY KEY (id)
  ```

- `sqlite_autoindex_accounts_2`

  ```sql
  UNIQUE (revision_id)
  ```

- `sqlite_autoindex_accounts_3`

  ```sql
  UNIQUE (protocol, address)
  ```
