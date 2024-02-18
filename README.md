# まぜそば大陸

<div align="center">
  ![Screenshot](docs/images/screenshot.png)
</div>

[![GitHub Release Date](https://img.shields.io/github/release-date/kakakaya/mazesoba-continent?style=flat)](https://github.com/kakakaya/mazesoba-continent/releases)
[![GitHub Downloads (all assets, latest release)](https://img.shields.io/github/downloads/kakakaya/mazesoba-continent/latest/total?sort=semver&style=flat)](https://github.com/kakakaya/mazesoba-continent/releases)
[![GitHub License](https://img.shields.io/github/license/kakakaya/mazesoba-continent?style=flat)](https://github.com/kakakaya/mazesoba-continent?tab=MIT-1-ov-file#readme)
[![GitHub Repo stars](https://img.shields.io/github/stars/kakakaya/mazesoba-continent)](https://github.com/kakakaya/mazesoba-continent)
[![codecov](https://codecov.io/gh/kakakaya/mazesoba-continent/branch/master/graph/badge.svg)](https://codecov.io/gh/kakakaya/mazesoba-continent)

投稿力の変わらないただ一つのBlueskyクライアント。そう、まぜそば大陸ならね。

Out of respect for [ラーメン大陸](https://forest.watch.impress.co.jp/docs/news/559014.html).

## 簡単な使い方

初回起動すると「設定ファイルが未設定だよ」みたいなことを言うので、設定ファイルの場所を開いて `config.toml` で以下の内容を入力してください。

```toml
[Credential]
  Identifier = "<ここにIdentifierを入れる>"
  Password = "<ここにパスワードを入れる>"
  Host = "https://bsky.social"
```

Windows や Linux なら <kbd>Ctrl</kbd> + <kbd>Enter</kbd> macOS なら <kbd>Cmd</kbd> + <kbd>Enter</kbd> で投稿ができます。

## 複雑な使い方

- [Slash Command](./docs/SLASH_COMMAND.md)
  - "/help"
