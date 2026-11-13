---
title: Instalação
description: Instale o Oiê com Homebrew, APT, Flatpak ou o instalador, rode com Docker ou compile a partir do código-fonte.
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

No Linux x86-64, baixe o [bundle Flatpak da v0.0.1-rc.9](https://github.com/araihu/oie/releases/download/v0.0.1-rc.9/oie-v0.0.1-rc.9-x86_64.flatpak). Adicione o Flathub se ainda não estiver cadastrado, para que o Flatpak possa baixar o runtime GNOME 51. Execute os comandos na pasta onde salvou o arquivo.

```sh
flatpak remote-add --user --if-not-exists flathub https://dl.flathub.org/repo/flathub.flatpakrepo
flatpak install --user ./oie-v0.0.1-rc.9-x86_64.flatpak
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

Use Go 1.27.4 ou mais recente e compile no sistema e na arquitetura em que pretende usar o aplicativo. Nos exemplos, substitua `v0.0.1-rc.9` pela versão que você escolheu na página [Baixar](../../baixar/index.md). O comando instala o executável em `~/.local/bin`.

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
  go install -tags production github.com/araihu/oie/cmd/oie-desktop@v0.0.1-rc.9
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
  github.com/araihu/oie/cmd/oie-desktop@v0.0.1-rc.9
```

Os parâmetros de compilação e linkedição mantêm o mesmo alvo. Consulte o [escopo verificado no guia do repositório](https://github.com/araihu/oie/blob/d6ce9ae95c6f67850cb7a8e2ab91979fe35171ea/docs/installation.md#platforms-and-verified-scope) para as versões de sistema testadas.

### Abrir pelo Terminal

```sh
"$HOME/.local/bin/oie-desktop"
```

Se estiver trabalhando em um checkout, execute o comando de compilação da sua plataforma na raiz do repositório, substituindo o último argumento por `./cmd/oie-desktop`.

### Miniaturas de vídeo

O FFmpeg é opcional e permite ao Oiê gerar miniaturas de vídeo. No Ubuntu, instale com `sudo apt install ffmpeg`. Se o aplicativo não encontrar o executável, passe o caminho absoluto com `oie-desktop --ffmpeg /caminho/para/ffmpeg`.

## Atualizar, remover e guardar seus dados

Se usou o instalador, encerre o Oiê e execute o instalador da versão desejada com o mesmo destino. Se compilou a partir do código-fonte, rode o comando de compilação com a nova versão. Os dados ficam em um perfil separado da instalação.

| Sistema | Diretório do perfil |
| --- | --- |
| macOS | `~/Library/Application Support/Oiê` |
| Linux | `$XDG_DATA_HOME/oie` quando `XDG_DATA_HOME` é um caminho absoluto, ou `~/.local/share/oie` |

O perfil guarda o histórico, as credenciais e a mídia retida. Encerre o aplicativo antes de copiar esse diretório. Use `--data-dir` para abrir o Oiê com outro perfil.

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

## Rodar pelo navegador ou usar o CLI

Para executar o motor no navegador ou em um servidor, consulte [Arquitetura e deploy](../arquitetura/index.md). Para usar o `oiectl`, veja o [guia do CLI](../cli/index.md).
