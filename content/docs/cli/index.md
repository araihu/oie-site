---
title: CLI
description: Compile o oiectl e consulte contas, conversas, mensagens e eventos do motor do Oiê.
language: pt-BR
---

# CLI do Oiê

O `oiectl` acessa um motor do Oiê em execução. Os exemplos usam o CLI e o motor da mesma revisão de desenvolvimento e uma instância local em `http://127.0.0.1:8080`. O comando `oiectl --version` mostra a versão do contrato OpenAPI usada para gerar o CLI.

## Compilar e conectar

Use Make e Go 1.27.4 ou mais recente para compilar e iniciar o Oiê a partir do [repositório do aplicativo](https://github.com/araihu/oie/tree/7ae805764c4d019d85eff7880617939955979527).

```sh
make build build-cli
bin/oie --role local --listen 127.0.0.1:8080 --data-dir ./data-cli
```

Em outro terminal, no mesmo checkout, confira a conexão com a API.

```sh
export OIECTL_SERVER=http://127.0.0.1:8080
bin/oiectl system health --output json
bin/oiectl system engine --output json
```

Em um perfil novo, abra [a interface local](http://127.0.0.1:8080) no navegador. [Adicione e conecte sua conta XMPP](../manual/index.md#contas-xmpp). Em **Nova conversa**, informe outro endereço XMPP e selecione **Abrir conversa**. Ela aparecerá em `conversations list`.

Para configurar uma instância sem navegador, inicie o motor com `--data-only` e siga o [exemplo de conexão e envio pela API](https://github.com/araihu/oie/blob/7ae805764c4d019d85eff7880617939955979527/README.md#connect-and-message-through-the-api). Para consultar um motor independente que já está em execução, ajuste `OIECTL_SERVER` ou use `--server` em um único comando. Para comandos que recebem JSON, use `--body-stdin` ou `--body-file`. Use um arquivo privado ou stdin para dados com senha.

Veja [Arquitetura e deploy](../arquitetura/index.md#motor-hospedado) para hospedar um motor e acessá-lo por uma conexão privada.

## Consultar e enviar mensagens

```sh
bin/oiectl accounts list --output json
bin/oiectl conversations list --output json
bin/oiectl messages search --q 'trilha' --output json
bin/oiectl messages search --q 'trilha' --include-edit-history --output json
```

Para enviar, substitua `CONVERSATION_UUID` pelo ID retornado em `conversations list` e `REQUEST_UUID` por um UUIDv7 novo para essa mensagem.

```sh
printf '%s\n' '{"originalText":"Oiê! Tudo certo para sábado?"}' |
  bin/oiectl messages send --conversation-id CONVERSATION_UUID \
    --idempotency-key REQUEST_UUID --body-stdin --output json
```

A resposta confirma que o motor aceitou a mensagem. Substitua `MESSAGE_UUID` pelo ID retornado no envio para acompanhar o estado.

```sh
bin/oiectl messages get --message-id MESSAGE_UUID --output json
```

Se perder a resposta e precisar repetir a mesma solicitação, mantenha a chave e o corpo originais para evitar um envio duplicado.

## Eventos e saída

```sh
bin/oiectl events stream --timeout 10m
bin/oiectl events stream --timeout 10m --last-event-id ENGINE_UUID:SEQUENCE
```

Na retomada, use o ID do último evento que você processou por completo. Use `--output raw` para downloads binários e `--output json` para dados de automação. Os códigos de saída são `2` para erro de uso, `3` para falha de transporte ou tempo limite e `4` para erro HTTP. Os diagnósticos vão para stderr.

## Referência de comandos

| Para | Comece com |
| --- | --- |
| Contas, conexão e histórico | `accounts list`, `accounts connection`, `accounts history` |
| Salas e arquivo | `conversations preview-room`, `conversations open-room`, `conversations rooms archive` |
| Buscar mensagens e consultar estrelas | `messages search`, `messages stars list` |
| Gerenciar envios pendentes | `messages list-scheduled-messages`, `messages update-pending-message`, `messages send-scheduled-message-now`, `messages outbox cancel` |
| Ver arquivos e links de conversas | `shared content list` |
| Retomar o fluxo de eventos | `events stream` |

Acrescente `--help` ao grupo ou comando para ver as opções da sua versão, por exemplo `bin/oiectl messages send --help`. Consulte a [OpenAPI do Oiê no Manja](https://oapi.araihu.com/catalogs/oie/documents/engine-next/) para conferir operações, campos e respostas.
