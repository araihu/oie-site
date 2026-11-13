---
title: Manual
description: Como configurar contas XMPP e usar salas, histórico e envios agendados no Oiê.
language: pt-BR
---

# Manual do Oiê

## Contas XMPP

Depois de [instalar o Oiê](../../baixar/index.md), abra **Contas** e adicione uma conta XMPP que você já usa. Informe o endereço completo, como `ana@exemplo.org`. Conecte a conta com a senha do seu servidor XMPP.

Use o seletor de contas para ver todas as conversas ou filtrar por uma conta. Ao iniciar uma conversa ou abrir um convite de sala, escolha a conta com que deseja participar. As conversas de cada conta são independentes.

![Seletor de contas do Oiê aberto, com as opções Todas as contas, Pessoal e Trabalho.](imagens/seletor-de-contas.webp)

## Salas, convites e entrada automática

Em **Nova conversa**, escolha a opção de entrar pelo endereço ou criar uma sala. Você pode colar um endereço de sala, como `trilha@salas.exemplo.org`, ou um convite no formato `xmpp:trilha@salas.exemplo.org?join`. Se houver uma ficha de sala na mensagem, **Ver sala** abre o formulário com o endereço preenchido.

A prévia pode mostrar nome, descrição, assunto, foto e número de participantes, conforme o que o serviço disponibiliza à conta selecionada. A entrada só acontece depois que você confirma. A criação de salas depende do suporte do serviço XMPP. A entrada por endereço aceita salas sem senha.

![Cartão de convite para a sala Trilhas de sábado com endereço, descrição e o link Ver sala.](imagens/convite-de-sala.webp)

Ao conectar a conta, o Oiê respeita os favoritos XMPP marcados para entrada automática. Usa a parte antes do `@` no endereço da conta como apelido inicial. Se você escolheu outro apelido ou saiu explicitamente de uma sala no Oiê, essa decisão tem prioridade.

## Histórico e busca em salas

As mensagens recentes de uma sala ficam disponíveis enquanto o Oiê busca o histórico mais antigo. Para priorizar a recuperação de uma sala, vá ao início da conversa e peça mensagens mais antigas.

A busca geral consulta as mensagens guardadas no Oiê. A busca dentro de uma sala também pode consultar o arquivo do servidor, quando ele oferece busca por texto. Ative **Incluir histórico de edições** na busca para encontrar versões anteriores de mensagens editadas. Os [atalhos de busca](../atalhos/index.md#busca-na-conversa) ajudam a percorrer os resultados.

![Busca por trilha em uma sala. O primeiro de três resultados está destacado na conversa e o campo oferece a opção de incluir o histórico de edições.](imagens/busca-na-sala.webp)

## Envio agendado

No menu junto ao botão de envio, escolha **Agendar mensagem** e confira a data, o horário e a mensagem. Em **Agendadas**, você pode editar, reagendar, cancelar ou pedir o envio imediato enquanto a mensagem ainda estiver aguardando o primeiro envio.

![Janela Agendar mensagem, com a prévia do lembrete e o campo de data e horário de envio.](imagens/envio-agendado.webp)

Se o aplicativo ou a conexão estiver indisponível no horário marcado, o Oiê tenta enviar quando voltar a funcionar e a conversa estiver disponível. Em salas, isso exige que a conta esteja na sala. Fechar a janela do aplicativo de desktop o mantém em execução. Para encerrá-lo, use [o atalho de saída](../atalhos/index.md#aplicativo-de-desktop).

## Citações e texto original

O Oiê exibe Markdown nas mensagens. Para citar várias linhas, comece cada uma com `>`. Uma linha sem o marcador fica fora da citação, mesmo quando vem logo abaixo. No menu da mensagem, use **Copiar mensagem original** para copiar o texto com os marcadores Markdown.

![Mensagem com a pergunta citada em um bloco e a resposta Sim! Levo o café abaixo, fora da citação.](imagens/citacao-markdown.webp)
