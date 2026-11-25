---
title: Manual
description: Escolha o motor local ou remoto, conecte ou remova contas XMPP, gerencie contatos e presença, marque conversas como lidas ou não lidas, gerencie salas, busque e encaminhe mensagens, configure notificações, consulte conversas recentes sem conexão, envie mensagens e anexos, consulte o diagnóstico e gerencie a retenção e exclusão de dados.
language: pt-BR
margo:
  imageViewer: true
---

# Manual do Oiê

Este manual acompanha o [Oiê v0.0.1-rc.54](../../baixar/index.md). A navegação dos painéis e as ações de mídia descritas aqui serão disponibilizadas na próxima versão.

## Motor local ou remoto

No desktop, entre em **Configurações** e abra **Motor**. **Local** é o padrão e roda o motor no seu dispositivo. Para acessar um motor hospedado, selecione **Remoto**, informe seu endereço completo, como `http://127.0.0.1:18084`, e clique em **Salvar**. O Oiê verifica a conexão antes de abrir as conversas do motor remoto. A escolha vale para esse perfil do desktop e é lembrada quando você reabre o aplicativo.

![Tela Motor com o botão Voltar às configurações, a opção Remoto selecionada, o endereço preenchido e o botão Salvar.](imagens/motor-remoto.webp "Configurações do motor remoto") [Ver captura completa](imagens/motor-remoto-original.webp)

As contas e o histórico de cada motor ficam onde foram gravados. Para voltar às suas conversas deste dispositivo, inclusive se o motor remoto sair do ar, selecione **Local** e salve. O motor remoto continua funcionando no servidor.

