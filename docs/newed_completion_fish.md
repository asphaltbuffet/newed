## newed completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	newed completion fish | source

To load completions for every new session, execute once:

	newed completion fish > ~/.config/fish/completions/newed.fish

You will need to start a new shell for this setup to take effect.


```
newed completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -c, --config string   configuration file
```

### SEE ALSO

* [newed completion](newed_completion.md)	 - Generate the autocompletion script for the specified shell

