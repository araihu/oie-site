---
title: "shared_content_state"
description: "Definição SQL, colunas e relações de shared_content_state no banco de dados do Oiê."
language: pt-BR
---

# `shared_content_state`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE shared_content_state (
 message_id TEXT PRIMARY KEY REFERENCES messages(id),
 parser_version INTEGER NOT NULL,
 caption TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| caption | TEXT |  | Não |  |
| message_id | TEXT |  | Sim | [messages](messages.md) |
| parser_version | INTEGER |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| sqlite_autoindex_shared_content_state_1 | PRIMARY KEY | PRIMARY KEY (message_id) |


## Índices

- `sqlite_autoindex_shared_content_state_1`

  ```sql
  PRIMARY KEY (message_id)
  ```


## Gatilhos

### `shared_catalog_state_delete`

```sql
CREATE TRIGGER shared_catalog_state_delete AFTER DELETE ON shared_content_state BEGIN
 INSERT INTO shared_catalog_refresh VALUES(OLD.message_id);
END
```

### `shared_catalog_state_insert`

```sql
CREATE TRIGGER shared_catalog_state_insert AFTER INSERT ON shared_content_state BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.message_id);
END
```

### `shared_catalog_state_update`

```sql
CREATE TRIGGER shared_catalog_state_update AFTER UPDATE OF parser_version ON shared_content_state BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.message_id);
END
```
