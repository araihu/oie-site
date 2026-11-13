# Site do Oiê

Site do [Oiê](https://github.com/araihu/oie), gerado com [Margo](https://github.com/araihu/margo/tree/v0.0.35). O conteúdo está em `content/` e a configuração em `site.yaml`.

## Desenvolvimento

Com Go 1.27.3 ou mais recente, sincronize o tema e gere os arquivos estáticos:

```sh
./scripts/sync-theme.sh
go run ./cmd/build-site
```

O comando gera o site com Margo. O código em `cmd/build-site` adapta os controles de busca, cópia e diagramas para pt-BR, sincroniza a seleção acessível da busca com a ordem dos resultados e inclui um sumário nas telas estreitas. Ao atualizar o Margo, confira esses controles no navegador. Para pré-visualizar a publicação com Python 3, sirva os arquivos gerados e abra `http://127.0.0.1:8080/`.

```sh
python3 -m http.server --bind 127.0.0.1 --directory dist 8080
```

O comando escreve em `dist/` e exige um destino que ainda não exista. Use `--output-dir` para gerar outra cópia.

Os links em `content/baixar/index.md` apontam para arquivos de uma versão específica. Ao atualizar a página, confira os nomes dos arquivos e as plataformas na publicação do Oiê.

Ao atualizar **Baixar**, confira também os comandos e requisitos do [guia de instalação](content/docs/instalacao/index.md) com o [guia do aplicativo](https://github.com/araihu/oie/blob/main/docs/installation.md).

## Documentação do Oiê

A página **Baixar** aponta para os binários e para o código-fonte correspondente da [revisão da v0.0.1-rc.8](https://github.com/araihu/oie/tree/1845d9df6c8f66d5bac41c8f4fc329074baa2850).

O manual, os atalhos, os exemplos do CLI e os guias de instalação e arquitetura foram revisados com a [main do aplicativo em 9929391](https://github.com/araihu/oie/tree/9929391579532045b61fcad8d8883f7e3a351215) e a [v0.0.1-rc.3](https://github.com/araihu/oie/releases/tag/v0.0.1-rc.3). Ao atualizar, registre o SHA das páginas revisadas e confira as notas da versão oferecida em **Baixar**. O contrato HTTP fica no [catálogo OpenAPI da Arai Hû](https://oapi.araihu.com/catalogs/oie/documents/engine-next/).

A seção de atualização via APT no guia de instalação foi conferida com a [main do aplicativo em baae7e6](https://github.com/araihu/oie/tree/baae7e69541d4f32e999fa277238d34c9808e427) e com o pacote publicado da [v0.0.1-rc.7](https://github.com/araihu/oie/releases/tag/v0.0.1-rc.7).

A migração do perfil Flatpak para `com.araihu.oie` foi documentada a partir da [revisão 16dd0c2 do aplicativo](https://github.com/araihu/oie/tree/16dd0c22001197d547e946c04ddd43999e98364f). O bundle da rc.8 usa `org.araihu.oie`.

A instalação do `oiectl` por release independente foi conferida com a [release cli/v0.30.1](https://github.com/araihu/oie/releases/tag/cli/v0.30.1), gerada a partir do [contrato no commit f21196a](https://github.com/araihu/oie/tree/f21196a6d8fe2d1b75180470ccd26089db5d6bb9/api).

As seções sobre mensagens não lidas, idioma e fuso horário foram conferidas com a [main do aplicativo em 43d6710](https://github.com/araihu/oie/tree/43d67105d77ef5393495025be26fc5258a0b2fc6).

As capturas da página inicial, do manual e dos atalhos mostram a interface em pt-BR com dados fictícios em instâncias isoladas. A captura da página inicial e as sete capturas de recursos mais antigas foram refeitas na [revisão e87ea75](https://github.com/araihu/oie/tree/e87ea75b0dac62915c0b4e7fe8cb49d01ecc6692). As capturas da busca geral e da solicitação de mensagem são de `6c8e8c0`, que tem a mesma interface. As imagens estão em WebP.

A seção Docker da página de arquitetura descreve a [CI da imagem em f9c6a01](https://github.com/araihu/oie/commit/f9c6a01625232e7132a358b27372759144588258).

## Publicação

O GitHub Actions roda a cada push em `main`, gera os arquivos do Margo e publica `dist/` no Cloudflare Workers com Wrangler. O domínio do Worker é [oie.araihu.com](https://oie.araihu.com/). O pipeline confere as URLs canônicas das páginas e compara o manifesto do Margo e o CSS publicado com os arquivos gerados. Ele usa o secret organizacional `CLOUDFLARE_API_TOKEN` e a variável do repositório `CLOUDFLARE_ACCOUNT_ID`.

## Identidade e licenças

- O tema Oiê e o catálogo Margo vêm de [araihu/assets](https://github.com/araihu/assets/tree/main/themes). Para atualizar os dois, altere `themes/assets-revision` para o SHA completo desejado e rode `./scripts/sync-theme.sh`. O script valida o SHA-256 do CSS. `themes/site.css` concentra os estilos da página.
- `content/brand/logo.svg` e `favicon.svg` são, respectivamente, `dist/brand/oie/logo/light-transparent-optical.svg` e `dist/icons/brand/oie-icon-light-transparent-optical.svg` do [Assets, revisão 428dbea](https://github.com/araihu/assets/tree/428dbea099d6bceab3333aeaef0ade4350d36a54). A versão escura em `logo-dark.svg` usa `dist/brand/oie/logo/dark-transparent-optical.svg` da mesma revisão. A imagem de compartilhamento usa o mesmo logotipo sobre um fundo da paleta do Oiê.
- O site usa GPL-3.0, a mesma licença do aplicativo. As licenças e os créditos de terceiros acompanham o repositório em `LICENSES/`.
