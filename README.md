# NagiosXI CLI

Command Line Interface (CLI) to interact with NagiosXI API.

[![Build Status](https://travis-ci.com/romainDavaze/nagiosxi-cli.svg?branch=master)](https://travis-ci.com/romainDavaze/nagiosxi-cli)

## Configuration

### Global

When starting, nagiosxi-cli reads a config file and load variables from it.

By default, if the parameter `--config` is not specified, it looks for a `.nagiosxi-cli.yml` file under user's home directory.
You can see an example [here](examples/nagiosxi-cli.yml).

**The API token must belong to a NagiosXI admin.**


### NagiosXI objects

You can add and delete several NagiosXI objects by providing a specific file as a parameter.

As of now, these objects are supported :
- Commands ([example](examples/commands.yml))
- Contactgroups ([example](examples/contactgroups.yml))
- Contacts ([example](examples/contacts.yml))
- Hostgroups ([example](examples/hostgroups.yml))
- Hosts ([example](examples/hosts.yml))
- Servicegroups ([example](examples/servicegroups.yml))
- Services ([example](examples/services.yml))



## Usage

Please refer to the [commands documentation](docs/nagiosxi-cli.md)


## Author

<a href="https://romaindavaze.github.io/">Romain Davaze</a>