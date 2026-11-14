---
title: "operations"
description: "Definição SQL, colunas e relações de operations no banco de dados do Oiê."
language: pt-BR
---

# `operations`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE "operations" (
 id TEXT PRIMARY KEY, fingerprint TEXT NOT NULL, resource_id TEXT NOT NULL,
 result TEXT NOT NULL, created_at TEXT NOT NULL, kind TEXT NOT NULL
)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| created_at | TEXT |  | Não |
| fingerprint | TEXT |  | Não |
| id | TEXT |  | Sim |
| kind | TEXT |  | Não |
| resource_id | TEXT |  | Não |
| result | TEXT |  | Não |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| id | PRIMARY KEY | PRIMARY KEY (id) |
| sqlite_autoindex_operations_1 | PRIMARY KEY | PRIMARY KEY (id) |


## Índices

- `original_send_receipts`

  ```sql
  CREATE INDEX original_send_receipts ON operations(resource_id) WHERE kind='send_message'
  ```

- `sqlite_autoindex_operations_1`

  ```sql
  PRIMARY KEY (id)
  ```
