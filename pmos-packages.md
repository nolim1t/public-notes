# Setup for PMOS PinePhone

Some documentation on what I did, so I can reproduce it later

## Reference material

- [PostmarketOS Source](https://gitlab.com/postmarketOS)
- [pmbootstrap source](https://gitlab.com/postmarketOS/pmbootstrap)

## Presetup stuff

### Prequisite packages

- git
- bash (installed already)

### Command line stuff

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


## Convergence Edition Default Install

What I set up

### Alpine Packages

- docker
- avahi
- gnupg
- alpine-sdk
- neovim
- python3 and python3-pip (not exact)

more coming soon
