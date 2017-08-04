# makeicon

[![Twitter: @beplusio](https://img.shields.io/badge/contact-@beplusio-blue.svg?style=flat)](https://twitter.com/beplusio)
[![License](https://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://github.com/beplus/makeicon/blob/master/LICENSE)
[![CircleCI Build Status](https://circleci.com/gh/beplus/makeicon.svg?style=shield)](https://circleci.com/gh/beplus/makeicon)
[![Release](https://img.shields.io/github/release/beplus/makeicon.svg?style=flat-square)](https://github.com/beplus/makeicon/releases/latest)

Generates mobile app icons in all resolutions for both **iOS** and **Android**.


## Table of contents
- [Requirements](#Requirements)
- [Installation](#Installation)
- [Usage](#Usage)
- [Contributors](#Contributors)
- [License](#License)


## Requirements:
* Mac OS X, Linux or Windows
* 64-bit version


## Installation

**Download** from Releases and extract _or_ use **Homebrew**

```bash
$ brew install https://raw.githubusercontent.com/beplus/makeicon/master/makeicon.rb
$ brew upgrade https://raw.githubusercontent.com/beplus/makeicon/master/makeicon.rb
```


## Usage
Command Line Interface

### Get Help on Usage
```bash
$ makeicon --help
# or
$ makeicon -h
```

### Generate the Icons
To create an `AppIcon` directory with generated Icons from a file `./Icon.png`
```bash
$ makeicon --file ./Icon.png
# or
$ makeicon -f ./Icon.png
```

**Important**: `Icon.png` file have to be a square image with `1024px` x `1024px` at least.

### Print the Current Version
```bash
$ makeicon --version
# or
$ makeicon -v
```


## Contributors
- [Jan Misek](https://github.com/janmisek1)
- [Igor Lamos](https://github.com/igorlamos)


## License
[MIT](./LICENSE) Copyright (c) 2017 BePlus s.r.o.
