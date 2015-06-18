# worktile-events-to-slack

> 推送Worktile通知到Slack

[![CI Status](http://img.shields.io/travis/cybertk/worktile-events-to-slack/master.svg?style=flat)](https://travis-ci.org/cybertk/worktile-events-to-slack)
[![Docker Repository on Quay.io](https://quay.io/repository/quanlong/worktile-events-to-slack/status "Docker Repository on Quay.io")](https://quay.io/repository/quanlong/worktile-events-to-slack)

**worktile-events-to-slack**可以将Worktile的事件通知转发到Slack, **Slack Incoming Webhook**的地址可以通过**QueryString**参数`slack_url`动态配置.

## Getting Started

部署worktile-events-to-slack

    docker run quanlong/worktile-events-to-slack

创建Slack Incoming Webhook

1. 进入在Slack的Integrations设置
1. 新建一个Incoming Webhook
1. 获取Incoming Webhook地址, 如https://XXX.slack.com/services/hooks/incoming-webhook?token=...

创建Worktile Webhook

1. 进入Worktile的项目设置中的Webhook设置
1. 添加Webhook
1. 通知URL字段填写worktile-events-to-slack webhook地址, 如 https://example.com?slack_url=https://hooks.slack.com/hooks/AAA

## Test

Development environment is backed by Docker, so enter dev env with

    make dev

Then test with

    make test
