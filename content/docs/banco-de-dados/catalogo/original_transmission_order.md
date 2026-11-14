---
title: "original_transmission_order"
description: "Definição SQL, colunas e relações de original_transmission_order no banco de dados do Oiê."
language: pt-BR
---

# `original_transmission_order`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE original_transmission_order (
 ordinal INTEGER PRIMARY KEY AUTOINCREMENT,
 message_id TEXT NOT NULL REFERENCES messages(id),
 conversation_id TEXT NOT NULL REFERENCES conversations(id)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| conversation_id | TEXT |  | Não | [conversations](conversations.md) |
| message_id | TEXT |  | Não | [messages](messages.md) |
| ordinal | INTEGER |  | Sim |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| ordinal | PRIMARY KEY | PRIMARY KEY (ordinal) |


## Índices

- `original_transmission_conversation`

  ```sql
  CREATE INDEX original_transmission_conversation ON original_transmission_order(conversation_id,ordinal DESC)
  ```

- `original_transmission_message`

  ```sql
  CREATE INDEX original_transmission_message ON original_transmission_order(message_id,ordinal DESC)
  ```
