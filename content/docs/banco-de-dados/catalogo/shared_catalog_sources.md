---
title: "shared_catalog_sources"
description: "Definição SQL, colunas e relações de shared_catalog_sources no banco de dados do Oiê."
language: pt-BR
---

# `shared_catalog_sources`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE shared_catalog_sources (
 message_id TEXT PRIMARY KEY REFERENCES messages(id) ON DELETE CASCADE,
 conversation_id TEXT NOT NULL,
 position_ms INTEGER NOT NULL,
 parser_version INTEGER NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| conversation_id | TEXT |  | Não |  |
| message_id | TEXT |  | Sim | [messages](messages.md) |
| parser_version | INTEGER |  | Não |  |
| position_ms | INTEGER |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE CASCADE MATCH NONE |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| sqlite_autoindex_shared_catalog_sources_1 | PRIMARY KEY | PRIMARY KEY (message_id) |


## Índices

- `shared_catalog_pending`

  ```sql
  CREATE INDEX shared_catalog_pending ON shared_catalog_sources(parser_version,position_ms DESC,message_id DESC)
  ```

- `sqlite_autoindex_shared_catalog_sources_1`

  ```sql
  PRIMARY KEY (message_id)
  ```


## Gatilhos

### `shared_catalog_count_delete`

```sql
CREATE TRIGGER shared_catalog_count_delete AFTER DELETE ON shared_catalog_sources BEGIN
 UPDATE shared_catalog_counts SET messages=messages-1
 WHERE conversation_id=OLD.conversation_id AND parser_version=OLD.parser_version;
 DELETE FROM shared_catalog_counts
 WHERE conversation_id=OLD.conversation_id AND parser_version=OLD.parser_version AND messages=0;
END
```

### `shared_catalog_count_insert`

```sql
CREATE TRIGGER shared_catalog_count_insert AFTER INSERT ON shared_catalog_sources BEGIN
 INSERT INTO shared_catalog_counts VALUES(NEW.conversation_id,NEW.parser_version,1)
 ON CONFLICT(conversation_id,parser_version) DO UPDATE SET messages=messages+1;
END
```

### `shared_catalog_count_update`

```sql
CREATE TRIGGER shared_catalog_count_update AFTER UPDATE OF parser_version,conversation_id ON shared_catalog_sources
WHEN OLD.parser_version<>NEW.parser_version OR OLD.conversation_id<>NEW.conversation_id BEGIN
 UPDATE shared_catalog_counts SET messages=messages-1
 WHERE conversation_id=OLD.conversation_id AND parser_version=OLD.parser_version;
 DELETE FROM shared_catalog_counts
 WHERE conversation_id=OLD.conversation_id AND parser_version=OLD.parser_version AND messages=0;
 INSERT INTO shared_catalog_counts VALUES(NEW.conversation_id,NEW.parser_version,1)
 ON CONFLICT(conversation_id,parser_version) DO UPDATE SET messages=messages+1;
END
```
