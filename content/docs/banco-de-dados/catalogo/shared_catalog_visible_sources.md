---
title: "shared_catalog_visible_sources"
description: "Definição SQL, colunas e relações de shared_catalog_visible_sources no banco de dados do Oiê."
language: pt-BR
---

# `shared_catalog_visible_sources`

[Catálogo de tabelas e views](index.md)

## Definição SQL

```sql
CREATE VIEW shared_catalog_visible_sources AS
SELECT m.id AS message_id,m.conversation_id,COALESCE(m.timeline_at_ms,m.created_at_ms) AS position_ms,
 COALESCE(s.parser_version,-1) AS parser_version
FROM messages m LEFT JOIN shared_content_state s ON s.message_id=m.id
WHERE EXISTS (SELECT 1 FROM conversations c WHERE c.id=m.conversation_id AND json_extract(c.payload,'$.Request') IS NULL)
AND NOT EXISTS (SELECT 1 FROM message_canonical_aliases a WHERE a.alias_id=m.id)
AND NOT EXISTS (SELECT 1 FROM hidden_scheduled_messages h WHERE h.message_id=m.id)
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| conversation_id | TEXT |  | Sim |
| message_id | TEXT |  | Sim |
| parser_version |  |  | Sim |
| position_ms |  |  | Sim |


## Objetos referenciados

| Nome | Colunas | Descrição | Tipo |
| ---- | ------- | ------- | ---- |
| [messages](messages.md) | 10 | Mensagens armazenadas pelo motor por conta e conversa. | table |
| [shared_content_state](shared_content_state.md) | 3 |  | table |
| [conversations](conversations.md) | 10 |  | table |
| [message_canonical_aliases](message_canonical_aliases.md) | 2 |  | table |
| [hidden_scheduled_messages](hidden_scheduled_messages.md) | 1 | Mensagens agendadas que ficam fora das consultas da linha do tempo. | view |
