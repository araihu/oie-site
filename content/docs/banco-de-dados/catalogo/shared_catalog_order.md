---
title: "shared_catalog_order"
description: "Definição SQL, colunas e relações de shared_catalog_order no banco de dados do Oiê."
language: pt-BR
---

# `shared_catalog_order`

[Catálogo de tabelas e views](index.md)

Disponível a partir da migração `000027`, posterior à `v0.0.1-rc.9`.

Entradas que ordenam as referências por escopo e aba do catálogo.

## Definição SQL

```sql
CREATE TABLE shared_catalog_order (
 scope TEXT NOT NULL,
 tab TEXT NOT NULL,
 position_ms INTEGER NOT NULL,
 reference_id TEXT NOT NULL REFERENCES shared_content(id) ON DELETE CASCADE,
 PRIMARY KEY(scope,tab,position_ms,reference_id)
) WITHOUT ROWID
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| position_ms | INTEGER |  | Não |  |
| reference_id | TEXT |  | Não | [shared_content](shared_content.md) |
| scope | TEXT |  | Não |  |
| tab | TEXT |  | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (reference_id) REFERENCES shared_content (id) ON UPDATE NO ACTION ON DELETE CASCADE MATCH NONE |
| position_ms | PRIMARY KEY | PRIMARY KEY (position_ms) |
| reference_id | PRIMARY KEY | PRIMARY KEY (reference_id) |
| scope | PRIMARY KEY | PRIMARY KEY (scope) |
| sqlite_autoindex_shared_catalog_order_1 | PRIMARY KEY | PRIMARY KEY (scope, tab, position_ms, reference_id) |
| tab | PRIMARY KEY | PRIMARY KEY (tab) |


## Índices

- `shared_catalog_order_reference`

  ```sql
  CREATE INDEX shared_catalog_order_reference ON shared_catalog_order(reference_id)
  ```

- `sqlite_autoindex_shared_catalog_order_1`

  ```sql
  PRIMARY KEY (scope, tab, position_ms, reference_id)
  ```
