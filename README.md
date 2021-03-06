# aws-config-manager

Manage multiple AWS credentials and config file pairs.

Do you work with a variety of AWS configs/credentials and want to manage them separately? Now you can.

## Install

```shell
brew install cdimascio/tap/aws_config_manager
```
or
```
brew tap cdimascio/tap
aws_config_manager
```

## Usage

```
➜ acm
NAME:
   acm - AWS Config and Credentials Manager

USAGE:
   acm command [command options] [<setting>]

DESCRIPTION:
   Manages many .aws/credentials and .aws/config files as settings

COMMANDS:
   cat           cat a credentials or config file.
   create        creates a new empty setting.
   current, cur  shows the current setting
   edit          edits a credentials or config file.
   list, ls      list all settings
   remove, rm    removes a setting
   use           sets the current setting
   help, h       Shows a list of commands or help for one command

EXAMPLES:
   acm ls
   acm use default
   acm create my-config
   acm edit -t creds
   acm edit -t conf my-config
```

**What is a setting?**

A setting is a name used to identify a pair of aws config and credentials files. The first time you use `acm`, a `default` setting is created. The `default` setting is a reference to your current `~/.aws/credentials` and `~/.aws/config`. To create additional settings, run `acm create my_setting`, then `acm edit -t config my_setting` and `acm edit -t creds my_setting` to edit the config and credentials file for the `my_setting` setting.

AWS enables you to use a single pair e.g. ~/.aws/crendentials and ~/.aws/config. This library allows you manage multiple pairs or 'settings'.



### List

List all settings

```
acm list
```

### Use

Use a specific setting

```
acm use default
```

### Edit

Edit a setting

```
# edit the currently active aws credentials file
acm edit -t credentials

# edit a specific aws credentials file
acm edit -t credentials my_creds

# edit the currently active aws config file
acm edit -t config

# edit a specific aws config file
acm edit -t config my_creds
```

### And more...

## License
MIT
