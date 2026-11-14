---
title: "messages"
description: "Definição SQL, colunas e relações de messages no banco de dados do Oiê."
language: pt-BR
---

# `messages`

[Catálogo de tabelas e views](index.md)

Mensagens armazenadas pelo motor por conta e conversa.

## Definição SQL

```sql
CREATE TABLE messages (
 id TEXT PRIMARY KEY, revision_id TEXT NOT NULL, account_id TEXT NOT NULL REFERENCES accounts(id),
 conversation_id TEXT NOT NULL REFERENCES conversations(id), provider_id TEXT NOT NULL,
 is_own INTEGER NOT NULL, state TEXT NOT NULL, payload TEXT NOT NULL
, created_at_ms INTEGER NOT NULL DEFAULT 0, timeline_at_ms INTEGER DEFAULT NULL)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia | Descrição |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- | --------- |
| account_id | TEXT |  | Não |  | [accounts](accounts.md) |  |
| conversation_id | TEXT |  | Não |  | [conversations](conversations.md) |  |
| created_at_ms | INTEGER | 0 | Não |  |  |  |
| id | TEXT |  | Sim | [archive_aliases](archive_aliases.md) [message_canonical_aliases](message_canonical_aliases.md) [message_content_heads](message_content_heads.md) [message_correction_facts](message_correction_facts.md) [message_edit_commands](message_edit_commands.md) [message_previews](message_previews.md) [message_reaction_commands](message_reaction_commands.md) [message_reads](message_reads.md) [message_references](message_references.md) [message_replies](message_replies.md) [message_revisions](message_revisions.md) [message_sender_identities](message_sender_identities.md) [message_star_revisions](message_star_revisions.md) [message_stars](message_stars.md) [original_transmission_order](original_transmission_order.md) [outbox](outbox.md) [shared_catalog_sources](shared_catalog_sources.md) [shared_content](shared_content.md) [shared_content_state](shared_content_state.md) |  |  |
| is_own | INTEGER |  | Não |  |  |  |
| payload | TEXT |  | Não |  |  |  |
| provider_id | TEXT |  | Não |  |  |  |
| revision_id | TEXT |  | Não |  |  |  |
| state | TEXT |  | Não |  |  |  |
| timeline_at_ms | INTEGER | NULL | Sim |  |  | Posição na linha do tempo, quando definida. A consulta usa `created_at_ms` como alternativa. |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_messages_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `messages_chronological`

  ```sql
  CREATE INDEX messages_chronological ON messages(conversation_id,created_at_ms DESC,id DESC)
  ```

- `messages_conversation_page`

  ```sql
  CREATE INDEX messages_conversation_page ON messages(conversation_id,id DESC)
  ```

- `messages_provider_identity`

  ```sql
  CREATE UNIQUE INDEX messages_provider_identity ON messages(account_id,conversation_id,provider_id,is_own) WHERE provider_id <> ''
  ```

- `messages_read_candidates`

  ```sql
  CREATE INDEX messages_read_candidates ON messages(conversation_id,is_own)
  ```

- `messages_timeline`

  ```sql
  CREATE INDEX messages_timeline ON messages(conversation_id, COALESCE(timeline_at_ms, created_at_ms) DESC, id DESC)
  ```

- `sqlite_autoindex_messages_1`

  ```sql
  PRIMARY KEY (id)
  ```


## Gatilhos

### `shared_catalog_message_insert`

```sql
CREATE TRIGGER shared_catalog_message_insert AFTER INSERT ON messages BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.id);
END
```

### `shared_catalog_message_position`

```sql
CREATE TRIGGER shared_catalog_message_position AFTER UPDATE OF timeline_at_ms,created_at_ms ON messages
WHEN OLD.timeline_at_ms IS NOT NEW.timeline_at_ms OR OLD.created_at_ms<>NEW.created_at_ms BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.id);
END
```
