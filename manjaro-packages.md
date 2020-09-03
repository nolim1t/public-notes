# PineBook Pro Configuration (ArchLinux / Manjaro Edition)

ArchLinux seems to be a little different to using `yum` or `pkg` or `apt`, but uses `pacman` for package management.

## Getting set up

Check out my collection of UNIX [dotfiles](https://gitlab.com/nolim1t/dotfiles). Feel free to suggest others.

```bash
mkdir src
cd src
git clone https://gitlab.com/nolim1t/dotfiles.git
cd ~
# my vimrc
ln -s src/dotfiles/vimrc .vimrc
# my shell profile
# or pick out stuff
ln -s src/dotfiles/profile .profile
# neovim
mkdir -p .config/nvim
cd .config/nvim
ln -s ~/src/dotfiles/nvim/init.vim init.vim
```

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
- neovim / neovim-qt (for a cool window). Starting to like this more than VIM
- patch (wow seriously?!)
- flex
- bison
- linux-headers
- bc
- uboot-tools
- mtools


## Other stuff of note

- [get-pip.py](https://github.com/pypa/get-pip) - because manjaro doesn't have pip installed?
- [fswatch](https://github.com/emcrisostomo/fswatch.git)
- [nvm](https://github.com/nvm-sh/nvm)
- [Tor Browser](https://git.torproject.org/tor-browser.git)
- [pyenv](https://github.com/pyenv/pyenv) (so we can install whatever version of python)
- [rbenv](https://github.com/rbenv/rbenv)

## Python libraries

- docker-compose

## Enabling services on start

```bash
# or replace with disable to switch off
systemctl enable avahi
systemctl enable docker
```

