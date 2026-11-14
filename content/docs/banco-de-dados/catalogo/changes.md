---
title: "changes"
description: "Definição SQL, colunas e relações de changes no banco de dados do Oiê."
language: pt-BR
---

# `changes`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE TABLE changes (
    sequence INTEGER PRIMARY KEY AUTOINCREMENT,
    payload TEXT NOT NULL,
    occurred_at TEXT NOT NULL
, kind TEXT NOT NULL DEFAULT 'account.registered')
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| kind | TEXT | 'account.registered' | Não |
| occurred_at | TEXT |  | Não |
| payload | TEXT |  | Não |
| sequence | INTEGER |  | Sim |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| sequence | PRIMARY KEY | PRIMARY KEY (sequence) |
