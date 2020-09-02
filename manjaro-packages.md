# PineBook Pro Configuration (ArchLinux / Manjaro Edition)

ArchLinux seems to be a little different to using `yum` or `pkg` or `apt`, but uses `pacman` for package management.

## Manjaro Linux Packages

Important packages to get working

- avahi
- htop
- keepass / keepassxc
- telegram-desktop
- wire-desktop
- docker
- gcc
- git
- gnupg
- iputils (may by default)
- jq
- keybase / keybase-gui / kbfs
- libffi
- libusb (may be default)
- lsof
- make
- net-tools (may be default)
- tmux
- vim (may be default)
- neovim-qt (for a cool window)

## Other stuff of note

- [get-pip.py](https://github.com/pypa/get-pip)
- [fswatch](https://github.com/emcrisostomo/fswatch.git)
- [nvm](https://github.com/nvm-sh/nvm)
- [Tor Browser](https://git.torproject.org/tor-browser.git)
- [pyenv](https://github.com/pyenv/pyenv) (so we can install whatever version of python)
- rbenv

## Python libraries

- docker-compose

## Enabling services on start

```bash
# or replace with disable to switch off
systemctl enable avahi
systemctl enable docker
```

