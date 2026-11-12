# Site do Oiê

Site do [Oiê](https://github.com/araihu/oie), gerado com [Margo](https://github.com/araihu/margo/tree/v0.0.35). O conteúdo está em `content/` e a configuração em `site.yaml`.

## Desenvolvimento

Com Go 1.27.3 ou mais recente:

```sh
go run github.com/araihu/margo/cmd/margo@v0.0.35 serve .
```

Abra o endereço que o Margo imprimir. Para gerar os arquivos estáticos:

```sh
go run github.com/araihu/margo/cmd/margo@v0.0.35 site site.yaml
```

O Margo escreve em `dist/` e exige um destino que ainda não exista. Use `--output-dir` para gerar outra cópia. O GitHub Actions gera e publica o site a cada atualização de `main` em [araihu.github.io/oie-site](https://araihu.github.io/oie-site/).

Os links em `content/baixar/index.md` apontam para arquivos de uma versão específica. Ao atualizar a página, confira os nomes dos arquivos e as plataformas na publicação do Oiê.

## Identidade e licenças

- O [catálogo de temas do Assets](https://github.com/araihu/assets/blob/428dbea099d6bceab3333aeaef0ade4350d36a54/manifests/themes.yaml) ainda não inclui uma folha de estilo para Oiê. O tema local em `themes/oie.css` aplica a [paleta do aplicativo](https://github.com/araihu/oie/blob/f819b9d0fe746d9f0cd3e89905379eec760d3fb5/docs/branding.md) aos tokens de Goshtoso e Margo. O catálogo deriva do [catálogo de exemplo do Margo](https://github.com/araihu/margo/blob/v0.0.35/themes/margo.tokens.yaml). Após editar a folha de estilo, atualize o `css_digest` em `themes/oie.tokens.yaml` com o SHA-256 dos novos bytes.
- `content/brand/logo.svg` e `favicon.svg` são, respectivamente, `dist/brand/oie/logo/light-transparent-optical.svg` e `dist/icons/brand/oie-icon-light-transparent-optical.svg` do [Assets, revisão 428dbea](https://github.com/araihu/assets/tree/428dbea099d6bceab3333aeaef0ade4350d36a54). A versão escura em `logo-dark.svg` usa `dist/brand/oie/logo/dark-transparent-optical.svg` da mesma revisão. A imagem de compartilhamento usa o mesmo logotipo sobre um fundo da paleta do Oiê.
- A captura de tela usa apenas dados de exemplo do [servidor de demonstração do Oiê](https://github.com/araihu/oie/tree/f819b9d0fe746d9f0cd3e89905379eec760d3fb5/cmd/oie-mock).
- O site usa GPL-3.0, a mesma licença do aplicativo. As licenças e os créditos de terceiros acompanham o repositório em `LICENSES/`.
