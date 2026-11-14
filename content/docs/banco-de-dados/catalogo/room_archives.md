---
title: "room_archives"
description: "Definição SQL, colunas e relações de room_archives no banco de dados do Oiê."
language: pt-BR
---

# `room_archives`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE room_archives (
 conversation_id TEXT PRIMARY KEY REFERENCES conversations(id),
 mode TEXT NOT NULL DEFAULT 'gradual' CHECK(mode IN ('gradual','on_demand')),
 progress TEXT NOT NULL DEFAULT '{}'
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| conversation_id | TEXT |  | Sim | [conversations](conversations.md) |
| mode | TEXT | 'gradual' | Não |  |
| progress | TEXT | '{}' | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(mode IN ('gradual','on_demand')) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| conversation_id | PRIMARY KEY | PRIMARY KEY (conversation_id) |
| sqlite_autoindex_room_archives_1 | PRIMARY KEY | PRIMARY KEY (conversation_id) |


## Índices

- `sqlite_autoindex_room_archives_1`

  ```sql
  PRIMARY KEY (conversation_id)
  ```
