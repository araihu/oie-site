---
title: Manual
description: Como agendar mensagens, gerenciar o envio agendado, conectar contas XMPP e buscar o histórico das salas.
language: pt-BR
---

# Manual do Oiê

Este manual acompanha a versão de desenvolvimento do Oiê. A entrada em salas novas por endereço ou convite, a criação de salas e o envio agendado chegaram depois do [instalador v0.0.1-rc.1](../../baixar/index.md).

## Contas XMPP

No rodapé da lista de conversas, abra **Trocar de conta** e selecione **Adicionar conta**. Informe o endereço completo da sua conta XMPP, como `ana@exemplo.org`, dê um nome à conta e selecione **Salvar conta**. No painel da conta, use **Conectar** com a senha do servidor XMPP. Para voltar a esse painel depois, use **Trocar de conta → Gerenciar contas**.

Use **Trocar de conta** para ver todas as conversas ou filtrar por uma conta. Ao iniciar uma conversa, escolha a conta que você quer usar. As conversas de cada conta são independentes.

![Seletor de contas do Oiê aberto, com as opções Todas as contas, Pessoal e Trabalho.](imagens/seletor-de-contas.webp)

[Ampliar captura do seletor de contas](imagens/seletor-de-contas.webp)

## Salas, convites e entrada automática

Ao conectar a conta, o Oiê respeita os favoritos XMPP marcados para entrada automática. Usa a parte antes do `@` no endereço da conta como apelido inicial. Se você escolheu outro apelido ou saiu explicitamente de uma sala no Oiê, essa decisão tem prioridade. Para entrar em uma sala salva, abra a conversa e use **Entrar na sala**.

Na versão de desenvolvimento, abra **Nova conversa** e escolha **Entrar na sala** ou **Criar sala**. Selecione a conta e informe um endereço, como `trilha@salas.exemplo.org`. Para entrar, você também pode colar um convite no formato `xmpp:trilha@salas.exemplo.org?join`. Em um cartão de convite, **Ver sala** abre o formulário com o endereço preenchido.

A prévia pode mostrar nome, descrição, assunto, foto e número de participantes, conforme o que o serviço disponibiliza à conta selecionada. A entrada só acontece depois que você confirma. A criação de salas depende do suporte do serviço XMPP. A entrada por endereço aceita salas sem senha.

![Cartão de convite para a sala Trilhas de sábado com endereço, descrição e o link Ver sala.](imagens/convite-de-sala.webp)

[Ampliar captura do convite de sala](imagens/convite-de-sala.webp)

## Histórico e busca em salas

Em **Detalhes da conversa**, use **Baixar histórico antigo automaticamente** para escolher como o Oiê recupera o histórico daquela sala. Com a opção ligada, ele baixa mensagens antigas aos poucos. Com a opção desligada, pode buscá-las quando você usa **Mensagens anteriores** no topo da conversa ou faz uma busca, conforme o servidor permitir.

A busca geral consulta as mensagens guardadas no Oiê. Dentro de uma sala, abra **Mais opções** e selecione **Pesquisar mensagens**. Essa busca também pode consultar o arquivo do servidor, quando ele oferece busca por texto. Ative **Incluir histórico de edições** para encontrar versões anteriores de mensagens editadas. Use os [atalhos de busca](../atalhos/index.md#busca-na-conversa) para percorrer os resultados.

![Busca por trilha em uma sala. O primeiro de três resultados está destacado na conversa e o campo oferece a opção de incluir o histórico de edições.](imagens/busca-na-sala.webp)

[Ampliar captura da busca na sala](imagens/busca-na-sala.webp)

## Envio agendado

Na versão de desenvolvimento, abra o menu junto ao botão de envio, escolha **Agendar mensagem** e confira a data, o horário e a mensagem. Em **Agendadas**, você pode editar, reagendar, cancelar ou pedir o envio imediato enquanto a mensagem ainda estiver aguardando o primeiro envio.

![Janela Agendar mensagem, com a prévia do lembrete e o campo de data e horário de envio.](imagens/envio-agendado.webp)

[Ampliar captura do envio agendado](imagens/envio-agendado.webp)

Se o aplicativo ou a conexão estiver indisponível no horário marcado, o Oiê tenta enviar quando voltar a funcionar e a conversa estiver disponível. Em salas, isso exige que a conta esteja na sala. Fechar a janela do aplicativo de desktop o mantém em execução. Para encerrá-lo, use [o atalho de saída](../atalhos/index.md#aplicativo-de-desktop).

## Citações e texto original

O Oiê exibe Markdown nas mensagens por padrão. Para citar várias linhas, comece cada uma com `>`. Uma linha sem o marcador fica fora da citação, mesmo quando vem logo abaixo. Em **Configurações → Este dispositivo**, use **Formatar mensagens com Markdown** para ajustar a exibição. No menu da mensagem, use **Copiar mensagem original** para copiar o texto com os marcadores Markdown.

![Mensagem com a pergunta citada em um bloco e a resposta Sim! Levo o café abaixo, fora da citação.](imagens/citacao-markdown.webp)

[Ampliar captura da citação em Markdown](imagens/citacao-markdown.webp)
