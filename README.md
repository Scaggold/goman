goman
=====

Scaffold tools

### How to use

#### Requirements

Need below tools to build:

- git
- go (1.3+)
- make

#### Build

Close this repository and type `make` command with your os name:

```
git clone git@github.com:Scaggold/goman.git
cd goman
make [darwin|linux|window]
```

Then, `gm` command binay created to `build/`. Copy or make symlink to this binay:

```
# Mac OS case
cd build/darwin

# copy or make symlink
cp ./gm /usr/local/bin/
ln -s ./gm /usr/local/bin/gm
```

### Usage

##### Getting template set

Getting template-set from repository:

```
gm get [template-name]
```

Template-sets are available at https://github.com/Scaggold

#### Generating temlate set

Move to diretory you want to create, and type below:

```
gm gen [template-set]
```

Will generated template-set these rules:

- dotfile is ignored (e.g. `.git`, `.rc`)

#### More

See command line help

```
gm help
```

### Rooadmap

We want to implement these features:

- Post generate script running
- Configuration with `setting.json`
- Templates installation switching Local/Remote

