# Slash Command

`/open search まぜそば大陸` ってポストしてみよう

## って何

DiscordとかSlackにあるみたいなやつ。

## コマンド一覧

| COMMAND                            | ARGS                                           | ACTION                              | SHORTCUT                     |
|------------------------------------|------------------------------------------------|-------------------------------------|------------------------------|
| `/help`                            |                                                | [README.md](/README.md) を開く       |                              |
| `/help command`                    |                                                | これを開く                           |                              |
| `/help config`                     |                                                | [CONFIG.md](./CONFIG.md) を開く     |                              |
| `/open config`                     |                                                | 設定ファイルの場所を開く              | <kbd>Ctrl</kbd>+<kbd>,</kbd> |
| `/open log`                        |                                                | ログの場所を開く                     | <kbd>Ctrl</kbd>+<kbd>.</kbd> |
| `/open profile (<HANDLE> / <DID>)` | HANDLE / DID : ユーザーのハンドル名またはDID      | ユーザープロフィールを開く           |                               |
| `/open search <TEXT>`              | TEXT : 検索したい文言（v10時点では）              | ユーザープロフィールを開く           |                              |
| `/open weather <ADDRESS>`          | ADDRESS : 地域                                  | 天気を調べる                        |                               |
| `/pizza <ADDRESS>`                 | ADDRESS : 郵便番号（7桁）                       | ピザ注文（ピザハット）                |                              |
| `/post chikuwa`                    |                                                | ちくわ。                            |                              |
| `/post earthquake`                 |                                                | 地震だ！                            |                              |
| `/post version`                    |                                                | 今使っているバージョン               |                              |
| `/quit`                            |                                                | 沈没                                | <kbd>Ctrl</kbd>+<kbd>q</kbd> |
| `/set foter <FOOTER>`              | FOOTER : 投稿の末尾に付与する文章（スペース区切り）   | フッターを設定する                    |                             |
| `/reset foter`                     |                                                | フッターをリセットする                 |                             |

<!-- 
/set footer #spam
/set header #eggs
/unset footer
/unset header
-->