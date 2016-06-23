slackp
------

Post a [message to a Slack incoming webhook](https://api.slack.com/incoming-webhooks#sending_messages). Useful for a CI system, maybe.

# examples

Clone the [Github repo](https://github.com/dylanmei/slackp), build and run.

```
make
bin/slackp -help
bin/slackp -url https://hooks.slack.com/services/blah/blah/blah \
  -channel #slack-tests \
  -user SlackP \
  -icon ":golang:"
  "hello from slackp!"
```

Pull the [Docker image](https://hub.docker.com/r/dylanmei/slackp) and run.

```
docker pull dylanmei/slackp
docker run --rm dylanmei/slackp -help
docker run --rm \
  -e "SLACKP_URL=https://hooks.slack.com/services/blah/blah/blah" \
  -e "SLACKP_CHANNEL=#slack-tests" \
  -e "SLACKP_USER=SlackP" \
  -e "SLACKP_ICON=:golang:" \
  dylanmei/slackp "hello from slackp!"
```
