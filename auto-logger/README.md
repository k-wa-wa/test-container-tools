# auto-logger

`interval`ミリ秒の感覚で、ログを出力し続ける。
`DEBUG / INFO / WARN / ERROR`の割合を、`ratio`で指定可能。

デフォルトでは`INFO`ログを1秒ごとに出力する。

## usage

```sh
docker run -it auto-logger /usr/local/bin/app -i 400 -r 0/1/0/1
```
