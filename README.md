# atcoder-cui
solve atcoder contest problems in cui  
以下の4コマンドを覚えるだけで、自身が用意したAtCoder用のテンプレートを用いて、AtCoderのコンテストに参加できる。

1. `/bin/prepare`: コンテストの準備  
   `/bin/prepare abc272 golang go run (${コンテスト名} ${使用する言語} ${実行方法(hogehoge main.xxx のhogehogeの部分)})`  
   でそのコンテストを何言語で解くかを決める
2. `/bin/solve`: コンテストの問題を解く  
   `/bin/solve`を問題を解き始めるときに実行することで、`/templates/yourlanguage/main.xxx`を`/solve/main.〇〇`にコピーし、`main.○○`に解法を書いてゆく
3. `/bin/test`: コンテストの問題をテストする  
   `/bin/test a (${問題名})` で`/solve/main.○○`のコードを指定した問題の入出力例に対してテスト
4. `/bin/submit`: コンテストの問題を提出する  
   `/bin/submit a (${問題名})`で`/solve/main.○○`のコードを提出

## Installation & How to Use

0. `online-judge-tools`と`atocoder-cli`をインストールし、`oj`コマンド、`acc`コマンドが使える状態にしておく
1. このリポジトリをクローン
2. templatesの中に`templates/langugageName/template.xxx`となるようにテンプレートを使用する。  
   `langugageName`は自由に決めることができ、`/bin/prepare`を使うときの第二引数となる
   （例 `templates/hoge/template.hogehoge`は、`bin/prepare abcXXX hoge` の後 `bin/solve`とすることで`template.hogehoge`が`solve/main.hogehoge`にコピーされ、解く準備が完了する）
3. bin/prepareを実行する。各引数は以下のように  
   - 第一引数: コンテスト名（例: abc274）
   - 第二引数: 使用する言語（手順2. で指定したディレクトリ名）
   - 第三引数: コードを実行するときのコマンド(ファイル名部分を除く)
     - Python: `python main.py`なので`python`が第三引数となる
     - golang: `go run main.go`なので`go run`が第三引数となる
     - ファイル名が末尾に来ない実行方法に対応していない。
