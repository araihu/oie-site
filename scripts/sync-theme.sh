#!/usr/bin/env bash
set -euo pipefail

cd "$(dirname "${BASH_SOURCE[0]}")/.."
revision=$(<themes/assets-revision)
if [[ ! "$revision" =~ ^[0-9a-f]{40}$ ]]; then
  echo 'Use um SHA completo de araihu/assets em themes/assets-revision.' >&2
  exit 1
fi

tmp=$(mktemp -d themes/.oie-theme.XXXXXX)
trap 'rm -rf "$tmp"' EXIT
for file in oie.css oie.tokens.yaml; do
  curl --fail --silent --show-error --location --retry 3 --max-time 30 \
    --proto '=https' --proto-redir '=https' \
    "https://raw.githubusercontent.com/araihu/assets/$revision/themes/$file" \
    --output "$tmp/$file"
done

# Margo also validates this digest when building the site.
expected=$(sed -n 's/^css_digest: sha256-//p' "$tmp/oie.tokens.yaml")
actual=$(shasum -a 256 "$tmp/oie.css" | cut -d' ' -f1)
if [[ "$actual" != "$expected" ]]; then
  echo 'O tema Oiê não corresponde ao SHA-256 declarado no catálogo.' >&2
  exit 1
fi

mv "$tmp/oie.css" themes/oie.css
mv "$tmp/oie.tokens.yaml" themes/oie.tokens.yaml
