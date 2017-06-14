### Slack CLI

Slack, the _"messaging app for teams"_ offers an API that has been used to build multiple projects around it, from bots to independent clients as well as integrations with other external services. This project aims to offer a low level experience for advanced users that want to either drop the web client or interact with the API for testing purpose.

### Installation

```
go get -u github.com/cixtor/slackcli
```

### Usage

Use a [session token](https://api.slack.com/web#authentication) to authenticate the HTTP requests against the API service. Slack automatically generates a token for your when you open a new session [here](https://slack.com/messages/); you can see this token in the JavaScript console of your web browser if you type `boot_data.api_token` but be aware that it will expire once you close the session, consider to use a [legacy token](https://api.slack.com/custom-integrations/legacy-tokens) instead.

```
$ export SLACK_TOKEN=xoxs-token
$ slackcli help
$ slackcli auth.test
$ slackcli chat.session
username:channel> :token xoxs-token
username:channel> :owner
username:channel> :exit
```

You can also export an environment variable `SLACK_VERBOSE=true` to print additional information during the execution of certain operations to troubleshoot issues with either the communication with th API or the program in itself.

### Features

The client is built on top of the [Bot Users](https://api.slack.com/bot-users) documentation. Most if not all the methods available in the API are implemented and can be executed placing a colon character as the suffix of each method.

Note that the client runs with the same chat session of the user that is using the program, but technically speaking the interaction is similar to that of a bot. This offers some advantages, for example, like other APIs and integrations, bot users are free. Unlike regular users, the actions they can perform are somewhat limited. For teams on the Free Plan, each bot user counts as a separate integration.
