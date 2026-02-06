# belt

A collection of tiny command-line tools written in Go.

> [!CAUTION]
> These tools are primarily for my personal use and may not be rigorously tested, production-ready, or extensively documented. Use them at your own discretion and risk. They may evolve rapidly or break without notice.
> If a tool within this collection becomes popular, it might be moved to a separate repository for further development.

## Installing and Running Tools

### Nix

This repository is a [Nix flake](https://nixos.wiki/wiki/Flakes).

You can run any of the tools by using `nix run`.

For example, to run `hasenfetch`:

```shell
nix run git+https://codeberg.org/usysrc/belt.git#hasenfetch
```

### Go

If you have Go installed and GOBIN in your PATH, you can install any of the tools by using `go install`.

For example, to run `hasenfetch`:

```shell
go install codeberg.org/usysrc/belt/hasenfetch@latest
```

## Tools

Here is a list of all the tools available:

*   **hasenfetch**: minimal system information
*   **hex**: hex viewer
*   **jenv**: json environment variables
*   **jo**: json object manipulation
*   **nibs**: project tool for lua development
*   **obs**: interact with Obsidian from the terminal
*   **pal**: palette files manipulation
*   **proceed**: continue execution of a program
*   **repo**: open the forges web page of a repository
*   **serve**: tiny web server for serving files
*   **slow**: slow down output of a program
*   **ssl-expiry**: check SSL certificate expiry of remote domains
*   **timezone**: output current time in different timezones
*   **urlencode**: encode and decode strings to URLs
*   **uuid**: generate UUIDs
*   **xls-format**: change formatting inside of excel files
