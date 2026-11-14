---
title: "shared_catalog_order_rows"
description: "Definição SQL, colunas e relações de shared_catalog_order_rows no banco de dados do Oiê."
language: pt-BR
---

# `shared_catalog_order_rows`

[Catálogo de tabelas e views](index.md)

Disponível a partir da migração `000027`, posterior à `v0.0.1-rc.9`.

## Definição SQL

```sql
CREATE VIEW shared_catalog_order_rows AS
WITH kinds AS (
 SELECT sc.id AS reference_id,sc.message_id,m.account_id,s.conversation_id,s.position_ms,
 CASE WHEN json_extract(p.payload,'$.Kind')='image' OR sc.kind IN ('image','video','audio') THEN 'media'
      WHEN sc.kind='link' THEN 'links' WHEN sc.kind='document' THEN 'documents' END AS tab
 FROM shared_content sc JOIN shared_catalog_sources s ON s.message_id=sc.message_id
 JOIN messages m ON m.id=sc.message_id LEFT JOIN message_previews p ON p.message_id=m.id AND p.url=sc.url
 WHERE s.parser_version>=0
)
SELECT CASE scope.kind WHEN 0 THEN '' WHEN 1 THEN 'account:'||account_id ELSE 'conversation:'||conversation_id END AS scope,
 tab,position_ms,reference_id,message_id
FROM kinds CROSS JOIN (SELECT 0 AS kind UNION ALL SELECT 1 UNION ALL SELECT 2) scope
WHERE tab IS NOT NULL
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| message_id | TEXT |  | Sim |
| position_ms | INTEGER |  | Sim |
| reference_id | TEXT |  | Sim |
| scope |  |  | Sim |
| tab |  |  | Sim |


## Objetos referenciados

| Nome | Colunas | Descrição | Tipo |
| ---- | ------- | ------- | ---- |
| [shared_content](shared_content.md) | 6 | Referências a mídia, links e documentos identificados nas mensagens. | table |
| [shared_catalog_sources](shared_catalog_sources.md) | 4 |  | table |
| [messages](messages.md) | 10 | Mensagens armazenadas pelo motor por conta e conversa. | table |
| [message_previews](message_previews.md) | 10 |  | table |
