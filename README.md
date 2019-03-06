
[![Built with Mage](https://magefile.org/badge.svg)](https://magefile.org)

mage-select
===

CLI frontend for [mage](https://magefile.org) based on [promptui](https://github.com/manifoldco/promptui).

![](.github/images/screenshot.png)	

# Why?	

Migrating from `Makefiles` I missed bash completion. [It doesn't look like](https://github.com/magefile/mage/issues/113#issuecomment-422638376) there will be bash completion for `mage` any time soon, so I wrote this little program.

# Installing

[Install mage](https://magefile.org) first.

Clone or `go get` this repository and run `mage install` in `$GOPATH/src/github.com/iwittkau/mage-select`.

```bash
GO111MODULE=off go get github.com/iwittkau/mage-select
cd $GOPATH/src/github.com/iwittkau/mage-select
mage install
```

The `mages` binary will be installed into `$GOPATH/bin`.

# Usage

To select a `mage` target run:

```bash
mages
``` 

To abort selection press `CTRL`+`C`.

**You can also just start to type to search for a target!**

The `mages` command passes all arguments to `mage`, so you can run `mages -h -debug` to show the help of a target while debug output is enabled, for example.

There is nothing else to configure and there are no other options or flags at the moment, just run `mages` and select the `mage` target.

One little exception: the `-version` flag is overwritten, because mage-select bundles its own mage.
To show the version of `mages` and its bundled mage, run:

```bash
mages -version
``` 

# Caveats

`mages` bundles its own `mage` and therefor needs to be updated when Mage gets updated. Run `mages -version` to show the bundled version.

# Issues

If you encounter bugs or missing features, feel free to open an [issue](https://github.com/iwittkau/mage-select/issues).

# Contributing

If you'd like to contribute, please fork this repository and create a feature branch. Pull requests are welcome.

# Links

- [Magefile project](https://magefile.org)
- [promptui](https://github.com/manifoldco/promptui)
- [README inspiration](https://github.com/jehna/readme-best-practices)
- [ssh-select](https://github.com/iwittkau/ssh-select): `ssh` CLI frontend.
