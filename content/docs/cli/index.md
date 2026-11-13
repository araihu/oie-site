---
title: CLI
description: Compile o oiectl e consulte contas, conversas, mensagens e eventos do motor do Oiê.
language: pt-BR
---

# CLI do Oiê

O `oiectl` acessa um motor do Oiê em execução. Os exemplos abaixo usam a API em `http://127.0.0.1:8080`. A versão de `oiectl --version` identifica o contrato OpenAPI usado para gerar o CLI.

## Compilar e conectar

Com Go 1.27.3 ou mais recente, no checkout do [repositório do Oiê](https://github.com/araihu/oie), rode `make build-cli` para gerar `bin/oiectl`. Para iniciar um motor local independente da interface, compile também com `make build` e rode:

```sh
bin/oie --role local --data-only --listen 127.0.0.1:8080 --data-dir ./data-cli
```

Em outro terminal, no mesmo checkout:

```sh
export OIECTL_SERVER=http://127.0.0.1:8080
bin/oiectl system health --output json
bin/oiectl system engine --output json
bin/oiectl accounts list --output json
```

O diretório `./data-cli` começa com um perfil vazio. Para consultar conversas existentes, aponte o CLI para o motor que já contém essas conversas. Use `--server` para escolher outro motor em um único comando. Para comandos que recebem JSON, use `--body-stdin` ou `--body-file`. Use um arquivo privado ou stdin para dados com senha.

## Consultar e enviar mensagens

```sh
bin/oiectl conversations list --output json
bin/oiectl messages search --q 'trilha' --output json
bin/oiectl messages search --q 'trilha' --include-edit-history --output json
```

Para enviar, substitua `CONVERSATION_UUID` pelo ID retornado em `conversations list` e `REQUEST_UUID` por um UUIDv7 novo para essa mensagem:

```sh
printf '%s\n' '{"originalText":"Oiê! Tudo certo para sábado?"}' |
  bin/oiectl messages send --conversation-id CONVERSATION_UUID \
    --idempotency-key REQUEST_UUID --body-stdin --output json
```

A resposta confirma que o motor aceitou a mensagem. Consulte o estado do envio com `messages get`, usando o ID retornado. Se perder a resposta e precisar repetir a mesma solicitação, mantenha a chave e o corpo originais para evitar um envio duplicado.

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
