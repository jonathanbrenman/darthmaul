```
______           _   _    ___  ___            _ 
|  _  \         | | | |   |  \/  |           | |
| | | |__ _ _ __| |_| |__ | .  . | __ _ _   _| |
| | | / _` | '__| __| '_ \| |\/| |/ _` | | | | |
| |/ / (_| | |  | |_| | | | |  | | (_| | |_| | |
|___/ \____|_|   \__|_| |_\_|  |_/\____|\____|_|
                                                

Generate command:
generate controller <name>
generate service <name>
generate repository <name>

```

### This cli provide a generator to speed up the development.

## How to install cli?
```
    $ git clone https://github.com/jonybrn/darthmaul.git
    $ echo 'export PATH="$PATH:$GOPATH/bin"' >> $HOME/.bash_profile
    $ cd darthmaul && go install && source $HOME/.bash_profile
```

## How to use it?
```
Example:

    $ cd folder_to_create_entity
    $ darthmaul generate controller <name>
    $ darthmaul generate service <name>
    $ darthmaul generate repository <name>
    $ darthmaul create-app <name>
```
