# worktile-events-to-slack

> 推送Worktile通知到Slack

[![CI Status](http://img.shields.io/travis/cybertk/worktile-events-to-slack/master.svg?style=flat)](https://travis-ci.org/cybertk/worktile-events-to-slack)

*worktile-events-to-slack* 可以将Worktile发送的webhook通知转发到Slack

## Getting Started

    docker run -e SLACK_URL="https://XXX.slack.com/services/hooks/incoming-webhook?token=TOKEN" quanlong/worktile-events-to-slack

`SLACK_URL`是Slack的Incoming Webhook地址
