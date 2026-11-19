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

As seções **Ler conversas sem conexão** e **Pesquisar mensagens sem conexão**
foram conferidas na [revisão 77af49c do Oiê](https://github.com/araihu/oie/tree/77af49ce87946f082f50571c505d51496e161a33),
posterior à rc.44. As capturas mostram buscas e mensagens de conversas fictícias
no Oiê local em pt-BR com o navegador desconectado.

A configuração **Armazenamento sem conexão** foi conferida na
[revisão 821db0d do Oiê](https://github.com/araihu/oie/tree/821db0db68442780e25fd49417193f2cb840f84c).

O controle **Manter disponível sem conexão** foi conferido na
[revisão 1261089 do Oiê](https://github.com/araihu/oie/tree/1261089fcd25b0858fff8cbc3ff6a4b992495808),
posterior à rc.44. A captura mostra os detalhes de uma conversa fictícia no
Oiê local em pt-BR.

As instruções dos repositórios DNF, pacman e Flatpak seguem a publicação da [revisão 5a2339c](https://github.com/araihu/oie/tree/5a2339c87e8dc42078afcfc0be2efbade2c6c423).

A página **Baixar** aponta para os binários e para o código-fonte correspondente
da [v0.0.1-rc.32](https://github.com/araihu/oie/tree/v0.0.1-rc.32).
O guia de instalação e os trechos do manual sobre mensagens de voz, arquivos
baixados e o contador de não lidas acompanham essa versão.

Os guias de autenticação hospedada e conexão do desktop acompanham a rc.28.
O guia do CLI documenta a API 0.38.0 da rc.32. O guia de Kubernetes usa o
[chart 0.2.10](https://github.com/araihu/oie/tree/helm-v0.2.10/charts/oie),
distribuído como OCI no GHCR.

As seções de contatos, disponibilidade e leitura do manual foram conferidas com a
[RC 28 no commit 4f46384](https://github.com/araihu/oie/tree/4f46384c9084ff9098481fde658f9d9b55cc4367).
As capturas mostram a interface em pt-BR com contas fictícias. O link direto
para `install.sh` usa o arquivo da rc.32. Os comandos do
`oiectl` para remover salas e conversas diretas foram conferidos na mesma revisão.

O manual, os atalhos, os exemplos do CLI e os guias de instalação e arquitetura foram revisados com a [main do aplicativo em 9929391](https://github.com/araihu/oie/tree/9929391579532045b61fcad8d8883f7e3a351215) e a [v0.0.1-rc.3](https://github.com/araihu/oie/releases/tag/v0.0.1-rc.3). Ao atualizar, registre o SHA das páginas revisadas e confira as notas da versão oferecida em **Baixar**. O contrato HTTP fica no [catálogo OpenAPI da Arai Hû](https://oapi.araihu.com/catalogs/oie/documents/engine-next/).

A seção de atualização via APT no guia de instalação foi conferida com a [main do aplicativo em baae7e6](https://github.com/araihu/oie/tree/baae7e69541d4f32e999fa277238d34c9808e427) e com o pacote publicado da [v0.0.1-rc.7](https://github.com/araihu/oie/releases/tag/v0.0.1-rc.7).

A seção de backup antes de atualizar o banco foi conferida na [revisão dd2e610 do aplicativo](https://github.com/araihu/oie/tree/dd2e6106e56f34c86a1b74c6dd989ecdbd350bc6).

As seções sobre anexos, leitura, idioma e fuso horário, a instalação Flatpak e os exemplos de compilação foram conferidos na [revisão d6ce9ae do aplicativo](https://github.com/araihu/oie/tree/d6ce9ae95c6f67850cb7a8e2ab91979fe35171ea), usada na rc.9. O bundle da rc.8 usa `org.araihu.oie`.

As seções **Mensagens favoritadas** e **Envio agendado** foram conferidas na [revisão ad3f9be do aplicativo](https://github.com/araihu/oie/tree/ad3f9bebd78cb2f9a62769a3f26aabf4a419e808), usada na rc.11. A captura de **Favoritas** usa mensagens fictícias. A seção e a captura de **Favoritar em Mídia e links** foram conferidas na [revisão 902efcd](https://github.com/araihu/oie/tree/902efcdfb5c70e9e942d75eadaa61da2efedb028), usada na rc.12.

As seções de prévia e reprodução de vídeo foram conferidas na [revisão c0bc373 do aplicativo](https://github.com/araihu/oie/tree/c0bc373423340724b909e2e7028f324d005640d1). As capturas mostram uma conversa fictícia em pt-BR. A miniatura no rascunho e a geração de prévias sem FFmpeg estão na [rc.12](https://github.com/araihu/oie/tree/902efcdfb5c70e9e942d75eadaa61da2efedb028).

A seção **Retenção de mensagens** foi conferida na [revisão 505f6c9 do aplicativo](https://github.com/araihu/oie/tree/505f6c90632a68eea50db9dae92c55e2d533ed13), usada na rc.16. A captura foi feita com dados fictícios em uma instância isolada durante o desenvolvimento do recurso.

A seção **Notificações de mensagens** e os controles de nome e foto no **Perfil** acompanham a rc.23. A captura de notificações mostra um áudio de exemplo importado numa instância de teste em pt-BR. [Revisão do Oiê](https://github.com/araihu/oie/tree/c5c3f37aeb815dd9ee4a1e94512ebf7fd2467071)

As seções de **Diagnóstico** no manual e no guia de instalação acompanham a [revisão da rc.24](https://github.com/araihu/oie/tree/68f530a8f9da8039f46a0c8eb9cdbf700298ea1f). A captura usa dados fictícios em pt-BR. O caminho pela aba **Geral**, a
retenção e a confirmação de saída na aba **Motor** e o atalho para recuperar
uma conexão hospedada no desktop foram conferidos no
[PR #40 do Oiê](https://github.com/araihu/oie/pull/40) e estão na rc.32.

A instalação do `oiectl` por release independente foi conferida com a [release cli/v0.30.1](https://github.com/araihu/oie/releases/tag/cli/v0.30.1), gerada a partir do [contrato no commit f21196a](https://github.com/araihu/oie/tree/f21196a6d8fe2d1b75180470ccd26089db5d6bb9/api).

A seção **Histórico e busca em salas** descreve o realce e a seleção de versões anteriores disponíveis desde a rc.13. [Revisão 699be01 do Oiê](https://github.com/araihu/oie/tree/699be01c6aacad2a03f1ddcd2bdb4eb195a64c69)

As capturas da página inicial, do manual e dos atalhos mostram a interface em pt-BR com dados fictícios em instâncias isoladas. A captura da página inicial e as sete capturas de recursos mais antigas foram refeitas na [revisão e87ea75](https://github.com/araihu/oie/tree/e87ea75b0dac62915c0b4e7fe8cb49d01ecc6692). As capturas da busca geral e da solicitação de mensagem são de `6c8e8c0`, que tem a mesma interface. As capturas dos anexos e da busca de fuso horário foram feitas na revisão da rc.9 (`d6ce9ae`). As imagens estão em WebP.

Os fluxos de **Mídia e links**, **Encaminhar mensagens** e **Descobrir e gerenciar salas**, a busca de favoritas no CLI e o índice de busca padrão foram conferidos na [revisão 4809005 do aplicativo](https://github.com/araihu/oie/tree/48090057513b2eab83d526d44df1d7764ac16a0b), usada na rc.14. As capturas de encaminhamento e gerenciamento de salas vieram dos PRs #31 e #22.

A seção Docker da página de arquitetura descreve a [CI da imagem em f9c6a01](https://github.com/araihu/oie/commit/f9c6a01625232e7132a358b27372759144588258).

O fluxo de remoção de conta e o comando `oiectl accounts remove` foram conferidos com a [revisão da rc.22](https://github.com/araihu/oie/tree/ffa480cb81a5b79287191a2220d554c4180be3ac) e o contrato 0.34.0. A captura usa uma conta fictícia.

O trecho **Conferir o endereço e remover salas** no manual e o comando
`oiectl conversations remove-room-from-chats` foram conferidos na revisão de
desenvolvimento `b9a50667a861fd1eb26559e4aba3a0fdc9e6aecb` do Oiê, posterior
à versão oferecida em **Baixar**. As capturas em pt-BR usam salas fictícias.

A remoção de conversas diretas e o comando
`oiectl conversations remove-direct-from-chats` foram conferidos na revisão de
desenvolvimento `c9f670574b3ca067ff65305f45eef54d784e4878` do Oiê.

A remoção da sala das listas da conta que confirma sua exclusão foi conferida
na revisão de desenvolvimento `e81d0d535f2eb5cb2e9f857efd4c14de8131861e` do Oiê. A captura da confirmação
em pt-BR usa uma sala fictícia.

O indicador de estado no avatar das salas e o rótulo **Sala excluída** foram
conferidos na revisão de desenvolvimento
`0149b3894eb4f895f5ba5fadff3ee9bd89dc8552` do Oiê. A captura em pt-BR usa uma
sala fictícia excluída por outra conta.

As bordas dos avatares por conta de origem foram conferidas na revisão de
desenvolvimento `5b2287a8b90205d6da174cd8fe4da65e12c8fbce` do Oiê. A captura usa
duas contas fictícias que entraram na mesma sala.

A abertura do histórico de conversas removidas por links foi conferida na
revisão de desenvolvimento `0b61995218a13e1b3a467fefb55687addfc0fa34` do Oiê.

As mudanças de salas e conversas deste manual foram integradas à
[main do Oiê](https://github.com/araihu/oie/tree/18de5883825e05a1bc5ed81dae251d0dafe5780c)
depois da versão v0.0.1-rc.26.

A retração de mensagens, a edição de respostas com citação, o encaminhamento
para conversas diretas e a opção `--include-retracted` do CLI foram conferidos na
[revisão 71c41cd do Oiê](https://github.com/araihu/oie/tree/71c41cd34221ac98d8e6fc71529f8680db2d19f7).
As capturas usam duas contas fictícias conectadas a um servidor de teste.
Esses recursos estão disponíveis desde a rc.30.

O menu de anexos ao lado de **Enviar** foi conferido na
[revisão f9b3c71 do Oiê](https://github.com/araihu/oie/tree/f9b3c71722f6a3b008fbad71d40ae3ce5ebe55be).
A captura em pt-BR usa uma conversa fictícia.

O gravador de mensagens de voz, a prévia da gravação e a forma de onda no
balão foram conferidos na [revisão 96ca2c6 do Oiê](https://github.com/araihu/oie/tree/96ca2c68fb4d8290dd06b58418f23ab32e49c835).
As capturas em pt-BR usam uma conversa fictícia e áudio de teste.

A seção **Enviar sem conexão** foi conferida na
[revisão de desenvolvimento 34a1d58 do Oiê](https://github.com/araihu/oie/tree/34a1d58ec9cfc1ba9c3953f734c5437fc5e58645),
posterior à rc.32 de **Baixar**. A captura mostra uma conversa fictícia em pt-BR.

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
