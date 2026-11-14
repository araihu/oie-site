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

A página **Baixar** aponta para os binários e para o código-fonte correspondente da [revisão da v0.0.1-rc.13](https://github.com/araihu/oie/tree/d6d0872739426f85a4774dd956888f7bb880cd97).

O manual, os atalhos, os exemplos do CLI e os guias de instalação e arquitetura foram revisados com a [main do aplicativo em 9929391](https://github.com/araihu/oie/tree/9929391579532045b61fcad8d8883f7e3a351215) e a [v0.0.1-rc.3](https://github.com/araihu/oie/releases/tag/v0.0.1-rc.3). Ao atualizar, registre o SHA das páginas revisadas e confira as notas da versão oferecida em **Baixar**. O contrato HTTP fica no [catálogo OpenAPI da Arai Hû](https://oapi.araihu.com/catalogs/oie/documents/engine-next/).

A seção de atualização via APT no guia de instalação foi conferida com a [main do aplicativo em baae7e6](https://github.com/araihu/oie/tree/baae7e69541d4f32e999fa277238d34c9808e427) e com o pacote publicado da [v0.0.1-rc.7](https://github.com/araihu/oie/releases/tag/v0.0.1-rc.7).

A seção de backup antes de atualizar o banco foi conferida na [revisão dd2e610 do aplicativo](https://github.com/araihu/oie/tree/dd2e6106e56f34c86a1b74c6dd989ecdbd350bc6).

As seções sobre anexos, leitura, idioma e fuso horário, a instalação Flatpak e os exemplos de compilação foram conferidos na [revisão d6ce9ae do aplicativo](https://github.com/araihu/oie/tree/d6ce9ae95c6f67850cb7a8e2ab91979fe35171ea), usada na rc.9. O bundle da rc.8 usa `org.araihu.oie`.

As seções **Mensagens favoritadas** e **Envio agendado** foram conferidas na [revisão ad3f9be do aplicativo](https://github.com/araihu/oie/tree/ad3f9bebd78cb2f9a62769a3f26aabf4a419e808), usada na rc.11. A captura de **Favoritas** usa mensagens fictícias. A seção e a captura de **Favoritar em Mídia e links** foram conferidas na [revisão 902efcd](https://github.com/araihu/oie/tree/902efcdfb5c70e9e942d75eadaa61da2efedb028), usada na rc.12.

As seções de prévia e reprodução de vídeo foram conferidas na [revisão c0bc373 do aplicativo](https://github.com/araihu/oie/tree/c0bc373423340724b909e2e7028f324d005640d1). As capturas mostram uma conversa fictícia em pt-BR. A miniatura no rascunho e a geração de prévias sem FFmpeg estão na [rc.12](https://github.com/araihu/oie/tree/902efcdfb5c70e9e942d75eadaa61da2efedb028).

A instalação do `oiectl` por release independente foi conferida com a [release cli/v0.30.1](https://github.com/araihu/oie/releases/tag/cli/v0.30.1), gerada a partir do [contrato no commit f21196a](https://github.com/araihu/oie/tree/f21196a6d8fe2d1b75180470ccd26089db5d6bb9/api).

A seção **Histórico e busca em salas** descreve o realce e a seleção de versões anteriores disponíveis na rc.13. [Revisão 699be01 do Oiê](https://github.com/araihu/oie/tree/699be01c6aacad2a03f1ddcd2bdb4eb195a64c69)

As capturas da página inicial, do manual e dos atalhos mostram a interface em pt-BR com dados fictícios em instâncias isoladas. A captura da página inicial e as sete capturas de recursos mais antigas foram refeitas na [revisão e87ea75](https://github.com/araihu/oie/tree/e87ea75b0dac62915c0b4e7fe8cb49d01ecc6692). As capturas da busca geral e da solicitação de mensagem são de `6c8e8c0`, que tem a mesma interface. As capturas dos anexos e da busca de fuso horário foram feitas na revisão da rc.9 (`d6ce9ae`). As imagens estão em WebP.

A seção Docker da página de arquitetura descreve a [CI da imagem em f9c6a01](https://github.com/araihu/oie/commit/f9c6a01625232e7132a358b27372759144588258).

### Referência do banco de dados

O catálogo e os diagramas registram a [revisão dd2e610 do Oiê](https://github.com/araihu/oie/tree/dd2e6106e56f34c86a1b74c6dd989ecdbd350bc6), fixada em `scripts/database-docs/oie-revision`. Para atualizar a referência, coloque nesse arquivo o SHA completo do Oiê que será documentado. Com Python 3 com SQLite e o [tbls v1.96.1](https://github.com/k1LoW/tbls/releases/tag/v1.96.1) no PATH, rode a partir da raiz do site.

```sh
git -C ../oie fetch origin main
./scripts/generate-database-docs.sh ../oie
```

O script aplica as migrações `.up.sql` desse commit a um SQLite temporário e gera o Markdown e os SVGs usados pelo Margo. Edite as descrições e os recortes em `scripts/database-docs/` e inclua os arquivos regenerados junto com a atualização do SHA. Confira na página **Banco de dados** quais migrações chegaram depois da versão oferecida em **Baixar**.

## Publicação

O GitHub Actions roda a cada push em `main`, gera os arquivos do Margo e publica `dist/` no Cloudflare Workers com Wrangler. O domínio do Worker é [oie.araihu.com](https://oie.araihu.com/). O pipeline confere as URLs canônicas das páginas e compara o manifesto do Margo e o CSS publicado com os arquivos gerados. Ele usa o secret organizacional `CLOUDFLARE_API_TOKEN` e a variável do repositório `CLOUDFLARE_ACCOUNT_ID`.

## Identidade e licenças

- O tema Oiê e o catálogo Margo vêm de [araihu/assets](https://github.com/araihu/assets/tree/main/themes). Para atualizar os dois, altere `themes/assets-revision` para o SHA completo desejado e rode `./scripts/sync-theme.sh`. O script valida o SHA-256 do CSS. `themes/site.css` concentra os estilos da página.
- `content/brand/logo.svg` e `favicon.svg` são, respectivamente, `dist/brand/oie/logo/light-transparent-optical.svg` e `dist/icons/brand/oie-icon-light-transparent-optical.svg` do [Assets, revisão 428dbea](https://github.com/araihu/assets/tree/428dbea099d6bceab3333aeaef0ade4350d36a54). A versão escura em `logo-dark.svg` usa `dist/brand/oie/logo/dark-transparent-optical.svg` da mesma revisão. A imagem de compartilhamento usa o mesmo logotipo sobre um fundo da paleta do Oiê.
- O site usa GPL-3.0, a mesma licença do aplicativo. As licenças e os créditos de terceiros acompanham o repositório em `LICENSES/`.
