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
