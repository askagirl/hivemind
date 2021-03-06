# Hivemind

[![Build Status](https://travis-ci.org/DarthSim/hivemind.svg?branch=master)](https://travis-ci.org/DarthSim/hivemind)

Hivemind is a tool for running processes of a development environment. At the moment, it supports Linux, FreeBSD, and Mac OS X.

<a href="https://evilmartians.com/?utm_source=hivemind">
<img src="https://evilmartians.com/badges/sponsored-by-evil-martians.svg" alt="Sponsored by Evil Martians" width="236" height="54">
</a>

#### Why did I decide to develop Hivemind?

I used to use [Foreman](https://github.com/ddollar/foreman) by @ddollar, but I noticed a few problems with it:

* Sometimes Foreman loses a part of apps' output;
* Foreman loses colors of most apps' output;
* Sometimes Foreman can't interrupt some apps.

So I decided to write an alternative that won't have these problems. Now - meet Hivemind.

## Installation

#### With Homebrew (macOS)

```bash
brew install hivemind
```

#### From Source

You need Go 1.5 or later to build the project.

```bash
$ go get -u -f github.com/DarthSim/hivemind
```
__Note:__ You need to set `GO15VENDOREXPERIMENT=1` to build hivemind with Go 1.5.

__Note:__ You can update Hivemind the same way.

## Usage

Hivemind works with a Procfile.

```Procfile
web: bin/rails server
worker: bundle exec sidekiq
assets: gulp watch
```

To get started, you just need to run Hivemind from your working directory containing Procfile.

```bash
$ hivemind
```

If Procfile isn't located in your working directory, you can specify it:

```bash
$ hivemind path/to/your/Procfile
```

Run `hivemind --help` to see other options.

## Author

Sergey "DarthSim" Aleksandrovich

Highly inspired by [Foreman](https://github.com/ddollar/foreman).

## License

Hivemind is licensed under the MIT license.

See LICENSE for the full license text.
