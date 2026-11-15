---
title: Instalação
description: Instale o Oiê para desktop ou rode no navegador com Docker e Kubernetes. Configure o chart Helm com Ingress ou Gateway API.
language: pt-BR
---

# Instalação do Oiê

Os pacotes do Oiê estão disponíveis para **macOS com Apple Silicon** e **Ubuntu 24.04 x86-64**. O repositório APT atende também ao **Trisquel 12 x86-64**, e o bundle Flatpak roda no **Linux x86-64**. Para começar a conversar, tenha uma conta XMPP e siga o [passo a passo do manual](../manual/index.md#contas-xmpp) depois de abrir o aplicativo.


## Instalar com APT

No **Trisquel 12** e no **Ubuntu 24.04**, em computadores x86-64, adicione
uma vez o repositório assinado da Arai Hû. Baixe o
[arquivo oie.sources](https://araihu.github.io/oie/apt/oie.sources) e execute
na pasta onde o salvou:

```sh
sudo install -m 0644 oie.sources /etc/apt/sources.list.d/oie.sources
sudo apt update
sudo apt install oie
```

Abra o Oiê pelo menu de aplicativos. O repositório inclui pré-lançamentos.

### Atualizar com APT

Se o repositório já estiver cadastrado, feche o Oiê e execute o comando abaixo.

```sh
sudo apt update && sudo apt upgrade oie
```

A atualização direta da `rc.5` para a `rc.7` passou no Trisquel 12 e no
Ubuntu 24.04 LTS, em ambientes x86-64 emulados, com o repositório já cadastrado.

Para remover o aplicativo e manter seu perfil, execute `sudo apt remove oie`.

## Instalar com Flatpak

No Linux x86-64, baixe o [bundle Flatpak da v0.0.1-rc.24](https://github.com/araihu/oie/releases/download/v0.0.1-rc.24/oie-v0.0.1-rc.24-x86_64.flatpak). Adicione o Flathub se ainda não estiver cadastrado, para que o Flatpak possa baixar o runtime GNOME 51. Execute os comandos na pasta onde salvou o arquivo.

```sh
flatpak remote-add --user --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo
flatpak install --user ./oie-v0.0.1-rc.24-x86_64.flatpak
```

Se você já usava o Flatpak da `rc.8` ou anterior, [migre seu perfil](#migrar-para-comaraihuoie) antes de abrir o novo Oiê. Depois, abra o aplicativo pelo menu ou execute `flatpak run com.araihu.oie`.

### Migrar para `com.araihu.oie`

Até a `rc.8`, o Flatpak usava o identificador `org.araihu.oie`. Para manter suas contas, histórico e mídia, encerre qualquer instância do Oiê em execução com **Ctrl + Q**. Após instalar a `rc.9` e antes de abri-la pela primeira vez, copie o perfil completo. O comando só copia quando a pasta do perfil novo ainda não existe.

```sh
perfil_antigo="$HOME/.var/app/org.araihu.oie/data/oie"
dados_novos="$HOME/.var/app/com.araihu.oie/data"
if [ ! -d "$perfil_antigo" ]; then
  printf '%s\n' 'Perfil do Flatpak antigo não encontrado.'
elif [ -e "$dados_novos/oie" ]; then
  printf '%s\n' 'Já existe um perfil na nova instalação.'
else
  mkdir -p "$dados_novos" && cp -a "$perfil_antigo" "$dados_novos/oie"
fi
```

Abra o novo pacote com `flatpak run com.araihu.oie` e confira suas contas e seu histórico antes de remover a instalação antiga. Se já existe um perfil na nova instalação, encerre o Oiê e mova `~/.var/app/com.araihu.oie/data/oie` para um local de backup antes de repetir a cópia. Essa cópia usa o perfil antigo completo. Guarde o backup para preservar os dados criados na nova instalação.

## Instalar com Homebrew

No **macOS com Apple Silicon**, instale o aplicativo pelo [tap da Arai Hû](https://github.com/araihu/homebrew-oie).

```sh
brew install --cask araihu/oie/oie
```

Abra o Oiê na pasta **Aplicativos**. Se o macOS bloquear a abertura, siga os passos de [primeira abertura](#primeira-abertura-no-macos).

No **Ubuntu 24.04 x86-64**, instale as bibliotecas do sistema e depois a fórmula do Oiê.

```sh
sudo apt update
sudo apt install libgtk-4-1 libwebkitgtk-6.0-4 libsoup-3.0-0
brew install --formula araihu/oie/oie
```

Abra o aplicativo em uma sessão gráfica com `oie-desktop`.

Para atualizar, encerre o Oiê e execute `brew upgrade --cask araihu/oie/oie` no macOS ou `brew upgrade --formula araihu/oie/oie` no Linux. Para remover o aplicativo, use `brew uninstall --cask araihu/oie/oie` ou `brew uninstall --formula araihu/oie/oie`.

## Instalar o pacote publicado

Na página [Baixar](../../baixar/index.md), abra **Notas da versão e instalador**. Nos arquivos da versão que você quer instalar, baixe `install.sh`. O instalador escolhe o pacote para seu sistema e confere a soma SHA-256 antes de extrair os arquivos.

No **Ubuntu 24.04**, instale primeiro as bibliotecas usadas pelo aplicativo.

```sh
sudo apt update
sudo apt install curl ca-certificates libgtk-4-1 libwebkitgtk-6.0-4 libsoup-3.0-0
```

O macOS já inclui as bibliotecas da interface. O instalador usa `curl`, `tar`, `gzip` e `sha256sum` ou `shasum`.

Abra o Terminal na pasta onde salvou o arquivo e execute o instalador com a sua conta de usuário.

```sh
sh ./install.sh
```

| Sistema | Onde o instalador coloca o Oiê | Como abrir |
| --- | --- | --- |
| macOS | `~/Applications/Oie.app` | Abra o Oiê no Finder. |
| Ubuntu | `~/.local/bin/oie-desktop` | Abra o Oiê pelo menu de aplicativos em uma sessão gráfica. |

O instalador usa por padrão a versão em que foi publicado. Para escolher outra versão publicada ou mudar o destino, consulte as opções `--version` e `--prefix` com `sh ./install.sh --help`.

### Primeira abertura no macOS

O pacote tem assinatura ad hoc. Se o macOS bloquear a abertura, tente abrir o Oiê uma vez. Depois, vá a **Ajustes do Sistema → Privacidade e Segurança** e escolha **Abrir Mesmo Assim** para o Oiê. [Instruções da Apple](https://support.apple.com/pt-br/102445).

## Compilar a partir do código-fonte

Use Go 1.27.4 ou mais recente e compile no sistema e na arquitetura em que pretende usar o aplicativo. Nos exemplos, substitua `v0.0.1-rc.24` pela versão que você escolheu na página [Baixar](../../baixar/index.md). O comando instala o executável em `~/.local/bin`.

### Ubuntu 24.04 x86-64

Instale as ferramentas e bibliotecas de desenvolvimento. Se o `pkg-config` indicar uma biblioteca ausente, confira o pacote correspondente na lista do APT.

```sh
sudo apt update
sudo apt install git ca-certificates build-essential pkg-config \
  libgtk-4-dev libwebkitgtk-6.0-dev libsoup-3.0-dev
pkg-config --modversion gtk4 webkitgtk-6.0 libsoup-3.0 gio-unix-2.0
```

```sh
mkdir -p "$HOME/.local/bin"
GOBIN="$HOME/.local/bin" GOFLAGS= GOWORK=off CGO_ENABLED=1 \
  GONOPROXY=github.com/araihu/oie \
  go install -tags production github.com/araihu/oie/cmd/oie-desktop@v0.0.1-rc.24
```

### macOS com Apple Silicon

Instale as ferramentas de linha de comando do Xcode com `xcode-select --install` se elas ainda estiverem ausentes.

```sh
mkdir -p "$HOME/.local/bin"
GOBIN="$HOME/.local/bin" GOFLAGS= GOWORK=off CGO_ENABLED=1 \
  GONOPROXY=github.com/araihu/oie \
  MACOSX_DEPLOYMENT_TARGET=13.0 \
  CGO_CFLAGS="-O2 -g -mmacosx-version-min=13.0" \
  CGO_LDFLAGS="-O2 -g -mmacosx-version-min=13.0" \
  go install -tags production -ldflags='-extldflags=-mmacosx-version-min=13.0' \
  github.com/araihu/oie/cmd/oie-desktop@v0.0.1-rc.24
```

Os parâmetros de compilação e linkedição mantêm o mesmo alvo. Consulte o [escopo verificado no guia do repositório](https://github.com/araihu/oie/blob/c5c3f37aeb815dd9ee4a1e94512ebf7fd2467071/docs/installation.md#platforms-and-verified-scope) para as versões de sistema testadas.

### Abrir pelo Terminal

```sh
"$HOME/.local/bin/oie-desktop"
```

Se estiver trabalhando em um checkout, execute o comando de compilação da sua plataforma na raiz do repositório, substituindo o último argumento por `./cmd/oie-desktop`.

### Miniaturas de vídeo

Na rc.12, o navegador ou o aplicativo de desktop gera miniaturas com o
suporte a vídeo do próprio dispositivo. A geração dispensa o FFmpeg. Veja
[Prévia e reprodução de vídeos](../manual/index.md#prvia-e-reproduo-de-vdeos).

## Atualizar, remover e guardar seus dados

Se usou o instalador, encerre o Oiê e execute o instalador da versão desejada com o mesmo destino. Se compilou a partir do código-fonte, rode o comando de compilação com a nova versão. Os dados ficam em um perfil separado da instalação.

| Sistema | Diretório do perfil |
| --- | --- |
| macOS | `~/Library/Application Support/Oiê` |
| Linux | `$XDG_DATA_HOME/oie` quando `XDG_DATA_HOME` é um caminho absoluto, ou `~/.local/share/oie` |
| Flatpak (`com.araihu.oie`) | `~/.var/app/com.araihu.oie/data/oie`. Para versões até a rc.8, veja [como migrar o perfil](#migrar-para-comaraihuoie). |

O perfil guarda o histórico, as credenciais e a mídia retida. Encerre o aplicativo antes de copiar esse diretório. Use `--data-dir` para abrir o Oiê com outro perfil.

### Backup antes de atualizar o banco

Disponível desde a `v0.0.1-rc.10`.

Quando o Oiê precisa atualizar um banco existente na inicialização, primeiro guarda uma cópia em `private/migration-backups/` dentro do perfil. O nome começa com `oie-before-v` e inclui a versão anterior do banco e o horário em UTC. Reserve espaço para uma segunda cópia do banco. Se a cópia falhar, o Oiê interrompe a inicialização antes de atualizar o banco.

Guarde essas cópias em local privado, pois o banco contém o histórico e pode incluir senhas de convites para salas. Para guardar também a mídia e as credenciais de conexão XMPP, encerre o Oiê e copie o perfil inteiro. A opção **Apagar todos os dados** também remove as cópias automáticas do banco.

Se a janela não abrir após uma atualização, inicie o Oiê pelo terminal para ver o erro. No macOS, se usou `install.sh` com o destino padrão, execute `~/Applications/Oie.app/Contents/MacOS/oie-desktop`. Se usou o Homebrew com o destino padrão, execute `/Applications/Oie.app/Contents/MacOS/oie-desktop`. No Linux, abra o executável `oie-desktop` no caminho em que foi instalado ou execute `flatpak run com.araihu.oie` para Flatpak. Se o erro indicar falha no backup, confira o espaço livre e as permissões de gravação no perfil e tente novamente.

### Remover o aplicativo

Para remover uma instalação feita com `install.sh`, encerre o Oiê. No macOS, mova `Oie.app` para o Lixo. Na instalação padrão no Ubuntu, remova estes arquivos e a pasta de documentação.

- `~/.local/bin/oie-desktop`
- `~/.local/share/applications/org.araihu.oie.desktop`
- `~/.local/share/icons/hicolor/512x512/apps/org.araihu.oie.png`
- `~/.local/share/doc/oie`

Se usou `--prefix`, procure os mesmos arquivos no destino escolhido. Se usou `go install`, basta remover o executável instalado. Preserve o perfil se quiser manter seus dados.

## Rodar com Docker

Com o Docker em execução, inicie o Oiê com a imagem publicada e um volume para guardar seus dados.

```sh
docker run -d --name oie \
  --publish 127.0.0.1:8080:8080 \
  --volume oie-data:/data \
  ghcr.io/araihu/oie:latest
```

Abra [127.0.0.1:8080](http://127.0.0.1:8080) no navegador e [conecte sua conta XMPP](../manual/index.md#contas-xmpp). Para escolher a tag da imagem, atualizar o contêiner ou acessar o Oiê de outro computador, veja [Docker em Arquitetura e deploy](../arquitetura/index.md#docker).

## Rodar no Kubernetes com Helm

Use um cluster com armazenamento persistente, o Helm e o `kubectl`.
O [chart Helm 0.2.2](https://github.com/araihu/oie/releases/tag/helm-v0.2.2)
tem versão própria e usa por padrão a imagem do Oiê `v0.0.1-rc.21`.

```sh
helm upgrade --install oie \
  https://github.com/araihu/oie/releases/download/helm-v0.2.2/oie-0.2.2.tgz \
  --namespace oie --create-namespace --wait --timeout 10m
kubectl --namespace oie port-forward service/oie 8080:80
```

Abra [127.0.0.1:8080](http://127.0.0.1:8080) no navegador e
[conecte sua conta XMPP](../manual/index.md#contas-xmpp). O chart roda uma
instância e reserva 10 GiB para contas, histórico e mídia. Para usar um volume
já existente, indique o nome do PVC com `persistence.existingClaim`. Use um
PVC por instância. O volume é mantido quando você desinstala a release Helm.

O Oiê não tem login próprio. Para acesso remoto, configure um proxy HTTPS
com autenticação em todas as rotas, incluindo a API. Use Ingress ou conecte
o Oiê a um Gateway existente com `httpRoute.enabled`, `httpRoute.parentRefs`
e `httpRoute.hostnames`. Se o Gateway estiver em outro namespace, permita
rotas do namespace do Oiê no listener. Veja exemplos no
[guia do chart](https://github.com/araihu/oie/tree/helm-v0.2.2/charts/oie#remote-access).

Se o timeout de requisição do Gateway interromper as atualizações ao vivo
das conversas, acrescente este trecho ao arquivo de valores do chart. Use
os CRDs da Gateway API v1.2.0 ou posterior e um controlador com suporte ao
timeout de requisição no HTTPRoute.

```yaml
httpRoute:
  timeouts:
    request: "0s"
```

O valor `"0s"` desativa o timeout de requisição para todos os caminhos da rota
do Oiê. Se você omitir o valor, o controlador aplica seu comportamento padrão.
[Timeouts do HTTPRoute na Gateway API](https://gateway-api.sigs.k8s.io/reference/api-types/httproute/#timeouts-optional)

### Acesso ao diagnóstico no servidor

Configure `OIE_DIAGNOSTICS_KEY` no processo do servidor para habilitar o
visualizador. No Docker, use um arquivo de ambiente privado. O chart no
código-fonte da rc.24 aceita `diagnostics.existingSecret` para apontar para
um Secret que contenha a chave. Proteja o acesso remoto com HTTPS e
autenticação no proxy.

O registro também fica disponível em `docker logs oie` ou
`kubectl logs -n oie statefulset/oie`. Consulte o
[manual do diagnóstico](../manual/#diagnstico-e-registro-de-erros)
para copiar ou baixar o arquivo pelo aplicativo.

## Rodar pelo navegador ou usar o CLI

Para executar o motor no navegador ou em um servidor, consulte [Arquitetura e deploy](../arquitetura/index.md). Para usar o `oiectl`, veja o [guia do CLI](../cli/index.md).
