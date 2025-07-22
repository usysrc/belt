{
  description = "A collection of command-line tools written in Go for my personal computing, DevOps, sysadmin, or software engineering needs.";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      nixpkgs,
      flake-utils,
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };

        make_tool = name: hash: {
          inherit name;
          src = ./${name};
          vendorHash = hash;

          meta = {
            license = pkgs.lib.licenses.mit;
          };

          buildInputs = [ pkgs.notmuch ];

        };

        tools = builtins.mapAttrs (_: tool: pkgs.buildGoModule tool) {
          hasenfetch = make_tool "hasenfetch" "sha256-L5FZufDpwhv9rFdB5ELeJpROrtnQVN3aaTww9u9DY8A=";
          hex = make_tool "hex" "sha256-+aMFr9k1itFXWCGh3Z2jy/XyiS/l303eEVf8kBCBj5M=";
          jenv = make_tool "jenv" null;
          jo = make_tool "jo" "sha256-9gO00c3D846SJl5dbtfj0qasmONLNxU/7V1TG6QEaxM=";
          nibs = make_tool "nibs" "sha256-lcjv0tCFPga4n2lc5rRXe5A1jIDLScKEsWrG0a/Sftc=";
          obs = (make_tool "obs" "sha256-+Ezs6+YOOIESXrQneAQAsfvo3L6LwIiBx3LEybgEqBw=") // {
            doCheck = false;
          };
          pal = make_tool "pal" null;
          proceed = make_tool "proceed" null;
          repo = make_tool "repo" null;
          serve = make_tool "serve" "sha256-JFvC9V0xS8SZSdLsOtpyTrFzXjYAOaPQaJHdcnJzK3s=";
          slow = make_tool "slow" null;
          ssl-expiry = make_tool "ssl-expiry" null;
          timezone = make_tool "timezone" "sha256-JFvC9V0xS8SZSdLsOtpyTrFzXjYAOaPQaJHdcnJzK3s=";
          urlencode = make_tool "urlencode" "sha256-OEXvKQ/dBxhz6/pbQNDYIjBf3O0x36ZE3Se/FqEgYRg=";
          uuid = make_tool "uuid" "sha256-PfJNr7t/27PSnwIwFv0kHV3f+er0fpHwqddS8yS7ofo=";
          xls-format = make_tool "xls-format" "sha256-IvH6IKMmJ/yM7ZbkNdVkmiuRlTJtXy1eHh5mCKolfKk=";
        };
      in
      {
        packages = tools // {
          default = pkgs.symlinkJoin {
            name = "belt";
            paths = builtins.attrValues tools;
          };
        };

        devShells.default = pkgs.mkShell {
          hardeningDisable = [ "fortify" ];
          buildInputs = with pkgs; [
            go
            go-tools # linter (`staticcheck`)
            delve # debugger
            golangci-lint # linter (`golangci-lint run`), formatter
          ];
        };
      }
    );
}
