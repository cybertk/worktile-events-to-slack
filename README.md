# worktile-events-to-slack

> 推送Worktile通知到Slack

[![CI Status](http://img.shields.io/travis/cybertk/worktile-events-to-slack/master.svg?style=flat)](https://travis-ci.org/cybertk/worktile-events-to-slack)
[![Docker Repository on Quay.io](https://quay.io/repository/quanlong/worktile-events-to-slack/status "Docker Repository on Quay.io")](https://quay.io/repository/quanlong/worktile-events-to-slack)

*worktile-events-to-slack* 可以将Worktile发送的webhook通知转发到Slack

## Getting Started

创建Slack Incoming Webhook

1. 进入在Slack的Integrations设置
1. 新建一个Incoming Webhook
1. 获取Incoming Webhook地址, 如https://XXX.slack.com/services/hooks/incoming-webhook?token=...

部署**worktile-events-to-slack**

    docker run -e SLACK_URL="https://XXX.slack.com/services/hooks/incoming-webhook?token=TOKEN" quanlong/worktile-events-to-slack

`SLACK_URL`是从Slack获取的Incoming Webhook地址

创建Worktile Webhook

1. 进入Worktile的项目设置中的Webhook设置
1. 添加Webhook
1. 通知URL字段填写**worktile-events-to-slack**部署后的URL
