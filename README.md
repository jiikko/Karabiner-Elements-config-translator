# Karabiner-Elements-config-translator

https://github.com/pqrs-org/Karabiner-Elements のための設定ファイルを簡単に書くためのツールです。

## モチベーション

以下の理由から Karabiner-Elements の設定をシンプルな記述で書けるようにするためのツールを作成しました。

- JIS, ANSI キーボードを両方使っているとそれらに対応した設定を書く必要があり冗長になりがち
- 繰り返す記述が多い
- 設定内容の多くに記号的な表現が多く、読み書きが若干難しい

## 使い方

```sh
command <input-file> --complex_modifications | ls ~/.config/karabiner/assets/complex_modifications/1234.json
```

## 例
