Keylime Tenant
==============

GoLang implementation of the Keylime Tenant

This project is not complete, don't use this yet. Use the Python
version over at [Keylime Python Repo](https://github.com/keylime/keylime)

```
# keylime_tenant --help
The Keylime Tenant CLI tool is used for the configuration
of KeyLime Agent nodes.

Usage:
  keylime_tenant [command]

Available Commands:
  add         Add an agent to the Verifier
  delete      Delete an Agent from the Verifier
  help        Help about any command
  list        List the operational state of an Agent
  reactivate  Reactivate and Agent
  regdelete   Delete an Agent from the Register
  status      Report the current status of the Agent
  update      Update the agent (delete & add)

Flags:
      --config string   config file (default is $HOME/.keylime_tenant.toml)
  -h, --help            help for keylime_tenant
  -t, --toggle          Help message for toggle

Use "keylime_tenant [command] --help" for more information about a command.

```

Sub command help is also Available

```
# keylime_tenant list --help

Lists the operational state of the an Agent

Usage:
  keylime_tenant list [flags]

Flags:
  -h, --help          help for list
      --uuid string   The UUID of the Agent to list

Global Flags:
      --config string   config file (default is $HOME/.keylime_tenant.toml)
```
