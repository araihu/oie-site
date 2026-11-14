---
title: "outbox"
description: "Definição SQL, colunas e relações de outbox no banco de dados do Oiê."
language: pt-BR
---

# `outbox`

[Catálogo de tabelas e views](index.md)

Fila de envio do motor, incluindo mensagens agendadas.

## Definição SQL

```sql
CREATE TABLE outbox (
 message_id TEXT PRIMARY KEY REFERENCES messages(id), state TEXT NOT NULL,
 attempt_count INTEGER NOT NULL DEFAULT 0, next_attempt_at TEXT NOT NULL
, scheduled_at TEXT NOT NULL DEFAULT '', wire_stage TEXT NOT NULL DEFAULT 'untracked'
 CHECK(wire_stage IN ('untracked', 'awaiting', 'started')))
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL | Referencia |
| ---- | ---- | ------ | ----------- | ---------- |
| attempt_count | INTEGER | 0 | Não |  |
| message_id | TEXT |  | Sim | [messages](messages.md) |
| next_attempt_at | TEXT |  | Não |  |
| scheduled_at | TEXT | '' | Não |  |
| state | TEXT |  | Não |  |
| wire_stage | TEXT | 'untracked' | Não |  |


## Restrições

| Nome | Tipo | Definição |
| ---- | ---- | ---------- |
| - | CHECK | CHECK(wire_stage IN ('untracked', 'awaiting', 'started')) |
| - (Foreign key ID: 0) | FOREIGN KEY | FOREIGN KEY (message_id) REFERENCES messages (id) ON UPDATE NO ACTION ON DELETE NO ACTION MATCH NONE |
| message_id | PRIMARY KEY | PRIMARY KEY (message_id) |
| sqlite_autoindex_outbox_1 | PRIMARY KEY | PRIMARY KEY (message_id) |


## Índices

- `outbox_dispatch_due`

  ```sql
  CREATE INDEX outbox_dispatch_due ON outbox(next_attempt_at,message_id) WHERE state='queued'
  ```

- `outbox_schedule`

  ```sql
  CREATE INDEX outbox_schedule ON outbox(scheduled_at, message_id)
WHERE state = 'queued' AND scheduled_at <> ''
  ```

- `sqlite_autoindex_outbox_1`

  ```sql
  PRIMARY KEY (message_id)
  ```


## Gatilhos

### `shared_catalog_outbox_delete`

```sql
CREATE TRIGGER shared_catalog_outbox_delete AFTER DELETE ON outbox
WHEN OLD.scheduled_at<>'' OR OLD.state='cancelled' BEGIN
 INSERT INTO shared_catalog_refresh VALUES(OLD.message_id);
END
```

### `shared_catalog_outbox_insert`

```sql
CREATE TRIGGER shared_catalog_outbox_insert AFTER INSERT ON outbox WHEN NEW.scheduled_at<>'' BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.message_id);
END
```

### `shared_catalog_outbox_update`

```sql
CREATE TRIGGER shared_catalog_outbox_update AFTER UPDATE OF state,scheduled_at ON outbox
WHEN (OLD.state<>NEW.state OR OLD.scheduled_at<>NEW.scheduled_at)
AND (OLD.scheduled_at<>'' OR NEW.scheduled_at<>'' OR OLD.state='cancelled' OR NEW.state='cancelled') BEGIN
 INSERT INTO shared_catalog_refresh VALUES(NEW.message_id);
END
```
