---
title: CLI
description: Instale ou compile o oiectl e consulte contas, conversas, mensagens e eventos do motor do Oiê.
language: pt-BR
---

# CLI do Oiê

O `oiectl` acessa um motor do Oiê em execução. Os exemplos usam uma instância local em `http://127.0.0.1:8080`. O [Oiê v0.0.1-rc.28](../../baixar/index.md) usa a API 0.36.0. A versão do CLI distribuída em separado acompanha o contrato 0.30.1.

## Instalar o CLI

Na release `cli/v0.30.1`, baixe o pacote para Linux x86-64 ou macOS Apple Silicon, confira o arquivo com `SHA256SUMS` e coloque o executável extraído no `PATH`. As notas também trazem o comando `go install` com a versão exata. [Baixar o oiectl v0.30.1](https://github.com/araihu/oie/releases/tag/cli/v0.30.1)

O `oiectl` v0.30.1 foi gerado a partir do [contrato OpenAPI 0.30.1 no catálogo da Arai Hû](https://oapi.araihu.com/catalogs/oie/documents/engine-v0.30.1/). Rode `oiectl --version` para conferir a versão do CLI e a tag do contrato.

```text
oiectl version v0.30.1 (spec openapi/engine/v0.30.1)
```

## Compilar e conectar

Use Make e Go 1.27.4 ou mais recente para compilar e iniciar o Oiê a partir do [código-fonte da v0.0.1-rc.28](https://github.com/araihu/oie/tree/4f46384c9084ff9098481fde658f9d9b55cc4367).

```sh
make build build-cli
bin/oie --role local --listen 127.0.0.1:8080 --data-dir ./data-cli
```

Os exemplos usam o CLI compilado com a rc.28. `bin/oiectl --version` mostra `oiectl version devel (OpenAPI 0.36.0)`. Se instalou o binário no `PATH`, use `oiectl` no lugar de `bin/oiectl`.

Em outro terminal, no mesmo checkout, confira a conexão com a API.

```sh
export OIECTL_SERVER=http://127.0.0.1:8080
bin/oiectl system health --output json
bin/oiectl system engine --output json
```

Em um perfil novo, abra [a interface local](http://127.0.0.1:8080) no navegador. [Adicione e conecte sua conta XMPP](../manual/index.md#contas-xmpp). Em **Nova conversa**, informe outro endereço XMPP e selecione **Abrir conversa**. Ela aparecerá em `conversations list`.

Para ver a conexão e o envio pela API em um ambiente de desenvolvimento, consulte o [exemplo com servidor XMPP local](https://github.com/araihu/oie/blob/4f46384c9084ff9098481fde658f9d9b55cc4367/README.md#connect-and-message-through-the-api). Para consultar um motor independente que já está em execução, ajuste `OIECTL_SERVER` ou use `--server` em um único comando. Para comandos que recebem JSON, use `--body-stdin` ou `--body-file`. Use um arquivo privado ou stdin para dados com senha.

Veja [Arquitetura e deploy](../arquitetura/index.md#motor-hospedado) para hospedar um motor e acessá-lo por uma conexão privada.

## Autenticar no motor hospedado

Use o `oiectl` compilado a partir da rc.28 ou posterior. No servidor, crie uma chave e guarde o arquivo
em local privado. Cada chave dá acesso completo à API da instância.

```sh
umask 077
oie auth key-create --data-dir /srv/oie/data --name automacao > oie.key
```

No cliente, forneça a chave pela variável `OIECTL_TOKEN` ao chamar o endereço
HTTPS do motor.

```sh
OIECTL_TOKEN="$(cat oie.key)" oiectl --server https://oie.exemplo.com system health
```

A chave é mostrada apenas na criação. Use `oie auth key-list` no servidor
para consultar o ID e a validade de cada chave. Para revogar, rode
`oie auth key-revoke --data-dir /srv/oie/data --id ID_DA_CHAVE`.


## Consultar e enviar mensagens

```sh
bin/oiectl accounts list --output json
bin/oiectl conversations list --output json
bin/oiectl messages search --q 'trilha' --output json
bin/oiectl messages search --q 'trilha' --include-edit-history --output json
```

Nas versões com retração, use `--include-retracted` em `messages search` ou `shared content list` para incluir conteúdo guardado de mensagens retraídas. Em `messages search`, combine a opção com `--include-edit-history` para incluir também as versões anteriores.

Repita `--account-ids` para pesquisar em mais de uma conta. Use os IDs de `accounts list` e reserve `--account-id` para a busca em uma única conta. `shared content list` também aceita o filtro de múltiplas contas.

```sh
bin/oiectl messages search --q 'trilha' \
  --account-ids PESSOAL_UUID --account-ids TRABALHO_UUID --output json
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

## Remover uma conversa da lista

Na rc.28, use o CLI compilado da mesma revisão do motor.
Em `conversations list --output json`,
confira o `id`, o `accountId` e o `revisionId` da sala. Substitua os valores
abaixo e gere um UUIDv7 novo para a solicitação.

```sh
printf '%s\n' '{"accountId":"ACCOUNT_UUID","expectedRevisionId":"REVISION_UUID"}' |
  oiectl conversations remove-room-from-chats --conversation-id CONVERSATION_UUID \
    --idempotency-key REQUEST_UUID --body-stdin
```

O comando tira a sala de **Conversas** e **Arquivadas** e interrompe a entrada
automática pelo Oiê. Mantém o histórico e a sala no servidor. Consulte o
[manual de salas](../manual/index.md#conferir-o-endereo-e-remover-salas)
para distinguir remoção local e exclusão para todos.

Para uma conversa direta, use `oiectl conversations remove-direct-from-chats`
com os mesmos parâmetros. O histórico fica salvo e a conversa volta ao
abrir o endereço novamente ou receber uma nova mensagem.

## Remover uma conta

Confira o ID da conta com `oiectl accounts list`. O comando abaixo remove a conta, a senha salva, as conversas e as mensagens pendentes do motor selecionado. As mensagens no servidor XMPP são mantidas. O CLI executa a remoção assim que você chama o comando.

```sh
oiectl accounts remove --account-id ID_DA_CONTA
```

Use o CLI compilado a partir do código-fonte da rc.23 e um motor com a API 0.34.0 ou posterior.

## Pesquisar mensagens favoritas

Com o CLI compilado da rc.23, consulte as mensagens favoritas pelo texto atual, remetente ou conversa de origem. Use `--account-id` ou `--conversation-id` para restringir a consulta.

```sh
bin/oiectl messages stars list --q 'trilha' --output json
```

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
| Salas e arquivo | `conversations discover-rooms`, `conversations open-room`, `conversations get-room-management`, `conversations rooms archive` |
| Buscar mensagens e consultar estrelas | `messages search`, `messages stars list` |
| Gerenciar envios pendentes | `messages list-scheduled-messages`, `messages update-pending-message`, `messages send-scheduled-message-now`, `messages outbox cancel` |
| Ver arquivos e links de conversas | `shared content list` |
| Consultar os prazos de retenção | `engine retention get`, `conversations retention get --conversation-id CONVERSATION_UUID` |
| Retomar o fluxo de eventos | `events stream` |

Com o CLI compilado da rc.9, `shared content list --include-first-page-counts` inclui a contagem da primeira página de mídia, links e documentos com os filtros usados na consulta.

Acrescente `--help` ao grupo ou comando para ver as opções da sua versão, por exemplo `bin/oiectl messages send --help`. Consulte o [contrato OpenAPI 0.30.1 da release do CLI](https://oapi.araihu.com/catalogs/oie/documents/engine-v0.30.1/), o [contrato 0.36.0 usado pela rc.28](https://github.com/araihu/oie/blob/4f46384c9084ff9098481fde658f9d9b55cc4367/api/openapi.yaml) ou a [prévia do contrato no catálogo](https://oapi.araihu.com/catalogs/oie/documents/engine-next/).
