# A tool for manipulating and exploiting Mattermost logs

Usage:
  mmlogctl [command]

Available Commands:
  clean        clean logs
  completion   Generate the autocompletion script for the specified shell
  help         Help about any command
  top-messages list the top messages found in the logs (default 10)

Flags:
      --config string   config file (default is ./mmlogctl.json)
  -h, --help            help for mmlogctl

## Clean

clean logs by removing lines containing one of the given strings, and also check for docker format logs that has a timestamp at the beginning of the line.

Usage:
  `mmlogctl clean [input file] [output file] [flags]`

Flags:
  -e, --ExcludeLinesWith stringArray   exclude lines containing one of those strings (default [Worker: Job is complete,No notification data available,Notification will be sent,Notification sent,Notification not sent,Notification received,websocket.slow])

## Top Messages

list the top messages found in the logs (default 15)

Usage:
  `mmlogctl top-messages [input file] [flags]`

Flags:
  -n, --number int   number to display (default 15)