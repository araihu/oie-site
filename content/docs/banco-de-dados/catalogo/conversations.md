---
title: "conversations"
description: "Definição SQL, colunas e relações de conversations no banco de dados do Oiê."
language: pt-BR
---

# `conversations`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE conversations (
 id TEXT PRIMARY KEY, revision_id TEXT NOT NULL, account_id TEXT NOT NULL REFERENCES accounts(id),
 address TEXT NOT NULL, peer_id TEXT NOT NULL, payload TEXT NOT NULL, activity_sequence INTEGER NOT NULL,
 unread_count INTEGER NOT NULL DEFAULT 0, archived INTEGER NOT NULL DEFAULT 0, favorite INTEGER NOT NULL DEFAULT 0,
 UNIQUE(account_id,address)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| account_id | TEXT |  | Não |  | [accounts](accounts.md) |
| activity_sequence | INTEGER |  | Não |  |  |
| address | TEXT |  | Não |  |  |
| archived | INTEGER | 0 | Não |  |  |
| favorite | INTEGER | 0 | Não |  |  |
| id | TEXT |  | Sim | [conversation_read_revisions](conversation_read_revisions.md) [conversation_read_states](conversation_read_states.md) [conversation_revisions](conversation_revisions.md) [message_edit_commands](message_edit_commands.md) [message_previews](message_previews.md) [message_reaction_commands](message_reaction_commands.md) [message_sender_identities](message_sender_identities.md) [messages](messages.md) [original_transmission_order](original_transmission_order.md) [room_archives](room_archives.md) [room_intents](room_intents.md) [room_invitation_credentials](room_invitation_credentials.md) [staged_attachments](staged_attachments.md) |  |
| payload | TEXT |  | Não |  |  |
| peer_id | TEXT |  | Não |  |  |
| revision_id | TEXT |  | Não |  |  |
| unread_count | INTEGER | 0 | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_conversations_1 | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_conversations_2 | UNIQUE | UNIQUE (account_id, address) |


## Índices

- `sqlite_autoindex_conversations_1`

  ```sql
  PRIMARY KEY (id)
  ```

- `sqlite_autoindex_conversations_2`

  ```sql
  UNIQUE (account_id, address)
  ```


## Gatilhos

### `shared_catalog_request_update`

```sql
CREATE TRIGGER shared_catalog_request_update AFTER UPDATE OF payload ON conversations
WHEN (json_extract(OLD.payload,'$.Request') IS NULL)<>(json_extract(NEW.payload,'$.Request') IS NULL) BEGIN
 INSERT INTO shared_catalog_refresh SELECT id FROM messages WHERE conversation_id=NEW.id;
END
```
