---
title: "message_canonical_aliases"
description: "Definição SQL, colunas e relações de message_canonical_aliases no banco de dados do Oiê."
language: pt-BR
---

# `message_canonical_aliases`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_canonical_aliases (
 alias_id TEXT PRIMARY KEY REFERENCES messages(id),
 canonical_id TEXT NOT NULL REFERENCES messages(id), CHECK(alias_id<>canonical_id)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| alias_id | TEXT |  | Sim | [messages](messages.md) |
| canonical_id | TEXT |  | Não | [messages](messages.md) |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(alias_id<>canonical_id) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (canonical_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (alias_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| alias_id | PRIMARY KEY | PRIMARY KEY (alias_id) |
| sqlite_autoindex_message_canonical_aliases_1 | PRIMARY KEY | PRIMARY KEY (alias_id) |


## Índices

- `correction_canonical_aliases`

  ```sql
  CREATE INDEX correction_canonical_aliases ON message_canonical_aliases(canonical_id)
  ```

- `sqlite_autoindex_message_canonical_aliases_1`

  ```sql
  PRIMARY KEY (alias_id)
  ```


## Gatilhos

### `shared_catalog_alias_delete`

```sql
CREATE TRIGGER shared_catalog_alias_delete AFTER DELETE ON message_canonical_aliases BEGIN
 INSERT INTO shared_catalog_refresh VALUES(OLD.alias_id);
END
```

### `shared_catalog_alias_insert`

```sql
CREATE TRIGGER shared_catalog_alias_insert AFTER INSERT ON message_canonical_aliases BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.alias_id);
END
```
