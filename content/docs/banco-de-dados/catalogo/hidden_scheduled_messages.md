---
title: "hidden_scheduled_messages"
description: "Definição SQL, colunas e relações de hidden_scheduled_messages no banco de dados do Oiê."
language: pt-BR
---

# `hidden_scheduled_messages`

[Catálogo de tabelas e views](index.md)

Mensagens agendadas que ficam fora das consultas da linha do tempo.

## Definição SQL

```sql
CREATE VIEW hidden_scheduled_messages AS
SELECT o.message_id FROM outbox o JOIN messages m ON m.id = o.message_id
WHERE (o.scheduled_at <> '' AND m.timeline_at_ms IS NULL AND o.state IN ('queued', 'dispatching', 'cancelled'))
   OR (o.state = 'cancelled' AND m.timeline_at_ms IS NOT NULL)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| message_id | TEXT |  | Sim |


## Objetos referenciados

| Nome | Colunas | Descrição | Tipo |
| ---- | ------- | ------- | ---- |
| [outbox](outbox.md) | 6 | Fila de envio do motor, incluindo mensagens agendadas. | table |
| [messages](messages.md) | 10 | Mensagens armazenadas pelo motor por conta e conversa. | table |
