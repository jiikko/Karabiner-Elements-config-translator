# Karabiner-Elements-config-yaml

https://github.com/pqrs-org/Karabiner-Elements のための設定ファイルを簡単に書くためのツールです。

## モチベーション

以下の理由から Karabiner-Elements の設定をシンプルな記述で書けるようにするためのツールを作成しました。

- JIS, ANSI キーボードを両方使っているとそれらに対応した設定を書く必要があり冗長になりがち
- 繰り返す記述が多い
- 設定内容の多くに記号的な表現が多く、読み書きが若干難しい
- 誤った記述してもエラーが出ないので不親切

## 使い方

```sh
command <input-file> | ~/.config/karabiner/assets/complex_modifications/1234.json
```

## 例

### 1

input

```yaml
:title:
  - my config
:maintainers:
  - foo
:rules:
  - description: disable command + m(最小化)
    :from:
      - m
      - command
    :to:
      - none
```

output

```json
{
  "title": "my config",
  "maintainers": ["foo"],
  "rules": [
    {
      "description": "disable command + m(最小化)",
      "manipulators": [
        {
          "from": {
            "key_code": "m",
            "modifiers": { "mandatory": ["command"] }
          },
          "to": [{ "key_code": "vk_none" }],
          "type": "basic"
        }
      ]
    }
  ]
}
```

### 2

- conditions がある

input

```yaml

```

output

```json
{
  "title": "my config",
  "maintainers": ["foo"],
  "rules": [
    {
      "description": "command + l to ctrl + l for Terminal(ターミナルでの画面クリア対策)",
      "manipulators": [
        {
          "from": {
            "key_code": "l",
            "modifiers": {
              "mandatory": ["command"]
            }
          },
          "to": [
            {
              "key_code": "l",
              "modifiers": ["right_control"]
            }
          ],
          "type": "basic",
          "conditions": [
            {
              "type": "frontmost_application_if",
              "bundle_identifiers": ["^com\\.apple\\.Terminal"]
            }
          ]
        }
      ]
    }
  ]
}
```

### 3

jis, ansi キーボードの設定

input

```yaml
:title:
  - UK→US Mac keyboard
:maintainers:
  - foo
:rules:
  - :description: "右Ctrl+1を押すと:bow:を入力する"
    :from:
      - 1
      - right_control
    :to:
      - japanese_eisuu
      - " "
      - - ":"
        - left_shift
      - b
      - o
      - w
      - - ":"
        - left_shift
```

output

```json
{
  "title": "my config",
  "maintainers": ["foo"],
  "rules": [
    {
      "description": "右Ctrl+1を押すと:bow:を入力する",
      "manipulators": [
        {
          "conditions": [
            { "type": "keyboard_type_if", "keyboard_types": ["ansi"] }
          ],
          "from": {
            "key_code": "1",
            "modifiers": {
              "mandatory": ["right_control"],
              "optional": ["any"]
            }
          },
          "to": [
            { "key_code": "japanese_eisuu" },
            { "key_code": "spacebar" },
            { "key_code": "semicolon", "modifiers": ["left_shift"] },
            { "key_code": "b" },
            { "key_code": "o" },
            { "key_code": "w" },
            { "key_code": "semicolon", "modifiers": ["left_shift"] }
          ],
          "type": "basic"
        },
        {
          "conditions": [
            { "type": "keyboard_type_if", "keyboard_types": ["jis"] }
          ],
          "from": {
            "key_code": "1",
            "modifiers": {
              "mandatory": ["right_control"],
              "optional": ["any"]
            }
          },
          "to": [
            { "key_code": "japanese_eisuu" },
            { "key_code": "spacebar" },
            { "key_code": "quote" },
            { "key_code": "b" },
            { "key_code": "o" },
            { "key_code": "w" },
            { "key_code": "quote" }
          ],
          "type": "basic"
        }
      ]
    }
  ]
}
```
