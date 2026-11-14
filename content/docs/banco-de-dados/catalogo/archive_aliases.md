---
title: "archive_aliases"
description: "Definição SQL, colunas e relações de archive_aliases no banco de dados do Oiê."
language: pt-BR
---

# `archive_aliases`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE archive_aliases (
 account_id TEXT NOT NULL REFERENCES accounts(id), scope TEXT NOT NULL, archive_id TEXT NOT NULL,
 message_id TEXT NOT NULL REFERENCES messages(id), PRIMARY KEY(account_id,scope,archive_id)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Não | [accounts](accounts.md) |
| archive_id | TEXT |  | Não |  |
| message_id | TEXT |  | Não | [messages](messages.md) |
| scope | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| account_id | PRIMARY KEY | PRIMARY KEY (account_id) |
| archive_id | PRIMARY KEY | PRIMARY KEY (archive_id) |
| scope | PRIMARY KEY | PRIMARY KEY (scope) |
| sqlite_autoindex_archive_aliases_1 | PRIMARY KEY | PRIMARY KEY (account_id, scope, archive_id) |


## Índices

- `sqlite_autoindex_archive_aliases_1`

  ```sql
  PRIMARY KEY (account_id, scope, archive_id)
  ```
