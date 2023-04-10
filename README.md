# README

投稿力の変わらないただ一つのBlueskyクライアント。そう、まぜそば大陸ならね。

Out of respect for [ラーメン大陸](www16.atpages.jp/nigore/soft/).

## About

チェックが付いた機能は実装済み(つまりまだなにもできていない)

- [x] 投稿
- ウィンドウ
    - [x] 表示
    - [ ] 残り文字数
    - [ ] 半透明表示
    - [ ] TL(現時点で実装予定なし)
    - [ ] フッター
- キーボード操作
    - [x] Ctrl+Enterで入力内容をポスト
    - [ ] F1キーでTL画面呼び出し(実装予定なし)
    - [ ] F2キーで設定画面呼び出し
    - [ ] F3キーでURL短縮
    - [ ] F4キーで入力文字列のデフラグ
    - [ ] F5キーでポスト画面に入力されている文字列の翻訳(実装予定なし)
    - [ ] F6キーでポスト画面に入力されている文字列の英日翻訳(実装予定なし)
    - [ ] F7キーでポスト画面に入力されている文字列の日英日翻訳(実装予定なし)
    - [ ] F8キーで新規付箋(実装予定なし)
    - [ ] F9キーで｢地震だ！｣機能。どこの田舎よりも素早い速報を！
    - [ ] F10キーでCDトレイの開閉。(実装予定なし)
    - [ ] F11キーでポストにフッターのオンオフ。(空文字無効)
    - [ ] F12キーでポストにハッシュタグのオンオフ。(空文字無効)
    - [ ] Escキーで終了or隠れる。
- マウス
    - [ ] 残り文字数の部分をダブルクリックでTL画面呼び出し。
    - [ ] 残り文字数の部分を右クリックでメニュー呼び出し。
    - [ ] 残り文字数の部分を掴んだら移動。
- アレ
    - [x] ちくわ
    - [ ] ハイパーステルスアレ
- おまけ
    - [ ] @最新版天童アリス：最新版の確認。
    - [ ] 新しいまぜそば大陸の存在…？：最新版の確認。
    - [ ] まさか…？！新しいまぜそば大陸？！：最新版の確認。
    - [ ] まぜそば大陸、用意できてる？：最新版の確認。
    - [ ] kakakayaくん…。：おいやめろ。

## Develop

### Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

### Building

To build a redistributable, production mode package, use `wails build`.
