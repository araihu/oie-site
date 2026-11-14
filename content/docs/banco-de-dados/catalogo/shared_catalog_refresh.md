---
title: "shared_catalog_refresh"
description: "Definição SQL, colunas e relações de shared_catalog_refresh no banco de dados do Oiê."
language: pt-BR
---

# `shared_catalog_refresh`

[Catálogo de tabelas e views](index.md)

Disponível a partir da migração `000027`, posterior à `v0.0.1-rc.9`.

## Definição SQL

```sql
CREATE VIEW shared_catalog_refresh AS SELECT '' AS message_id WHERE 0
```

## Colunas

| Nome | Tipo | Padrão | Aceita NULL |
| ---- | ---- | ------ | ----------- |
| message_id |  |  | Sim |


## Gatilhos

### `shared_catalog_refresh_source`

```sql
CREATE TRIGGER shared_catalog_refresh_source INSTEAD OF INSERT ON shared_catalog_refresh BEGIN
 DELETE FROM shared_catalog_sources WHERE message_id=NEW.message_id
 AND NOT EXISTS (SELECT 1 FROM shared_catalog_visible_sources WHERE message_id=NEW.message_id);
 INSERT INTO shared_catalog_sources SELECT * FROM shared_catalog_visible_sources WHERE message_id=NEW.message_id
 ON CONFLICT(message_id) DO UPDATE SET conversation_id=excluded.conversation_id,
 position_ms=excluded.position_ms,parser_version=excluded.parser_version;
 DELETE FROM shared_catalog_order WHERE reference_id IN (SELECT id FROM shared_content WHERE message_id=NEW.message_id);
 INSERT INTO shared_catalog_order SELECT scope,tab,position_ms,reference_id FROM shared_catalog_order_rows WHERE message_id=NEW.message_id;
END
```
