---
title: Banco de dados
description: Estrutura SQLite do Oiê, diagramas por assunto e referência de tabelas, views, índices e gatilhos.
language: pt-BR
margo:
  imageViewer: true
---

# Banco de dados

O motor do Oiê usa SQLite para guardar contas, conversas, mensagens e envios pendentes. Para entender onde o banco fica em cada modo de execução, consulte [Arquitetura e deploy](../arquitetura/index.md).

Esta referência descreve o schema da versão de desenvolvimento. As tabelas e views do catálogo de conteúdo compartilhado adicionadas pela migração `000027` são posteriores à `v0.0.1-rc.9`.

## Catálogo de tabelas e views

Abra o [catálogo](catalogo/index.md) para consultar colunas, definições SQL, restrições, índices e gatilhos. As chaves estrangeiras levam às páginas dos objetos relacionados.

## Diagramas por assunto

Cada diagrama mostra as chaves primárias e as colunas ligadas pelas chaves estrangeiras do recorte. O [DER completo em SVG](diagramas/completo.svg) reúne todas as tabelas e views, com todas as colunas.

### Contas, contatos e conversas

[Contatos](catalogo/contacts.md) e [conversas](catalogo/conversations.md) pertencem a uma conta XMPP. Cada [mensagem](catalogo/messages.md) referencia tanto a conta quanto a conversa.

![Chaves estrangeiras entre contas, contatos, conversas, suas revisões e mensagens](diagramas/conversas.svg "DER de contas, contatos e conversas")

### Histórico e leitura

Os [checkpoints](catalogo/archive_checkpoints.md) registram o progresso da recuperação do histórico. As [revisões de leitura](catalogo/conversation_read_revisions.md) registram o estado que o motor associa à conversa e às mensagens lidas.

![Chaves estrangeiras dos checkpoints, identificadores de arquivo e registros de leitura](diagramas/historico.svg "DER de histórico e leitura")

### Envios e anexos

A [fila de envio](catalogo/outbox.md) inclui os envios agendados. Os [anexos em preparação](catalogo/staged_attachments.md) ficam vinculados à conversa até a composição do envio.

![Chaves estrangeiras entre conversas, mensagens, fila de envio e anexos em preparação](diagramas/envios.svg "DER de envios e anexos")

### Interações com mensagens

As tabelas deste recorte guardam [respostas](catalogo/message_replies.md), [estrelas](catalogo/message_stars.md), [comandos de edição](catalogo/message_edit_commands.md) e [comandos de reação](catalogo/message_reaction_commands.md).

![Chaves estrangeiras das respostas, estrelas, comandos de edição e reação, e revisões de mensagens](diagramas/interacoes.svg "DER de interações com mensagens")

### Conteúdo compartilhado

[Referências extraídas das mensagens](catalogo/shared_content.md) alimentam as abas de mídia, links e documentos. As [entradas ordenadas](catalogo/shared_catalog_order.md) registram a posição das referências em cada escopo e aba.

![Chaves estrangeiras das tabelas de conteúdo compartilhado, prévias e catálogo](diagramas/conteudo.svg "DER de conteúdo compartilhado")
