# ConfigFile

config.yaml
```
blog: a.github.io
best_authors: ["fengyun","lee","park"]
desc:
  counter: 521
  plist: [3, 4]
remote_url: ["127.0.0.1:514", "10.103.30.28:514"]
tag: test
```

解析结果
```
blog: a.github.io
best_authors: ["fengyun","lee","park"]
desc:
  counter: 521
  plist: [3, 4]
remote_url: ["127.0.0.1:514", "10.103.30.28:514"]
tag: test

--- t:
{a.github.io [fengyun lee park] {0 [3 4]} [127.0.0.1:514 10.103.30.28:514] test}

--- t:
{this is Blog [fengyun lee park myself] {99 [3 4]} [127.0.0.1:514 10.103.30.28:514] syslog}

--- t dump:
blog: this is Blog
best_authors: [fengyun, lee, park, myself]
desc:
  Counter: 99
  plist: [3, 4]
remote_url:
- 127.0.0.1:514
- 10.103.30.28:514
tag: syslog
```
