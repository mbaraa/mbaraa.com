Assuming you're here because you want LSP(Language Server Protocol) for your neovim setup, but your Linux distro doesn't provide the latest version of lua-language-server, so in this article we'll install lua-language-server from source.

#### Dependencies installation

Make sure that you have, Ninja Build, GCC(some distros need G++ as well), and Clang.

To install the dependencies on Gentoo run:

```bash
sudo emerge -qav sys-devel/gcc sys-devel/clang dev-util/ninja
```

#### Compiling LSP's source code

##### Cloning the repo:

- Clone LSP's repo into a directory where you keep bins and stuff
- I use `~/.local/bin` so I'll just clone it there
- `git clone https://github.com/LuaLS/lua-language-server ~/.local/bin/lua-language-server`
- `cd ~/.local/bin/lua-language-server`

##### Compile the stuff:

- Download the sub-modules of the cloned repo
- `git submodule update --recursive`
- Download the `ninja` luamake rules
- `cd 3rd/luamake`
- `git submodule update --init`
- Run the compile script
- `compile/install.sh`
- `cd ../../`
- `./3rd/luamake/luamake rebuild`

##### Add the the executables' path to your path

```bash
SHELL_NAME=`basename $SHELL`
SHELL_RC="./.${SHELL_NAME}rc"
echo 'export PATH="${HOME}/.local/bin/lua-language-server/bin:${PATH}"' >> $SHELL_RC
```

Now re-login or run

```bash
source $SHELL_RC
$SHELL_NAME
```

##### Require the new installed lsp server

Add this line to `~/.config/nvim/init.lua` or to where you put lsp's config in Neovim

```lua
require('lspconfig').sumneko_lua.setup {}
```

Restart Neovim and your good to go.
