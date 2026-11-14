---
title: Catálogo de tabelas e views
description: Colunas, chaves, índices, gatilhos e definições SQL do banco SQLite do Oiê.
language: pt-BR
---

# Catálogo de tabelas e views

Estrutura SQLite gerada a partir das [migrações do Oiê](https://github.com/araihu/oie/tree/dd2e6106e56f34c86a1b74c6dd989ecdbd350bc6/internal/adapters/sqlite/migrations). Consulte os [diagramas por assunto](../index.md#diagramas-por-assunto) para localizar as relações principais.

As tabelas e views introduzidas pela migração `000027` exigem uma versão posterior à `v0.0.1-rc.9`.

## Tabelas e views

| Nome | Colunas | Descrição | Tipo |
| ---- | ------- | ------- | ---- |
| [accounts](accounts.md) | 6 |  | Tabela |
| [archive_aliases](archive_aliases.md) | 4 |  | Tabela |
| [archive_checkpoints](archive_checkpoints.md) | 3 |  | Tabela |
| [avatar_revisions](avatar_revisions.md) | 5 |  | Tabela |
| [avatars](avatars.md) | 5 |  | Tabela |
| [changes](changes.md) | 4 |  | Tabela |
| [contact_revisions](contact_revisions.md) | 5 |  | Tabela |
| [contact_sources](contact_sources.md) | 4 |  | Tabela |
| [contacts](contacts.md) | 9 |  | Tabela |
| [conversation_read_revisions](conversation_read_revisions.md) | 6 |  | Tabela |
| [conversation_read_states](conversation_read_states.md) | 3 |  | Tabela |
| [conversation_revisions](conversation_revisions.md) | 5 |  | Tabela |
| [conversations](conversations.md) | 10 |  | Tabela |
| [engine_identity](engine_identity.md) | 2 |  | Tabela |
| [hidden_scheduled_messages](hidden_scheduled_messages.md) | 1 | Mensagens agendadas que ficam fora das consultas da linha do tempo. | View |
| [history_status](history_status.md) | 2 |  | Tabela |
| [media_files](media_files.md) | 4 |  | Tabela |
| [message_canonical_aliases](message_canonical_aliases.md) | 2 |  | Tabela |
| [message_content_heads](message_content_heads.md) | 5 |  | Tabela |
| [message_correction_facts](message_correction_facts.md) | 7 |  | Tabela |
| [message_edit_commands](message_edit_commands.md) | 9 |  | Tabela |
| [message_previews](message_previews.md) | 10 |  | Tabela |
| [message_reaction_commands](message_reaction_commands.md) | 8 |  | Tabela |
| [message_reaction_observations](message_reaction_observations.md) | 18 |  | Tabela |
| [message_reads](message_reads.md) | 2 |  | Tabela |
| [message_references](message_references.md) | 5 |  | Tabela |
| [message_replies](message_replies.md) | 3 |  | Tabela |
| [message_revisions](message_revisions.md) | 5 |  | Tabela |
| [message_sender_identities](message_sender_identities.md) | 5 |  | Tabela |
| [message_star_revisions](message_star_revisions.md) | 5 |  | Tabela |
| [message_stars](message_stars.md) | 5 |  | Tabela |
| [messages](messages.md) | 10 | Mensagens armazenadas pelo motor por conta e conversa. | Tabela |
| [operations](operations.md) | 6 |  | Tabela |
| [original_transmission_order](original_transmission_order.md) | 3 |  | Tabela |
| [outbox](outbox.md) | 6 | Fila de envio do motor, incluindo mensagens agendadas. | Tabela |
| [preview_revisions](preview_revisions.md) | 5 |  | Tabela |
| [revisions](revisions.md) | 5 |  | Tabela |
| [room_archives](room_archives.md) | 3 |  | Tabela |
| [room_intents](room_intents.md) | 10 |  | Tabela |
| [room_invitation_credentials](room_invitation_credentials.md) | 2 |  | Tabela |
| [room_reaction_privacy](room_reaction_privacy.md) | 5 |  | Tabela |
| [shared_catalog_counts](shared_catalog_counts.md) | 3 |  | Tabela |
| [shared_catalog_order](shared_catalog_order.md) | 4 | Entradas que ordenam as referências por escopo e aba do catálogo. | Tabela |
| [shared_catalog_order_rows](shared_catalog_order_rows.md) | 5 |  | View |
| [shared_catalog_refresh](shared_catalog_refresh.md) | 1 |  | View |
| [shared_catalog_sources](shared_catalog_sources.md) | 4 |  | Tabela |
| [shared_catalog_visible_sources](shared_catalog_visible_sources.md) | 4 |  | View |
| [shared_content](shared_content.md) | 6 | Referências a mídia, links e documentos identificados nas mensagens. | Tabela |
| [shared_content_state](shared_content_state.md) | 3 |  | Tabela |
| [staged_attachments](staged_attachments.md) | 6 |  | Tabela |
