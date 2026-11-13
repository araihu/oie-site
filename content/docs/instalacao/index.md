---
title: Instalação
description: Instale o Oiê no macOS ou Ubuntu, rode com Docker ou compile a partir do código-fonte.
language: pt-BR
---

# Instalação do Oiê

Os pacotes do Oiê estão disponíveis para **macOS com Apple Silicon** e **Ubuntu 24.04 x86-64**. Para começar a conversar, tenha uma conta XMPP e siga o [passo a passo do manual](../manual/index.md#contas-xmpp) depois de abrir o aplicativo.

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

Use Go 1.27.4 ou mais recente e compile no sistema e na arquitetura em que pretende usar o aplicativo. Nos exemplos, substitua `v0.0.1-rc.3` pela versão que você escolheu na página [Baixar](../../baixar/index.md). O comando instala o executável em `~/.local/bin`.

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
  go install -tags production github.com/araihu/oie/cmd/oie-desktop@v0.0.1-rc.3
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
  github.com/araihu/oie/cmd/oie-desktop@v0.0.1-rc.3
```

Os parâmetros de compilação e linkedição mantêm o mesmo alvo. Consulte o [escopo verificado no guia do repositório](https://github.com/araihu/oie/blob/f9c6a01625232e7132a358b27372759144588258/docs/installation.md#platforms-and-verified-scope) para as versões de sistema testadas.

### Abrir pelo Terminal

```sh
"$HOME/.local/bin/oie-desktop"
```

Se estiver trabalhando em um checkout, execute o comando de compilação da sua plataforma na raiz do repositório, substituindo o último argumento por `./cmd/oie-desktop`.

### Miniaturas de vídeo

O FFmpeg é opcional e permite ao Oiê gerar miniaturas de vídeo. No Ubuntu, instale com `sudo apt install ffmpeg`. Se o aplicativo não encontrar o executável, passe o caminho absoluto com `oie-desktop --ffmpeg /caminho/para/ffmpeg`.

## Atualizar, remover e guardar seus dados

Para atualizar, encerre o Oiê e execute o instalador da versão desejada com o mesmo destino. Se compilou a partir do código-fonte, rode o comando de compilação com a nova versão. Os dados ficam em um perfil separado da instalação.

| Sistema | Diretório do perfil |
| --- | --- |
| macOS | `~/Library/Application Support/Oiê` |
| Linux | `$XDG_DATA_HOME/oie` quando `XDG_DATA_HOME` é um caminho absoluto, ou `~/.local/share/oie` |

O perfil guarda o histórico, as credenciais e a mídia retida. Encerre o aplicativo antes de copiar esse diretório. Use `--data-dir` para abrir o Oiê com outro perfil.

Para remover o aplicativo, encerre o Oiê. No macOS, mova `Oie.app` para o Lixo. Para a instalação padrão no Ubuntu, remova estes arquivos e a pasta de documentação.

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