Para um motor local de teste em outro computador, restrinja o acesso com um
[túnel SSH ou uma rede privada](../arquitetura/index.md#motor-hospedado).

### Conectar a um servidor OIDC

Use a rc.28 ou posterior no servidor e no desktop.

1. Abra o site do Oiê hospedado e selecione **Entrar**.
2. Vá a **Configurações → Chaves de acesso**. Dê um nome que
   identifique o dispositivo, selecione **Criar chave de acesso** e depois
   **Copiar chave**. A chave dá acesso completo e vale por 90 dias.
3. No desktop, abra **Configurações → Motor → Remoto**. Informe o
   endereço HTTPS, cole a chave em **Chave de acesso (se exigida)** e
   selecione **Salvar**.

O desktop verifica a conexão antes de salvar o endereço e a chave no perfil.
Se mudar o endereço depois de colar uma chave, cole a chave do endereço final.
Se sair da tela antes de salvar, cole a chave novamente ao voltar.
Para trocar uma chave vencida, crie outra no site do Oiê hospedado e cole-a na
mesma tela do desktop.

Se um motor remoto salvo passar a exigir
autenticação, o desktop mostra **Entrar pelo navegador** e **Inserir chave de
acesso**. Use a primeira opção para entrar no servidor e criar a chave e a
segunda para colá-la no desktop.

### Quando a conexão cai

Se o motor ficar indisponível, o Oiê tenta reabrir a conversa assim que a conexão voltar, mesmo se você abriu a janela durante a queda. Se você já estava escrevendo, o rascunho permanece no campo de mensagem. No desktop, em **Motor local**, o Oiê também tenta reiniciar o motor com o mesmo perfil se ele parar inesperadamente.

Se a conexão com o servidor XMPP cair, o Oiê tenta reconectar a conta e retomar a sessão quando o servidor permite. Os envios que já estavam na fila do motor são retomados quando a conexão volta. Se algum envio ficar sem confirmação, confira o status na conversa. Se a conta parar por causa da senha ou de outra configuração, corrija o problema em **Perfil** e selecione **Conectar**.

### Sair do Oiê hospedado

Abra **Configurações**, selecione **Sair** no rodapé e confirme. Isso encerra
a sessão em todas as abas deste navegador. O motor continua rodando no servidor.

## Diagnóstico e registro de erros

Abra **Configurações → Diagnóstico**.

![Configurações com aparência, idioma, fuso horário, idiomas dos atalhos de emoji e acesso às preferências do dispositivo e do motor.](imagens/configuracoes-geral.webp "Página inicial das configurações") [Ver captura completa](imagens/configuracoes-geral-original.webp)

Escolha um período e use **Atualizar** para ver os eventos recentes.
Abra uma linha para ver os detalhes. Use
**Copiar** para colocar o trecho na área de transferência. Para salvar o
arquivo `oie-diagnostics.jsonl`, abra a seta ao lado de **Copiar** e
escolha **Baixar**.

![Painel de Diagnóstico com um aviso de conexão expandido e o menu Baixar ao lado de Copiar.](imagens/diagnostico.webp "Visualizador de diagnóstico") [Ver captura completa](imagens/diagnostico-original.webp)

No aplicativo para computador, **Este dispositivo** continua disponível
se o motor falhar na inicialização ou perder a conexão. Quando o motor
remoto oferecer diagnóstico, selecione **Servidor** e informe a chave de
acesso fornecida por quem administra a instalação.

O registro inclui horários e detalhes de falhas. Mensagens, rascunhos,
nomes de contas, endereços, credenciais e caminhos de arquivos ficam
fora do registro. O arquivo só é compartilhado quando você decide enviá-lo.

## Apagar todos os dados

Abra **Configurações → Ações destrutivas → Apagar todos os dados**.

Leia o aviso e digite `delete` para confirmar. A exclusão apaga as contas, o histórico e os arquivos guardados pelo motor selecionado. Em um motor compartilhado, isso afeta todos os dispositivos conectados.

## Contas XMPP

No rodapé da lista de conversas, abra **Trocar de conta** e selecione **Adicionar conta**. Informe o endereço completo da sua conta XMPP, como `ana@exemplo.org`, dê um nome à conta e selecione **Salvar conta**. No painel da conta, use **Conectar** com a senha do servidor XMPP. Para voltar a esse painel depois, use **Trocar de conta → Gerenciar contas**.

Use **Trocar de conta** para ver todas as conversas ou filtrar por uma conta. Ao iniciar uma conversa, escolha a conta que você quer usar. As conversas de cada conta são independentes.

Em **Gerenciar contas**, abra o **Perfil** da conta e use **Editar nome** e **Salvar nome** para alterar o nome mostrado no Oiê. Com a conta conectada, use **Editar foto**, escolha uma imagem PNG, JPEG ou GIF de até 512 KB e confira a prévia antes de selecionar **Salvar foto**.

A borda dos avatares das conversas usa a cor da conta de origem, assim como o
avatar em **Trocar de conta**. O nome e o endereço identificam a conta por
escrito.

[![A mesma sala nas contas Alice e Bob, com bordas nas cores mostradas no seletor de contas.](imagens/seletor-de-contas.webp)](imagens/seletor-de-contas-original.webp)

### Contatos

Em **Trocar de conta → Gerenciar contas**, abra o **Perfil** da conta e escolha
**Contatos**. Com a conta conectada, informe o endereço XMPP da pessoa e
selecione **Adicionar contato**. Em **Solicitações de contato**, use **Aceitar**
ou **Recusar**.

A presença aparece na lista quando o contato compartilha a disponibilidade
com sua conta. Para tirar um endereço da lista, selecione **Remover** ao lado
dele. O Oiê mantém a conversa e as mensagens guardadas.

[![Painel Contatos da conta Pessoal, com o endereço XMPP de Bruno e os botões Adicionar contato e Remover.](imagens/contatos.webp)](imagens/contatos-original.webp)

### Disponibilidade e digitação

No **Perfil** da conta, abra **Disponibilidade**, escolha **Disponível**,
**Ausente** ou **Não perturbe** e selecione **Salvar**. Em **Não perturbe**,
as mensagens continuam chegando sem avisos ou sons para aquela conta.

Desative **Mostrar quando estou digitando** e selecione **Salvar** para
deixar de enviar o indicador nas conversas diretas compatíveis.

[![Preferências da conta com Não perturbe selecionado, a opção Mostrar quando estou digitando e o botão Salvar.](imagens/disponibilidade-conta.webp)](imagens/disponibilidade-conta-original.webp)

### Remover uma conta

Em **Trocar de conta → Gerenciar contas**, abra a conta e selecione **Remover conta** no rodapé dos detalhes. Confira o endereço e confirme em **Remover conta**. Escolha **Desconectar** se quiser apenas sair do servidor e manter as conversas no Oiê.

A remoção apaga a senha salva, as conversas e as mensagens pendentes dessa conta no Oiê. As outras contas e as mensagens no servidor XMPP são mantidas. Se você usa um motor remoto, a conta também some dos outros dispositivos conectados a ele.

[![Confirmação de remoção da conta Pessoal no Oiê, com as opções Cancelar e Remover conta.](imagens/remover-conta.webp "Confirmação de remoção")](imagens/remover-conta-original.webp)

## Percorrer a lista de conversas

Quando a lista tiver mais de 100 conversas, selecione **Conversas mais antigas** no fim da lista para abrir a próxima página. Use **Conversas mais recentes** para voltar à primeira página. A conversa aberta e os filtros acompanham a navegação. No celular, **Voltar para as conversas** retorna à lista na posição em que você estava, mesmo se a conexão cair.

[![Fim da lista de conversas com a opção Conversas mais antigas.](imagens/paginacao-conversas.webp)](imagens/paginacao-conversas.webp)

## Solicitações de mensagem e convites recebidos

A primeira mensagem de uma conversa direta sem histórico anterior pode aparecer como **Solicitação de mensagem**. Confira o endereço no cabeçalho. Use **Aceitar** para ler e responder, ou **Arquivar** para tirar a solicitação da lista principal. O conteúdo fica oculto até você aceitar.

![Solicitação de mensagem arquivada com o endereço do remetente no cabeçalho e o botão Aceitar.](imagens/solicitacao-de-mensagem.webp "Solicitação de mensagem arquivada")

Uma solicitação arquivada continua em **Arquivadas**, abaixo dos filtros da lista de conversas, mesmo se chegarem mais mensagens. Selecione **Arquivadas** novamente para voltar à lista principal. Se você aceitar uma solicitação de mensagem por lá, a conversa volta à lista principal. Em **Convite para sala**, use **Entrar na sala** para pedir a entrada. O conteúdo aparece quando o servidor XMPP confirma. Você também pode arquivar o convite e decidir depois.

## Remover uma conversa direta

Em **Mais opções → Remover das conversas**, confirme a remoção para tirar
a conversa de **Conversas** e **Arquivadas**. O histórico fica salvo.
Abra o endereço novamente com a mesma conta, ou receba uma nova mensagem
do contato, para a conversa voltar à lista.

Nas salas e nas conversas diretas removidas, os links da busca, de **Favoritas**
e de **Mídia e links** abrem o histórico salvo em modo de leitura. A conversa
continua fora da lista.

Para ver **Detalhes da conversa**, clique na foto ou nas iniciais do cabeçalho.

## Salas, convites e entrada automática

Ao conectar a conta, o Oiê respeita os favoritos XMPP marcados para entrada automática. Usa a parte antes do `@` no endereço da conta como apelido inicial. Se você escolheu outro apelido ou saiu explicitamente de uma sala no Oiê, essa decisão tem prioridade. Para entrar em uma sala salva, abra a conversa e use **Entrar na sala**.

Abra o menu **Nova conversa** e escolha **Conversa direta**, **Entrar na sala** ou **Criar sala**. Para entrar ou criar uma sala, siga o formulário da opção escolhida. Selecione a conta e informe um endereço, como `trilha@salas.exemplo.org`. Para entrar, você também pode colar um convite no formato `xmpp:trilha@salas.exemplo.org?join`. Em um cartão de convite, **Ver sala** abre o formulário com o endereço preenchido.

A prévia pode mostrar nome, descrição, assunto, foto e número de participantes, conforme o que o serviço disponibiliza à conta selecionada. A entrada só acontece depois que você confirma. A criação de salas depende do suporte do serviço XMPP. Se a sala exigir senha, informe-a no formulário de entrada. Se a conexão cair depois que você confirmar, tente novamente. Ao reabrir o Oiê, ele verifica se o pedido foi aceito e só pede a senha outra vez se precisar reenviá-lo.

![Cartão de convite para a sala Trilhas de sábado com endereço, descrição e o link Ver sala.](imagens/convite-de-sala.webp "Cartão de convite para uma sala")

### Descobrir e gerenciar salas

- Em **Nova conversa → Entrar na sala**, abra **Explorar salas públicas** e selecione uma sala para preencher o endereço e ver a prévia. Confirme quando quiser entrar.
- Selecione a imagem da sala no cabeçalho da conversa e abra **Gerenciar sala**, logo abaixo do nome. A seta à esquerda de **Gerenciar sala** leva de volta aos detalhes da conversa. O **×** fecha o painel. Em **Detalhes**, escolha **Editar nome e tópico**. Altere os campos, selecione **Revisar alterações** e confira os valores antigos e novos antes de **Publicar alterações**. As opções dependem das permissões da sua conta. A descrição aparece quando o serviço a fornece. Use **Mais opções → Convidar pessoas** para enviar um convite quando a sala permitir.
- Em **Participantes**, busque pelo apelido e abra o nome da pessoa para ver as ações disponíveis. Dependendo das permissões, você pode alterar a função, remover a pessoa ou ajustar o acesso à sala. Em **Banimentos**, administradores e proprietários podem buscar um banimento, banir por endereço ou retirar um banimento.
- Se o servidor oferecer moderação e sua conta tiver permissão, use **Remover mensagem** no menu da mensagem. A remoção aparece quando a sala a confirma.

[![Gerenciar sala com o botão Editar nome e tópico e os valores atuais.](imagens/gerenciar-sala.webp)](imagens/gerenciar-sala-original.webp)

### Conferir o endereço e remover salas

Ao criar uma sala, use o domínio do serviço de salas XMPP. Por exemplo, se a
conta for `ana@exemplo.org`, o endereço da sala pode ser
`trilha@salas.exemplo.org`. Confira com quem administra o servidor qual
domínio usar. Se o endereço não puder ser confirmado, o Oiê mantém o
formulário aberto para você corrigir ou tentar de novo.

- **Mais opções → Remover das conversas** tira a sala de **Conversas** e
  **Arquivadas** e interrompe a entrada automática pelo Oiê. Isso funciona
  também para uma sala que não conectou. A sala e o histórico salvo continuam
  disponíveis. Para trazê-la de volta, use **Entrar na sala** e informe o endereço.
- Se você for proprietário, entre na sala e vá a **Mais opções → Gerenciar
  sala → Detalhes → Excluir sala para todos**. Confirme para excluir a sala
  do servidor e remover todos os participantes. Quando o servidor confirmar,
  a sala sai de **Conversas** e **Arquivadas** da conta que pediu a exclusão.
  O histórico salvo será mantido. Quem recebe o aviso de exclusão vê **Sala
  excluída** e pode remover a sala da própria lista.

[![Confirmação para remover uma sala das conversas, com as opções Cancelar e Remover das conversas.](imagens/remover-sala.webp)](imagens/remover-sala-original.webp)

[![Confirmação para excluir a sala para todos, com aviso de que a ação não pode ser desfeita.](imagens/excluir-sala.webp)](imagens/excluir-sala-original.webp)

### Ver o estado de uma sala

O ponto no avatar mostra o estado da sala. O cabeçalho exibe o estado por
escrito, e a lista destaca quando você está fora da sala ou quando há um
problema.

- **Na sala** indica que você já entrou.
- **Conta desconectada** aparece quando sua conta está sem conexão.
- **Sala excluída** aparece quando o servidor confirma a exclusão. O histórico
  que você já tem continua acessível.
- **Acesso restrito** indica que sua entrada foi recusada. **Sala
  indisponível** indica que o Oiê não conseguiu acessar a sala.

[![Uma sala na lista de conversas com um ponto vermelho no avatar e o rótulo Sala excluída.](imagens/estado-sala-excluida.webp)](imagens/estado-sala-excluida-original.webp)

## Histórico e busca em salas

Para abrir **Detalhes da conversa**, clique na foto ou nas iniciais no cabeçalho da conversa, ou abra **Mais opções → Detalhes da conversa**. Use **Baixar histórico antigo automaticamente** para escolher como o Oiê recupera o histórico daquela sala. Com a opção ligada, ele baixa mensagens antigas aos poucos. Com a opção desligada, pode buscá-las quando você usa **Mensagens anteriores** no topo da conversa ou faz uma busca, conforme o servidor permitir.

Dentro de uma sala, abra **Mais opções** e selecione **Pesquisar mensagens**. Essa busca também pode consultar o arquivo do servidor, quando ele oferece busca por texto. Ative **Incluir histórico de edições** para encontrar versões anteriores de mensagens editadas. Use os [atalhos de busca](../atalhos/index.md#busca-na-conversa) para percorrer os resultados.

Buscar por `arqiuvo` nas mensagens guardadas pode realçar **arquivo** no balão. Se o resultado vier de uma versão anterior ou dos detalhes de um link, o realce aparece no trecho do resultado. Ao recarregar a página com a busca aberta, o Oiê volta à versão anterior selecionada.

![Busca por trilha em uma sala. O primeiro de três resultados está destacado na conversa e o campo oferece a opção de incluir o histórico de edições.](imagens/busca-na-sala.webp "Busca no histórico da sala")

## Retenção de mensagens

Abra **Configurações → Motor → Retenção de mensagens**. O Oiê salva os prazos padrão de **Salas** e **Conversas diretas** ao sair do campo alterado.

Os prazos padrão valem para o motor do Oiê em uso, inclusive quando ele é compartilhado entre dispositivos. Para mudar o prazo de uma conversa, abra **Mais opções → Detalhes da conversa → Preferências da conversa → Retenção de mensagens**. A escolha é salva ao selecionar. Para um prazo personalizado, digite os dias e pressione Enter ou saia do campo. Use **Usar o padrão** para voltar ao prazo global.

Por padrão, as salas têm prazo de **180 dias** e as conversas diretas usam **Sem limite**. O Oiê preserva as **100 mensagens mais recentes já guardadas de cada sala** e todas as mensagens favoritadas. Reduzir o prazo pode apagar outras mensagens antigas do histórico guardado pelo motor. Em um envio agendado, o prazo conta a partir do envio. Se a mensagem original de uma resposta expirar, o texto da resposta continua visível e a citação fica indisponível.

O prazo também limita a recuperação do histórico no XMPP. Uma instalação nova pode recuperar menos de 100 mensagens de uma sala, dependendo do prazo e dos dados ainda disponíveis no servidor. Com **Sem limite**, as mensagens deixam de expirar por idade, o histórico pode ocupar mais espaço e a recuperação pode alcançar as mensagens mais antigas que o servidor ainda oferece à sua conta.

![Página Motor com 180 dias para salas e Sem limite para conversas diretas.](imagens/retencao-de-mensagens.webp "Retenção de mensagens no Motor") [Ver captura completa](imagens/retencao-de-mensagens-original.webp)

## Mensagens não lidas

Ao selecionar uma conversa na lista, o Oiê abre na primeira mensagem não lida que estiver no histórico. Se todas estiverem lidas, mostra as mensagens mais recentes. Resultados de busca e respostas citadas continuam abrindo a mensagem selecionada.

Depois de percorrer o histórico, role até o fim ou use o botão com seta para baixo **Ir para as mais recentes**, acima do campo de mensagem. Quando há mensagens não lidas, o botão mostra o contador da conversa e exibe **+99** para valores maiores. Ao permanecer no fim com a janela ou aba em foco e sem buscas ou diálogos abertos, o Oiê marca como lidas também as mensagens anteriores que ficaram fora da tela. O contador de não lidas e o indicador de menções são atualizados na lista de conversas.

### Marcar uma conversa como lida ou não lida

No cabeçalho da conversa, abra **Mais opções → Marcar como lida** para zerar
o contador e limpar qualquer lembrete, inclusive se você estiver no meio do
histórico.

Use **Mais opções → Marcar como não lida** para incluir a conversa no filtro
**Não lidas**. Se todas as mensagens já foram lidas, a lista mostra um ponto
no lugar do contador. Para tirar o lembrete, escolha **Marcar como lida**.

[![Menu Mais opções de uma conversa com a ação Marcar como não lida.](imagens/marcar-conversa-como-nao-lida.webp)](imagens/marcar-conversa-como-nao-lida-original.webp)

## Notificações de mensagens

Em **Configurações → Este dispositivo → Notificações**, ative **Notificações de mensagens** para receber banners e conceda a permissão quando o navegador ou o sistema pedir.

O som vem ativado para mensagens diretas e, nas salas, para menções e respostas a mensagens suas. Em **Som nas salas**, você pode escolher **Todas as mensagens**, manter **Menções e respostas a você** ou selecionar **Desativado**. Há controles para mensagens diretas, salas arquivadas e a conversa em foco. **Som de notificação** desliga todos os toques.

No iPhone e no iPad, você pode abrir o Oiê, entrar em uma conversa e digitar enquanto ouve música em outro aplicativo. Se começar a gravar uma mensagem de voz enquanto uma notificação toca, o Oiê interrompe o toque.

Use **Escolher arquivo** em **Toque** para usar um áudio seu de até 5 MB. O Oiê usa os primeiros dois segundos e guarda o toque neste dispositivo. Use o botão de reprodução para ouvi-lo ou o **X** ao lado do nome do arquivo para voltar ao toque do Oiê.

![A seção Notificações mostra o arquivo Toque de exemplo.wav com o botão de reprodução, o X e a opção Escolher arquivo.](imagens/notificacoes.webp "Toque de notificação personalizado") [Ver captura completa](imagens/notificacoes-original.webp)

Para silenciar uma conversa ou sala neste dispositivo, abra **Detalhes da conversa → Preferências da conversa** e ative **Silenciar notificações neste dispositivo**. Ao receber várias mensagens em sequência, o Oiê toca uma vez a cada dois segundos. O contador de mensagens não lidas continua funcionando.

Os alertas chegam enquanto o Oiê está aberto e conectado. Selecione o banner para abrir a mensagem. No desktop, o clique restaura a janela se ela estiver escondida. No navegador ou no app web instalado, um banner que ainda estiver na central de notificações pode reabrir o Oiê após fechar a janela. Se o navegador bloquear áudio ao abrir o Oiê, interaja com a janela para habilitar a reprodução.

## Busca geral e filtros por conta

Digite no campo **Pesquisar**, na lista de conversas. **Tudo** reúne conversas, contatos salvos e mensagens guardadas. Use **Conversas** ou **Mensagens** para restringir os resultados. Em **Contas**, marque uma ou mais contas ou limpe a seleção para pesquisar em todas.

Os resultados destacam os termos encontrados e mostram a conta e a conversa de origem. Abra um resultado para ir à mensagem. Se ativar **Incluir histórico de edições**, um resultado marcado como **Versão anterior** mostra o trecho guardado e abre a mensagem atual na conversa.

![Busca por plataformas com duas contas selecionadas e o termo destacado no texto e no link da mensagem.](imagens/busca-geral.webp "Busca em várias contas")

Os atalhos **Mídia**, **Links** e **Documentos** abrem o catálogo com o mesmo texto de busca e as contas selecionadas. Para uma consulta rápida em todas as contas, use [Ctrl + K ou ⌘ + K](../atalhos/index.md#busca-geral). A consulta dessa janela é independente da busca da lista.

## Mídia e links

Abra **Mídia e links** na barra lateral para encontrar o que foi compartilhado. As abas **Mídia**, **Links** e **Documentos** agrupam os itens por mês. Use **Contas** para buscar por nome ou endereço e marcar uma ou mais contas. **Limpar seleção** volta a mostrar itens de todas as contas. Sua escolha em **Trocar de conta** continua valendo nas outras telas.

No menu do cartão, use **Carregar imagem** quando a prévia ainda não estiver visível ou **Ir para a mensagem** para voltar à conversa de origem. Selecione um documento para abri-lo. Se a prévia falhar, use **Tentar novamente** quando essa opção aparecer. O cabeçalho também permite baixar o arquivo ou abrir o endereço original.

Na conversa, se a imagem ainda não apareceu, abra **Ações da mensagem** e escolha **Carregar imagem**. Se houver várias imagens na mensagem, escolha o arquivo pelo nome ou pelo número.

### Reprodução de áudio

Na aba **Mídia**, selecione **Reproduzir áudio** no cartão. Quando a forma de onda aparece, selecione um ponto nela para avançar ou voltar. A reprodução pelo cartão está disponível desde a rc.52.

Abra as ações do cartão para escolher **Baixar áudio** ou, após uma falha de reprodução, **Tentar novamente**. Para os controles da conversa, consulte [Mensagens de voz](#mensagens-de-voz).

### Arquivos baixados

Com o motor local, abra **Configurações → Este dispositivo → Armazenamento → Arquivos baixados**. Em **Manter arquivos por**, a escolha é salva ao selecionar uma opção. O padrão é **Até eu excluir**. Para liberar espaço agora, selecione **Excluir arquivos baixados** e confirme.

![Preferências de arquivos baixados com a opção Até eu excluir e o botão Excluir arquivos baixados.](imagens/arquivos-baixados.webp "Retenção de arquivos baixados") [Ver captura dos arquivos baixados](imagens/arquivos-baixados.webp)

Quando um anexo compatível inclui dados de verificação, o Oiê confere a integridade do download completo. Se o conteúdo divergir do original anunciado, o Oiê recusa o download.

## Mensagens favoritadas

No menu de uma mensagem, escolha **Favoritar**. Abra **Favoritas** na navegação lateral e use **Pesquisar mensagens favoritas** para buscar pelo texto, pelo remetente ou pela conversa de origem. Os cartões incluem mensagens de conversas arquivadas. Use **Trocar de conta** para restringir a lista a uma conta ou **Todas as contas** para reunir as mensagens de todas elas.

No menu do cartão, use **Ir para a mensagem** para abrir a conversa no ponto em que ela foi enviada. Você também pode **Copiar mensagem original** ou **Encaminhar**. Para tirar uma mensagem da lista, selecione a estrela. Para localizar conversas favoritas, use o filtro **Favoritas** na lista de conversas.

![Grade de mensagens favoritadas, com o remetente, a conversa e o balão da mensagem em cada cartão.](imagens/favoritas.webp "Mensagens favoritadas")

### Favoritar em Mídia e links

Em **Mídia e links**, clique na estrela no canto superior esquerdo do
cartão para favoritar a mensagem de origem. A estrela cinza aparece ao
passar o mouse ou focar o cartão pelo teclado. Nos favoritos, ela fica
amarela e visível. Em telas de toque, o botão fica sempre disponível.
Itens da mesma mensagem compartilham a estrela.

![Estrela cinza no primeiro cartão e estrela amarela no terceiro, ambas no canto superior esquerdo.](imagens/favoritar-midia.webp "Favoritar em Mídia e links")

## Editar e retrair mensagens enviadas

Na última mensagem própria enviada, abra **Ações da mensagem → Editar mensagem**. Com o campo de mensagem vazio, você também pode pressionar **↑**. Quando a edição está disponível, o Oiê carrega o texto no campo de mensagem. Depois de salvar, o ícone de edição ao lado do horário abre o **Histórico da mensagem**.

**Editar mensagem** também pode aparecer em respostas com citação enviadas na sessão atual. O Oiê conserva a citação quando você salva o novo texto.

Para retrair, abra **Ações da mensagem → Retrair mensagem** em uma mensagem própria e confirme. A ação depende do suporte do servidor. Em salas, também depende de o Oiê conseguir verificar a autoria. Outros participantes ainda podem ter uma cópia.

No histórico, **Mensagem retraída** ocupa o lugar do conteúdo. Se o Oiê ainda guardou a mensagem, use o ícone **Revelar original** ao lado do horário. O **Histórico da mensagem** mostra a retração como estado atual, seguida das versões guardadas. A busca geral, a busca na conversa e **Mídia e links** ocultam esse conteúdo por padrão. Selecione **Incluir retraídas** para encontrá-lo. Para buscar versões anteriores de mensagens retraídas, ative também **Incluir histórico de edições**.

![Histórico da mensagem com a retração no topo e as versões guardadas abaixo.](imagens/revelar-original.webp "Revelar uma mensagem retraída")

[Abrir a captura da retração em tamanho original](imagens/revelar-original.webp)

## Encaminhar mensagens

Abra **Encaminhar** no menu de uma mensagem em **Conversas**, **Favoritas** ou **Mídia e links**. Confira a prévia, escolha uma ou mais conversas diretas ou salas e confirme em **Encaminhar**. A lista mostra a conta que enviará cada cópia.

O Oiê prepara o texto e os anexos para cada destino. Se algum envio falhar ou ficar sem confirmação, veja o aviso junto ao destino e use **Tentar novamente**. Ao tentar novamente, o Oiê pula os destinos cuja cópia já entrou na fila.

![Janela Encaminhar mensagem com uma conversa direta selecionada e a conta de origem ao lado de cada destino.](imagens/encaminhar-mensagem.webp "Escolher destinos para encaminhar")

[Abrir a captura do encaminhamento em tamanho original](imagens/encaminhar-mensagem.webp)

## Anexos no rascunho

Abra **Mais opções** na seta ao lado de **Enviar** e escolha **Imagem** ou **Arquivo**. No navegador, você também pode colar um arquivo no campo ou arrastá-lo para a conversa. Escreva a legenda no próprio campo de mensagem.

![Menu ao lado de Enviar com as opções Imagem, Arquivo e Agendar mensagem.](imagens/menu-de-anexos.webp "Menu de envio")

[Abrir a captura do menu de envio em tamanho original](imagens/menu-de-anexos.webp)

Você pode clicar no nome de um anexo pronto para baixar o original. No aplicativo de desktop, isso abre a janela do sistema para salvar o arquivo. Use **Remover anexo** para tirar um item do rascunho. Se o preparo falhar, use **Tentar novamente** ou remova o arquivo. Se usou o seletor do Flatpak, escolha o arquivo novamente. Quando todos estiverem prontos, envie a mensagem ou [agende o envio](#envio-agendado).

![Rascunho com uma imagem, dois arquivos de texto e a legenda no campo de mensagem. Cada anexo tem um botão de remoção.](imagens/rascunho-anexos.webp "Anexos no rascunho")

O Oiê aceita até 16 arquivos por mensagem. O limite por arquivo depende do serviço XMPP e aparece no compositor quando está disponível. O serviço também pode restringir anexos naquela conversa. Quem recebe a mensagem pode baixar os arquivos pelos links enviados junto com a legenda. Os anexos prontos ficam no rascunho ao recarregar a página. Aguarde o preparo antes de trocar de conversa.

### Qualidade de áudio e vídeo

No menu de anexos, escolha **Áudio ou vídeo**, selecione o arquivo e escolha a qualidade. **SD** reduz o tamanho e vem selecionada em um dispositivo novo. **HD** preserva mais detalhes. **OG** mantém o original quando o formato é compatível e o tamanho cabe no limite informado pelo servidor. Em **Configurações → Este dispositivo → Mídia e exibição → Qualidade da mídia**, você pode mudar a preferência.

Em MP3 com capa PNG ou JPEG, **SD** prepara um M4A com a capa e os metadados. Se a economia for menor que 10% e o original couber no limite, o Oiê mantém o MP3. **HD** e **OG** também mantêm o MP3 original quando ele cabe.

O Oiê reconhece um MP3 mesmo com a extensão trocada e aceita vídeos MP4 com capa. Quando uma capa tem descrição própria ou identifica a contracapa, escolha **OG** para preservar essas informações. Para manter a transparência de um vídeo, escolha **Anexar como arquivo**. Se o Oiê avisar que um MP4 ou WAV parece incompleto, procure uma cópia íntegra antes de enviá-lo.

Escolha **Adicionar à mensagem** para preparar o anexo no dispositivo. Depois você pode enviar ou agendar a mensagem. No Flatpak, o arquivo selecionado fica no dispositivo até você confirmar a qualidade. O vídeo mantém as dimensões e a taxa de quadros. Se a conversão afetar uma faixa de áudio ou uma legenda, escolha **Anexar como arquivo** para anexar o original. Quando o servidor informa o tamanho máximo, o Oiê mostra o limite ao escolher a qualidade. Ao reabrir um rascunho com vídeo, você pode dar play mesmo que o vídeo seja grande demais para gerar a miniatura.

![Janela de qualidade para um MP3, com HD selecionada e as opções SD e OG disponíveis.](imagens/qualidade-de-midia.webp) [Ampliar a captura](imagens/qualidade-de-midia.webp)

## Mensagens de voz

Com o rascunho sem texto nem anexos, selecione **Gravar áudio** e permita o acesso ao microfone. O botão aparece nas conversas que aceitam anexos, em navegadores e aplicativos compatíveis com a gravação.

Enquanto grava, o campo mostra o tempo e a forma de onda. Selecione **Parar gravação** e use **Reproduzir gravação** para ouvir. Na versão com preparo de mídia, escolha também a qualidade depois de parar a gravação. Para refazer, use **Cancelar** no rascunho de áudio e selecione o microfone novamente. Quando estiver pronto, selecione **Enviar mensagem**.

![Campo de mensagem durante a gravação, com o tempo, a forma de onda e o botão de parar.](imagens/gravacao-mensagem-de-voz.webp "Gravação em andamento") [Ver captura da gravação](imagens/gravacao-mensagem-de-voz.webp)

![Gravação pronta para ouvir, descartar ou enviar.](imagens/rascunho-mensagem-de-voz.webp "Prévia da gravação") [Ver captura da prévia](imagens/rascunho-mensagem-de-voz.webp)

Para ouvir um áudio da conversa, use **Reproduzir áudio** no balão. Quando disponível, a forma de onda aparece com a duração total. Selecione um ponto nela para avançar ou voltar. Para salvar o original, abra o menu da mensagem e selecione **Baixar áudio**.

![Balão de mensagem com botão de reprodução, forma de onda e duração total.](imagens/onda-de-audio-no-balao.webp "Player de áudio no balão") [Ver captura do player](imagens/onda-de-audio-no-balao.webp)

## Prévia e reprodução de vídeos

Para anexar um vídeo, escolha **Arquivo** no [menu de anexos](#anexos-no-rascunho) e selecione o vídeo.
Use **Reproduzir vídeo** na miniatura para assistir antes de enviar. O nome
do arquivo e a opção **Remover anexo** continuam disponíveis se a miniatura
não puder ser criada.

![Vídeo no rascunho com botão de reprodução, nome do arquivo e botão para remover o anexo.](imagens/video-no-rascunho.webp "Vídeo no rascunho")

Em uma conversa direta ou sala, escolha **Reproduzir vídeo** no cartão para
assistir. Se a reprodução falhar, use o mesmo botão para tentar de novo.
Para baixar o arquivo, abra **Ações da mensagem** e escolha **Baixar original**. Em **Configurações → Este dispositivo → Mídia e exibição**, ajuste **Gerar prévias de
vídeo automaticamente**. Com essa opção desativada, você pode pedir uma prévia
em **Ações da mensagem → Carregar prévia do vídeo**. Na aba **Mídia**, abra as ações do cartão para carregar a prévia ou escolher **Baixar original**.

Ao tocar em **Reproduzir vídeo**, o Oiê tenta abrir o vídeo no navegador. Se o formato precisar de outro reprodutor, o Oiê o carrega no próprio dispositivo. A primeira reprodução pode demorar um pouco mais. Se o vídeo travar, use **Baixar original** no menu do cartão ou da mensagem para abrir o arquivo em outro aplicativo.

![Cartão de vídeo recebido em uma sala, com uma imagem estática e o botão de reprodução ao centro.](imagens/video-recebido.webp "Vídeo recebido em uma sala")

## Enviar sem conexão

Disponível desde a rc.34.

Se a conversa já estiver aberta quando a conexão cair, selecione **Enviar**
para guardar a mensagem neste navegador. O ícone de sinal cortado ao lado do
horário indica **Aguardando conexão**. O Oiê tenta enviá-la quando a
conexão voltar e o aplicativo estiver aberto neste navegador. Você pode
reabri-lo pela lista de conversas.

Enquanto a mensagem estiver aguardando conexão, selecione **Cancelar envio**
abaixo dela ou no menu **Ações da mensagem**. A mensagem volta ao rascunho,
onde você pode editar e enviar de novo.

Se outro rascunho estiver aberto, selecione **Envio pendente** para voltar à
mensagem que aguarda conexão. **Outro rascunho** alterna entre os rascunhos
guardados na conversa.

[![Mensagem com um ícone de sinal cortado, um aviso de conexão e a ação Cancelar envio.](imagens/envio-sem-conexao.webp)](imagens/envio-sem-conexao-original.webp)

Se aparecer **Verificar e tentar novamente**, use essa ação na mensagem. Se
o navegador não conseguir guardar a mensagem, o texto continua disponível
no campo de edição. Uma mensagem deixada no rascunho só é enviada quando
você seleciona **Enviar**.

## Ler conversas sem conexão

Disponível desde a rc.41.

No Oiê local, abra no navegador as conversas que quer consultar. Se perder
acesso ao Oiê, volte ao mesmo endereço no mesmo navegador. Na lista de
conversas salvas, selecione a conversa para ler as mensagens de texto guardadas.
**Tentar novamente** volta à conversa normal quando a conexão retornar. Seu
rascunho reaparece no campo de mensagem.

[![Conversa Casa aberta no navegador com o aviso Sem conexão.](imagens/historico-sem-conexao.webp)](imagens/historico-sem-conexao-original.webp)

O navegador guarda até 12 conversas, com até 50 mensagens de texto recentes
por conversa dentro do espaço disponível. Conversas recentes são guardadas
por até sete dias. Anexos e edição de rascunhos ficam disponíveis ao reconectar.
O recurso está disponível no Oiê local aberto em localhost ou HTTPS.

### Manter uma conversa disponível sem conexão

Disponível desde a rc.45.

Com a conversa aberta, selecione **Detalhes da conversa** e ative
**Manter disponível sem conexão**. O Oiê protege as mensagens salvas
do limite de sete dias e da remoção automática para liberar espaço.
A opção vale para este navegador.

[![Controle Manter disponível sem conexão ativado nos detalhes da conversa.](imagens/manter-disponivel-sem-conexao.webp)](imagens/manter-disponivel-sem-conexao.webp)

Se faltar espaço, desative a opção em outra conversa e tente novamente.
Se uma conversa protegida crescer e as mensagens mais recentes não couberem,
o Oiê avisa que a cópia está desatualizada ou indisponível. Depois de liberar
espaço, reabra a conversa com conexão para atualizar a cópia.

### Ajustar o espaço para conversas salvas

Disponível desde a rc.46.

Em **Configurações → Este dispositivo → Armazenamento → Armazenamento sem conexão**, confira o
espaço ocupado pelo texto das conversas e escolha o **Limite para conversas
salvas**. O padrão é 1 MiB. Você também pode escolher 256 KiB, 512 KiB ou
2 MiB. A escolha vale para este navegador.

[![Limite de 2 MiB selecionado em Armazenamento sem conexão, com o espaço ocupado pelas conversas salvas.](imagens/limite-conversas-salvas.webp)](imagens/limite-conversas-salvas.webp)

Ao reduzir o limite, o Oiê libera espaço removendo primeiro as cópias mais
antigas sem proteção. Se as conversas protegidas já ultrapassarem o novo
limite, o Oiê mantém o anterior e informa quanto espaço liberar. Nesse caso,
desative **Manter disponível sem conexão** em uma ou mais conversas e tente
novamente. Rascunhos, mensagens aguardando envio e arquivos baixados ficam
guardados separadamente.

### Pesquisar mensagens sem conexão

Disponível desde a rc.45.

Na lista de conversas salvas, use **Pesquisar mensagens salvas neste dispositivo**
para procurar uma palavra ou frase. Para procurar em uma única conversa, abra-a
e use **Pesquisar mensagens salvas nesta conversa**. Selecione um resultado para
ir à mensagem. Use **Voltar** no navegador para rever os resultados ou limpe
o campo de pesquisa para voltar às conversas ou às mensagens do chat.

[![Busca sem conexão por “endereço” com resultados nas conversas Casa e Viagem.](imagens/pesquisar-mensagens-sem-conexao.webp)](imagens/pesquisar-mensagens-sem-conexao.webp)

A busca alcança o texto guardado neste navegador. Se a mensagem que você procura
ficou fora da cópia salva, reconecte o Oiê para pesquisar o histórico completo
disponível.

## Envio agendado

Abra o menu junto ao botão de envio, escolha **Agendar mensagem** e confira a data, o horário e a mensagem. Em **Agendadas**, você pode editar, reagendar, cancelar ou pedir o envio imediato enquanto a mensagem ainda estiver aguardando o primeiro envio.

Em **Agendadas**, abra **Conversas** para buscar um destinatário. Você pode marcar várias conversas. Na rc.49, com o motor atualizado, a busca também encontra destinatários de envios pendentes que só aparecem nas páginas seguintes da agenda.

![Janela Agendar mensagem, com a prévia do lembrete e o campo de data e horário de envio.](imagens/envio-agendado.webp "Prévia e horário do envio agendado")

Na agenda da navegação lateral, use **Conversas** para escolher quais destinos quer ver. Os envios do mesmo dia aparecem em colunas quando há espaço.

Para editar uma mensagem sem sair da conversa, abra **Mais opções → Agendadas → Editar mensagem**. O campo de mensagem passa a mostrar o conteúdo agendado. Ao salvar ou cancelar a edição, o Oiê restaura o rascunho que você já tinha naquela conversa.

Se o aplicativo ou a conexão estiver indisponível no horário marcado, o Oiê tenta enviar quando voltar a funcionar e a conversa estiver disponível. Em salas, isso exige que a conta esteja na sala. Fechar a janela do aplicativo de desktop o mantém em execução. Para encerrá-lo, use [o atalho de saída](../atalhos/index.md#aplicativo-de-desktop).

## Citações e texto original

O Oiê exibe Markdown nas mensagens por padrão. Para citar várias linhas, comece cada uma com `>`. Uma linha sem o marcador fica fora da citação, mesmo quando vem logo abaixo. Em **Configurações → Este dispositivo → Mídia e exibição**, use **Formatar mensagens com Markdown** para ajustar a exibição. Abra **Ações da mensagem → Copiar mensagem original** para copiar o texto com os marcadores Markdown. Se a conexão cair com a conversa aberta, você ainda pode copiar o texto.

O lápis ao lado do horário indica uma mensagem editada. Selecione-o ou use **Ver histórico de edições** no menu da mensagem para consultar as versões que o Oiê guardou.

![Mensagem com a pergunta citada em um bloco e a resposta Sim! Levo o café abaixo, fora da citação.](imagens/citacao-markdown.webp "Citação em Markdown")

## Idioma e fuso horário

Em **Configurações**, use a busca em **Idioma** ou **Fuso horário de exibição**. A lista de fusos mostra o deslocamento atual em relação ao UTC, incluindo o horário de verão quando aplicável. Você pode buscar pela cidade ou pelo deslocamento, como `UTC+09:00`.

Escolha **Fuso horário do sistema** para acompanhar o fuso do dispositivo. O idioma e o fuso escolhidos ficam salvos neste aplicativo ou navegador. O agendamento também usa esse fuso para interpretar a data e o horário que você informa.

![Busca de fuso horário em Geral, com o fuso do sistema e o deslocamento UTC do resultado.](imagens/fuso-horario.webp "Busca por fuso horário")
