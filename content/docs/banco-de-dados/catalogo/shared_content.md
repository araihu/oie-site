---
title: "shared_content"
description: "Definição SQL, colunas e relações de shared_content no banco de dados do Oiê."
language: pt-BR
---

# `shared_content`

[Catálogo de tabelas e views](index.md)

Referências a mídia, links e documentos identificados nas mensagens.

## Definição SQL

```sql
CREATE TABLE shared_content (
 id TEXT PRIMARY KEY, revision_id TEXT NOT NULL,
 message_id TEXT NOT NULL REFERENCES messages(id),
 url TEXT NOT NULL, kind TEXT NOT NULL, file_name TEXT NOT NULL,
 UNIQUE(message_id,url)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| file_name | TEXT |  | Não |  |  |
| id | TEXT |  | Sim | [shared_catalog_order](shared_catalog_order.md) |  |
| kind | TEXT |  | Não |  |  |
| message_id | TEXT |  | Não |  | [messages](messages.md) |
| revision_id | TEXT |  | Não |  |  |
| url | TEXT |  | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_shared_content_1 | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_shared_content_2 | UNIQUE | UNIQUE (message_id, url) |


## Índices

- `shared_content_kind`

  ```sql
  CREATE INDEX shared_content_kind ON shared_content(kind,message_id,id)
  ```

- `shared_content_message`

  ```sql
  CREATE INDEX shared_content_message ON shared_content(message_id,id)
  ```

- `sqlite_autoindex_shared_content_1`

  ```sql
  PRIMARY KEY (id)
  ```

- `sqlite_autoindex_shared_content_2`

  ```sql
  UNIQUE (message_id, url)
  ```


## Gatilhos

### `shared_catalog_reference_insert`

```sql
CREATE TRIGGER shared_catalog_reference_insert AFTER INSERT ON shared_content BEGIN
 INSERT INTO shared_catalog_order SELECT scope,tab,position_ms,reference_id FROM shared_catalog_order_rows WHERE reference_id=NEW.id;
END
```

### `shared_catalog_reference_update`

```sql
CREATE TRIGGER shared_catalog_reference_update AFTER UPDATE OF kind ON shared_content WHEN OLD.kind<>NEW.kind BEGIN
 DELETE FROM shared_catalog_order WHERE reference_id=OLD.id;
 INSERT INTO shared_catalog_order SELECT scope,tab,position_ms,reference_id FROM shared_catalog_order_rows WHERE reference_id=NEW.id;
END
```
