---
title: "contacts"
description: "Definição SQL, colunas e relações de contacts no banco de dados do Oiê."
language: pt-BR
---

# `contacts`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE contacts (
 id TEXT PRIMARY KEY,
 revision_id TEXT NOT NULL,
 account_id TEXT NOT NULL REFERENCES accounts(id),
 address TEXT NOT NULL,
 display_name TEXT NOT NULL,
 subscription TEXT NOT NULL CHECK(subscription IN ('none','to','from','both')),
 groups_json TEXT NOT NULL CHECK(json_valid(groups_json)),
 active INTEGER NOT NULL CHECK(active IN (0,1)),
 payload TEXT NOT NULL CHECK(json_valid(payload)),
 UNIQUE(account_id,address)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| account_id | TEXT |  | Não |  | [accounts](accounts.md) |
| active | INTEGER |  | Não |  |  |
| address | TEXT |  | Não |  |  |
| display_name | TEXT |  | Não |  |  |
| groups_json | TEXT |  | Não |  |  |
| id | TEXT |  | Sim | [contact_revisions](contact_revisions.md) |  |
| payload | TEXT |  | Não |  |  |
| revision_id | TEXT |  | Não |  |  |
| subscription | TEXT |  | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(subscription IN ('none','to','from','both')) |
| - | CHECK | CHECK(json_valid(groups_json)) |
| - | CHECK | CHECK(active IN (0,1)) |
| - | CHECK | CHECK(json_valid(payload)) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_contacts_1 | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_contacts_2 | UNIQUE | UNIQUE (account_id, address) |


## Índices

- `contacts_account`

  ```sql
  CREATE INDEX contacts_account ON contacts(account_id,id)
  ```

- `sqlite_autoindex_contacts_1`

  ```sql
  PRIMARY KEY (id)
  ```

- `sqlite_autoindex_contacts_2`

  ```sql
  UNIQUE (account_id, address)
  ```
