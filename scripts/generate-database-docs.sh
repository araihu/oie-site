#!/usr/bin/env bash
set -euo pipefail

if [[ $# != 1 ]]; then
  echo "Uso: $0 /caminho/do/repositorio/oie" >&2
  exit 1
fi

oie_repo=$(cd "$1" && pwd)
cd "$(dirname "${BASH_SOURCE[0]}")/.."
revision=$(<scripts/database-docs/oie-revision)
revision=$(git -C "$oie_repo" rev-parse --verify "$revision^{commit}")
export OIE_SCHEMA_REVISION="$revision"
tbls_bin=${TBLS_BIN:-tbls}

stage=$(mktemp -d)
trap 'rm -rf "$stage"' EXIT
python3 - "$oie_repo" "$revision" "$stage/oie.sqlite" <<'PY'
import sqlite3
import subprocess
import sys
from contextlib import closing

repo, revision, database = sys.argv[1:]
root = "internal/adapters/sqlite/migrations"
files = subprocess.check_output(
    ["git", "ls-tree", "-r", "--name-only", revision, "--", root],
    cwd=repo, text=True,
).splitlines()
migrations = sorted(path for path in files if path.endswith(".up.sql"))
if not migrations:
    raise SystemExit(f"Nenhuma migração encontrada em {revision}.")

with closing(sqlite3.connect(database)) as db:
    db.execute("PRAGMA foreign_keys = ON")
    for path in migrations:
        sql = subprocess.check_output(["git", "show", f"{revision}:{path}"], cwd=repo, text=True)
        try:
            db.executescript("BEGIN;\n" + sql + "\nCOMMIT;")
        except sqlite3.Error as error:
            raise SystemExit(f"Falha ao aplicar {path}: {error}") from error
    if db.execute("PRAGMA foreign_key_check").fetchone() is not None:
        raise SystemExit("O schema migrado tem violações de chave estrangeira.")
PY

dsn="sqlite://$stage/oie.sqlite"
"$tbls_bin" doc "$dsn" "$stage/catalogo" -c scripts/database-docs/catalogo.yaml --sort --without-er
mv "$stage/catalogo/README.md" "$stage/catalogo/index.md"

mkdir "$stage/diagramas"
"$tbls_bin" out "$dsn" -c scripts/database-docs/catalogo.yaml --sort -t svg -o "$stage/diagramas/completo.svg"
for theme in conversas historico envios interacoes conteudo; do
  "$tbls_bin" out "$dsn" -c scripts/database-docs/diagramas.yaml --sort --viewpoint "$theme" -t svg -o "$stage/diagramas/$theme.svg"
done

# Substitua os arquivos gerados somente após a migração e todas as exportações.
output=content/docs/banco-de-dados
for directory in catalogo diagramas; do
  rm -rf "$output/$directory"
  mv "$stage/$directory" "$output/$directory"
done
