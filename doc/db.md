# redis 键值说明

## 礼品码

* `gift:count:{code}` 记录礼品码剩余次数，如果无限次则置-1，服务器保证该值或者大于等于0，或者为-1
* `gift:user:{code}` 记录礼品对应的用户，如果通用则记为空""，需要保证用户Id不为空
* `gift:detil:{code}` 记录礼品细节，如金币多少等，目前没有区分道具，士兵，英雄的种类。

## 用户
* `army:{userId}` 记录用户的士兵数
* `hero:{userId}` 记录用户的英雄数
* `coin:{userId}` 记录用户的金币
* `diamond:{userId}` 记录用户的钻石
* `prop:{userId}` 记录用户的道具

