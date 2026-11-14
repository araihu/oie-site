---
title: "message_reaction_observations"
description: "Definição SQL, colunas e relações de message_reaction_observations no banco de dados do Oiê."
language: pt-BR
---

# `message_reaction_observations`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_reaction_observations (
 ordinal INTEGER PRIMARY KEY AUTOINCREMENT,
 id TEXT NOT NULL UNIQUE,
 account_id TEXT NOT NULL REFERENCES accounts(id),
 address TEXT NOT NULL,
 kind TEXT NOT NULL,
 actor TEXT NOT NULL,
 provider_id TEXT NOT NULL,
 archive_scope TEXT NOT NULL,
 archive_id TEXT NOT NULL,
 target_id TEXT NOT NULL,
 target_scope TEXT NOT NULL,
 target_author TEXT NOT NULL,
 fingerprint TEXT NOT NULL,
 occurred_at TEXT NOT NULL,
 payload TEXT NOT NULL,
 command_id TEXT NOT NULL DEFAULT ''
, archive_order INTEGER NOT NULL DEFAULT 0, room_epoch INTEGER NOT NULL DEFAULT 0)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| account_id | TEXT |  | Não | [accounts](accounts.md) |
| actor | TEXT |  | Não |  |
| address | TEXT |  | Não |  |
| archive_id | TEXT |  | Não |  |
| archive_order | INTEGER | 0 | Não |  |
| archive_scope | TEXT |  | Não |  |
| command_id | TEXT | '' | Não |  |
| fingerprint | TEXT |  | Não |  |
| id | TEXT |  | Não |  |
| kind | TEXT |  | Não |  |
| occurred_at | TEXT |  | Não |  |
| ordinal | INTEGER |  | Sim |  |
| payload | TEXT |  | Não |  |
| provider_id | TEXT |  | Não |  |
| room_epoch | INTEGER | 0 | Não |  |
| target_author | TEXT |  | Não |  |
| target_id | TEXT |  | Não |  |
| target_scope | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (account_id) REFERENCES accounts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| ordinal | PRIMARY KEY | PRIMARY KEY (ordinal) |
| sqlite_autoindex_message_reaction_observations_1 | UNIQUE | UNIQUE (id) |


## Índices

- `reaction_observation_archive`

  ```sql
  CREATE INDEX reaction_observation_archive ON message_reaction_observations(account_id,archive_scope,archive_id)
  ```

- `reaction_observation_archive_order`

  ```sql
  CREATE INDEX reaction_observation_archive_order ON message_reaction_observations(account_id,archive_scope,occurred_at,archive_order) WHERE archive_id<>''
  ```

- `reaction_observation_command`

  ```sql
  CREATE INDEX reaction_observation_command ON message_reaction_observations(command_id) WHERE command_id<>''
  ```

- `reaction_observation_identity`

  ```sql
  CREATE INDEX reaction_observation_identity ON message_reaction_observations(account_id,address,actor,provider_id)
  ```

- `reaction_observation_target`

  ```sql
  CREATE INDEX reaction_observation_target ON message_reaction_observations(account_id,address,target_id)
  ```

- `sqlite_autoindex_message_reaction_observations_1`

  ```sql
  UNIQUE (id)
  ```
