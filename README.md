GitHub is usually one SPOF for a lot of applications, this agent takes a look
and notify on your slack channel it's the GitHub's status change.

```bash
github-spof -p 10 -t slackToken -c channel
```

* **p** is the ping time, default it's 10 second but you change it for example to
  ping the api each second (-p 1)
* **t** is the slack token, grab it from your account
* **c** is the name of the channel to sent the status
