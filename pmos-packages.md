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


## Convergence Edition 

* Insert JumptDrive SD card and boot up
* Plug into PC
* go to [`pmbootstrap`](https://gitlab.com/postmarketOS/pmbootstrap.git) repo and run `./pmbootstrap.py init`
* To export an image Follow the instructions at https://wiki.postmarketos.org/wiki/PINE64_PinePhone_(pine64-pinephone)#Internal_eMMC



### Alpine Packages

- docker
- avahi
- gnupg
- alpine-sdk
- vim
- neovim
- python3
- libffi-dev
- py3-pip
- python3-dev
- openssl-dev
- tmux
- go

### RC Update stuff after

```bash
rc-update add docker
rc-update add avahi-daemon boot
```
