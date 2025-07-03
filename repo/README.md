# repo

`repo` is a simple command-line tool that opens the current git repository's remote URL in your default web browser. If you are in a subdirectory of the repository, it will open the corresponding tree view for the current branch.

## Installation

To install, you can use `go install`:

```sh
go install codeberg.com/usysrc/belt/repo@latest
```

## Usage

Simply run `repo` in any directory within a git repository:

```sh
repo
```

The tool will automatically detect the remote URL and open it in your browser.

## Supported Forges

`repo` currently supports the following git forges:

*   GitHub
*   GitLab
*   Codeberg
*   Gitea

If your forge is not listed, the tool will default to the Codeberg style URL structure, which may or may not work.
