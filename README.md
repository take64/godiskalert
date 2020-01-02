# godiskalert

Golang製のディスクアラートをslack通知するツールです

slackのincoming webhookのURLを取得し、引数に渡してあげます

```
./godiskalert -slack=https://hooks.slack.com/services/hoge/fuga
```

crontabに登録してあげると便利
