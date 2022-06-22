# pf
## Path format/fix command line tool.

A *very* simple command line tool to help you fix or understand your current path.

Obviously it's best to get your path right in your environment and/or set up the path correctly in your
shell, but sometimes you're in a shell that you don't really want to restart and it would be nice to
just read or tweak the path directly.

## Examples:

### Move a directory or directories to the start of the path:

```bash
$ eval $(pf first Go/bin)
```

This will move any folder that matches "Go/bin" to the front of the path. The `eval` is necessary
because a tool like `pf` can't modify its parent's environment otherwise; `pf` actually prints out
a line that looks like `PATH=/your:/new:/path...` to be consumed by `bash`. I usually use `bash`, so
I'm probably not going to be motivated to support other shells, but pull requests that detect the
containing shell and modify the output appropriately are welcome.

All path expressions in `pf` use Go regular expressions to filter the paths, so
this also works:

```bash
$ eval $(pf first /bin$)
```

The above will rewrite the path so all path components that end in `/bin` will be
brought to the front of the path.

Finally, you can pull multiple directory patterns to the front of the path, and it will add
them in order of the patterns on the command line.

```bash
$ eval $(pf first Go /bin$)
```

The Go/bin binary folder followed by all (other) path components that end in `/bin`
will be moved to the front of the path.

### Search the current path

```bash
$ pf search go local
/home/tim-m/.local/bin
/home/tim-m/GCloud/google-cloud-sdk/bin
/home/tim-m/go/bin
/usr/local/sbin
/usr/local/bin
/usr/local/games
```

Search all path components for any paths that match the given regular expressions. Useful to see
if a folder is present in a larger path, or to see what order two paths fall in the path order.
The example above prints all paths that either match "go" or "local", and then shows only the
relevant components in their path order.

### Delete a path component

```bash
$ eval $(pf delete Go/bin$)
```

Delete a matching path component from the path. It's recommended that you first `search` using the same
pattern before you use delete. *All* paths that match the regular expression will be deleted from the
path.

### Print the current path, one component per line, Linux host:

```bash
$ pf print

/usr/local/sbin
/usr/local/bin
/usr/sbin
/usr/bin
/sbin
/bin
/usr/games
/usr/local/games
# ...
```

### Print the current path, one component per line, Windows host from MSYS 2.0 Bash:

```bash
$ pf print

C:\Windows\system32
C:\Windows
C:\Windows\System32\Wbem
C:\Windows\System32\WindowsPowerShell\v1.0
# ...
```

## Credits
Written by Tim Mensch, who is not much of a Go programmer, so any suggestions to improve the idiomatic usage
of Go would be appreciated.