# THIP - Thingiverse Zip Downloader CLI

![License](https://img.shields.io/badge/license-GPL3-green?style=plastic)
![Go](https://img.shields.io/github/go-mod/go-version/thomasboom89/thip/main)

A cli scraper to bundle a zip archive with all stl files of one model found
on [Thingiverse](https://www.thingiverse.com/).

Since they struggled around a lot with the **_download all files_** function on their website, they now completely
removed it. For me, it is pita to download every file on it's own. I decided to make a quick poc for my personal use,
and now I like to share it with maker/creator/programmer and many more :)

## Installation

Installation can be different, depending on what OS you are on. Guide is written for linux x86 OS.
\
Requirements: installed working go environment on build machine >= 1.17

Clone the repo and build binary

```zsh
git clone https://github.com/ThomasBoom89/thip
cd thip/thip
go build -o ../build/thip -ldflags "-w -s"
```

Now you can use it directly from the build folder

```zsh
cd ../build
./thip
```

## Usage

To download all files into a zip archive, you have to take the model id from the thingiverse url.
\
For example: thingiverse.com/thing:123456 take the last part, this is the model id: 123456.
\
Now you can type ```./thip -m 123456``` and you will get an archive if the model is found and there are no errors.
\
You will find the archive in your current directory.

## Contribution

Contributions are always welcome.
\
If you find a bug, report it. If you have question, ask for help (in Disskussions). If you want to improve something,
make a PR. Also, there are several points that could be worked off to help other users to get THIP running:

- ~~Check installation on linux x86 (report also your version) and add a section to the installation guide~~
- Check installation on Windows (report also your version) and add a section to the installation guide
- Check installation on macOS (report also your version) and add a section to the installation guide
- Check installation on linux arm (report also your version) and add a section to the installation guide
- Generate GitHub action to build binaries for all working OS at release

## License

THIP - Thingiverse Zip Downloader CLI, a cli scraper to bundle a zip archive with all stl files of one model found
on [Thingiverse](https://www.thingiverse.com/).\
Copyright (C) 2022 ThomasBoom89. GNU GPL3 license (GNU General Public License Version 3).

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public
License as published by the Free Software Foundation, either version 3 of the License, or (at your option)
any later version.

This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with this program. If not, see:

1. The LICENSE file in this repository.
2. https://www.gnu.org/licenses/.
3. https://www.gnu.org/licenses/gpl-3.0.txt.

However, THIP - Thingiverse Zip Downloader CLI includes several third-party Open-Source libraries, which are licensed
under
their own respective Open-Source licenses. Libraries or projects directly included:

- spf13/cobra: [APACHE-2.0](https://github.com/spf13/cobra/blob/master/LICENSE.txt)