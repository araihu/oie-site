---
title: "avatars"
description: "Definição SQL, colunas e relações de avatars no banco de dados do Oiê."
language: pt-BR
---

# `avatars`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE avatars (
 id TEXT PRIMARY KEY,revision_id TEXT NOT NULL,account_id TEXT NOT NULL REFERENCES accounts(id),
 address TEXT NOT NULL,payload TEXT NOT NULL,UNIQUE(account_id,address)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referenciada por | Referencia |
| ---- | ---- | ------ | ----------- | ---------------- | ---------- |
| account_id | TEXT |  | Não |  | [accounts](accounts.md) |
| address | TEXT |  | Não |  |  |
| id | TEXT |  | Sim | [avatar_revisions](avatar_revisions.md) |  |
| payload | TEXT |  | Não |  |  |
| revision_id | TEXT |  | Não |  |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_avatars_1 | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_avatars_2 | UNIQUE | UNIQUE (account_id, address) |


## Índices

- `sqlite_autoindex_avatars_1`

  ```sql
  PRIMARY KEY (id)
  ```

- `sqlite_autoindex_avatars_2`

  ```sql
  UNIQUE (account_id, address)
  ```
