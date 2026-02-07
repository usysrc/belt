export GITHUB_REPO:="https://github.com/usysrc/belt"
export CODEBERG_REPO:="https://codeberg.org/usysrc/belt.git"

update:
  #!/bin/bash
  nix flake check

  read -p "Nix flake check successful. Proceed with update? (y/N) " -n 1 -r
  echo

  if [[ $REPLY =~ ^[Yy]$ ]]; then
      nix flake update
  else
      echo "Update cancelled."
  fi

push:
  #!/bin/sh
  git push origin HEAD
  # if the github remote is set up, push there as well
  if git remote get-url github >/dev/null 2>&1; then
    git push github HEAD
  fi
  # if the codeberg remote is set up, push there as well
  if git remote get-url codeberg >/dev/null 2>&1; then
    git push codeberg HEAD
  fi

git-setup:
  #!/bin/bash
  set -euo pipefail
  origin_url="$(git remote get-url origin)"
  github_url=$GITHUB_REPO
  codeberg_url=$CODEBERG_REPO

  if [[ "$origin_url" == *codeberg.org* ]]; then
    if ! git remote get-url github >/dev/null 2>&1; then
      git remote add github "$github_url"
      echo "Added github remote -> $github_url"
    fi
  fi
  if [[ "$origin_url" == *github.com* ]]; then
    if ! git remote get-url codeberg >/dev/null 2>&1; then
      git remote add codeberg "$codeberg_url"
      echo "Added codeberg remote -> $codeberg_url"
    fi
  fi
