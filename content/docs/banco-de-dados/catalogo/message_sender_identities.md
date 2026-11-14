---
title: "message_sender_identities"
description: "Definição SQL, colunas e relações de message_sender_identities no banco de dados do Oiê."
language: pt-BR
---

# `message_sender_identities`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_sender_identities (
    message_id TEXT NOT NULL REFERENCES messages(id) ON DELETE CASCADE,
    account_id TEXT NOT NULL REFERENCES accounts(id),
    conversation_id TEXT NOT NULL REFERENCES conversations(id),
    personal_address TEXT NOT NULL CHECK(length(personal_address) BETWEEN 1 AND 1024),
    source_hash TEXT NOT NULL CHECK(length(source_hash)=64),
    PRIMARY KEY(message_id,personal_address,source_hash)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Não | [accounts](accounts.md) |
| conversation_id | TEXT |  | Não | [conversations](conversations.md) |
| message_id | TEXT |  | Não | [messages](messages.md) |
| personal_address | TEXT |  | Não |  |
| source_hash | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(length(personal_address) BETWEEN 1 AND 1024) |
| - | CHECK | CHECK(length(source_hash)=64) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (conversation_id) REFERENCES conversations (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 2) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE CASCADE MATCH NONE |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| personal_address | PRIMARY KEY | PRIMARY KEY (personal_address) |
| source_hash | PRIMARY KEY | PRIMARY KEY (source_hash) |
| sqlite_autoindex_message_sender_identities_1 | PRIMARY KEY | PRIMARY KEY (message_id, personal_address, source_hash) |


## Índices

- `sqlite_autoindex_message_sender_identities_1`

  ```sql
  PRIMARY KEY (message_id, personal_address, source_hash)
  ```
