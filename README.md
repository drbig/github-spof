Why `github-spof`? Because GitHub is usually one SPOF for a lot of
applications, this agent takes a look and notify on your slack channel it's the
GitHub's status change.

You can install it as a simple wget or curl, you use our compiled binary.
```
curl -LSs https://github.com/gianarb/github-spof/releases/download/0.2.1/github_spof_0.2.1_darwin_386 > github_spof
curl -LSs https://github.com/gianarb/github-spof/releases/download/0.2.1/github_spof_0.2.1_linux_386 > github_spof
curl -LSs https://github.com/gianarb/github-spof/releases/download/0.2.1/github_spof_0.2.1_linux_arm > github_spof
```
Make it executable and go!


```bash
github-spof -p 10 -t slackToken -c channelCode
```

* **p** is the ping time, default it's 10 second but you change it for example
  to ping the api each second (-p 1)
* **t** is the slack token, grab it from your account
* **c** is the code of the channel to sent the status
