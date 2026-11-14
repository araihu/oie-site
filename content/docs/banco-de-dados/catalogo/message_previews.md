---
title: "message_previews"
description: "Definição SQL, colunas e relações de message_previews no banco de dados do Oiê."
language: pt-BR
---

# `message_previews`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_previews (
 id TEXT PRIMARY KEY, revision_id TEXT NOT NULL,
 account_id TEXT NOT NULL REFERENCES accounts(id),
 conversation_id TEXT NOT NULL REFERENCES conversations(id),
 message_id TEXT NOT NULL REFERENCES messages(id),
 url TEXT NOT NULL, state TEXT NOT NULL, download_state TEXT NOT NULL,
 updated_at TEXT NOT NULL, payload TEXT NOT NULL,
 UNIQUE(message_id,url)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| account_id | TEXT |  | Não |  | [accounts](accounts.md) |
| conversation_id | TEXT |  | Não |  | [conversations](conversations.md) |
| download_state | TEXT |  | Não |  |  |
| id | TEXT |  | Sim | [preview_revisions](preview_revisions.md) |  |
| message_id | TEXT |  | Não |  | [messages](messages.md) |
| payload | TEXT |  | Não |  |  |
| revision_id | TEXT |  | Não |  |  |
| state | TEXT |  | Não |  |  |
| updated_at | TEXT |  | Não |  |  |
| url | TEXT |  | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 2) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_message_previews_1 | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_message_previews_2 | UNIQUE | UNIQUE (message_id, url) |


## Índices

- `message_previews_conversation`

  ```sql
  CREATE INDEX message_previews_conversation ON message_previews(conversation_id,message_id)
  ```

- `message_previews_jobs`

  ```sql
  CREATE INDEX message_previews_jobs ON message_previews(download_state,state,updated_at)
  ```

- `sqlite_autoindex_message_previews_1`

  ```sql
  PRIMARY KEY (id)
  ```

- `sqlite_autoindex_message_previews_2`

  ```sql
  UNIQUE (message_id, url)
  ```


## Gatilhos

### `shared_catalog_preview_delete`

```sql
CREATE TRIGGER shared_catalog_preview_delete AFTER DELETE ON message_previews
WHEN json_extract(OLD.payload,'$.Kind')='image' BEGIN
 INSERT INTO shared_catalog_refresh VALUES(OLD.message_id);
END
```

### `shared_catalog_preview_insert`

```sql
CREATE TRIGGER shared_catalog_preview_insert AFTER INSERT ON message_previews
WHEN json_extract(NEW.payload,'$.Kind')='image' BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.message_id);
END
```

### `shared_catalog_preview_update`

```sql
CREATE TRIGGER shared_catalog_preview_update AFTER UPDATE OF payload ON message_previews
WHEN (json_extract(OLD.payload,'$.Kind') IS 'image')<>(json_extract(NEW.payload,'$.Kind') IS 'image') BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.message_id);
END
```
