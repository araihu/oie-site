---
title: "contact_revisions"
description: "Definição SQL, colunas e relações de contact_revisions no banco de dados do Oiê."
language: pt-BR
---

# `contact_revisions`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE contact_revisions (
 id TEXT PRIMARY KEY,
 entity_id TEXT NOT NULL REFERENCES contacts(id),
 parents TEXT NOT NULL CHECK(json_valid(parents)),
 payload TEXT NOT NULL CHECK(json_valid(payload)),
 created_at TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| created_at | TEXT |  | Não |  |
| entity_id | TEXT |  | Não | [contacts](contacts.md) |
| id | TEXT |  | Sim |  |
| parents | TEXT |  | Não |  |
| payload | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(json_valid(parents)) |
| - | CHECK | CHECK(json_valid(payload)) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (entity_id) REFERENCES contacts (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_contact_revisions_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `sqlite_autoindex_contact_revisions_1`

  ```sql
  PRIMARY KEY (id)
  ```
