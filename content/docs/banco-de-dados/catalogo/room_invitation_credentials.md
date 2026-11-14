---
title: "room_invitation_credentials"
description: "Definição SQL, colunas e relações de room_invitation_credentials no banco de dados do Oiê."
language: pt-BR
---

# `room_invitation_credentials`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE room_invitation_credentials (
    conversation_id TEXT PRIMARY KEY REFERENCES conversations(id) ON DELETE CASCADE,
    password TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| conversation_id | TEXT |  | Sim | [conversations](conversations.md) |
| password | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE CASCADE MATCH NONE |
| conversation_id | PRIMARY KEY | PRIMARY KEY (conversation_id) |
| sqlite_autoindex_room_invitation_credentials_1 | PRIMARY KEY | PRIMARY KEY (conversation_id) |


## Índices

- `sqlite_autoindex_room_invitation_credentials_1`

  ```sql
  PRIMARY KEY (conversation_id)
  ```
