## newed completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(newed completion bash)

To load completions for every new session, execute once:

#### Linux:

	newed completion bash > /etc/bash_completion.d/newed

#### macOS:

	newed completion bash > $(brew --prefix)/etc/bash_completion.d/newed

You will need to start a new shell for this setup to take effect.


```
newed completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -c, --config string       configuration file
  -t, --templates strings   template(s) to apply
```

### SEE ALSO

* [newed completion](newed_completion.md)	 - Generate the autocompletion script for the specified shell

