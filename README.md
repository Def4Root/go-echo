# go-echo
- linuxのechoコマンドをGoで作成しています。
- 名前がややこしいのでオウム返しを意味するparrotにしました。
## version
- v1.0: コマンドライン引数をそのまま返す
- v1.1: 環境変数を実装 
- v1.2: `-n`オプションによる改行なし指定を実装
- v1.3: `-e`オプションでエスケープ文字(`\\n`)を改行に置換
- v2.0: オプションをos.Argsからflag.Argに変更
## caution
linuxのechoコマンドをGoで作成し、開発環境もlinuxなので`build/parrot.exe`が正常に動作しない可能性があります。
