---
title: "message_correction_facts"
description: "Definição SQL, colunas e relações de message_correction_facts no banco de dados do Oiê."
language: pt-BR
---

# `message_correction_facts`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE message_correction_facts (
 message_id TEXT NOT NULL REFERENCES messages(id),
 scope TEXT NOT NULL, author TEXT NOT NULL, remote_id TEXT NOT NULL, target_id TEXT NOT NULL,
 source_hash TEXT NOT NULL, source_revision_id TEXT NOT NULL REFERENCES message_revisions(id),
 PRIMARY KEY(message_id,scope,author,remote_id,target_id,source_hash)
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| author | TEXT |  | Não |  |
| message_id | TEXT |  | Não | [messages](messages.md) |
| remote_id | TEXT |  | Não |  |
| scope | TEXT |  | Não |  |
| source_hash | TEXT |  | Não |  |
| source_revision_id | TEXT |  | Não | [message_revisions](message_revisions.md) |
| target_id | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (source_revision_id) REFERENCES message_revisions (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| - (Foreign key ID: 1) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| author | PRIMARY KEY | PRIMARY KEY (author) |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| remote_id | PRIMARY KEY | PRIMARY KEY (remote_id) |
| scope | PRIMARY KEY | PRIMARY KEY (scope) |
| source_hash | PRIMARY KEY | PRIMARY KEY (source_hash) |
| sqlite_autoindex_message_correction_facts_1 | PRIMARY KEY | PRIMARY KEY (message_id, scope, author, remote_id, target_id, source_hash) |
| target_id | PRIMARY KEY | PRIMARY KEY (target_id) |


## Índices

- `correction_source_identity`

  ```sql
  CREATE INDEX correction_source_identity ON message_correction_facts(scope,author,remote_id)
  ```

- `correction_target_identity`

  ```sql
  CREATE INDEX correction_target_identity ON message_correction_facts(scope,author,target_id)
  ```

- `sqlite_autoindex_message_correction_facts_1`

  ```sql
  PRIMARY KEY (message_id, scope, author, remote_id, target_id, source_hash)
  ```
