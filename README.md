## notify-fs
[![AUR](https://img.shields.io/github/license/codexlynx/notify-fs)](LICENSE) [![](https://goreportcard.com/badge/github.com/codexlynx/notify-fs)](https://goreportcard.com/report/github.com/codexlynx/notify-fs)

Connecting filesystem events with instant messaging.

### Behaviour:
When a file is created in the watched directory, it is automatically sent to the configured chat.

### Technology:
This service uses the [fsnotify](https://github.com/fsnotify/fsnotify) library. This library 
uses the system call `SYS_INOTIFY_INIT1` (294) on Linux systems and `ReadDirectoryChangesW` on Windows systems.

### Configuration:
This service is configured via environment variables.

| Variable Name      | Description             |
|--------------------|-------------------------|
| `TARGET_DIRECTORY` | Directory to watch      |
| `TELEGRAM_TOKEN`   | Telegram bot token      |
| `TARGET_CHAT_ID`   | Telegram target chat id |
| `ONLY_IMAGES`      | Send only images        |

### Run:

You can compile the binary.

#### Compile:
* Requirements:
    * A version of __Docker__ with __BuildKit__ support.
    * GNU __make__ utility.

* Procedure:
    * Run: `make`.
    * Check the correct creation of `dist` directory.
