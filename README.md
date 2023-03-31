# report-binding

pdf記事を製本するコマンドラインアプリケーション。
用意した記事や表紙を結合して製本する。目次pdf生成機能付き。

## インストール
[release](https://github.com/kienn-HCl/report-binding/releases/latest)にバイナリがあるので各自の環境に合わせてインストールしてください。

## 使い方

### 製本の一通りの流れ
記事pdfのあるディレクトリで


```shell
$ ./report-binding i
```

とすることでディレクトリに`reportData.csv`,`FrontCover`,`TableOfContents`,`UnitedReport`,`BackCover`が生成されます。


```shell
$ lsd --tree
 .
├──  article1.pdf
├──  article2.pdf
├──  article3.pdf
└──  report-binding

$ ./report-binding i

$ lsd --tree
 .
├──  article1.pdf
├──  article2.pdf
├──  article3.pdf
├──  BackCover
├──  FrontCover
├──  report-binding
├──  reportData.csv
├──  TableOfContents
└──  UnitedReport

```

`reportData.csv`はpdfのメタデータを読み取って作られます。読み取れなかった部分は空となります。

```
┌───────────┬────────┬─────────────┬──────────────┐
│ PageCount │ Author │ Title       │ Filename     │
├───────────┼────────┼─────────────┼──────────────┤
│ 1         │        │             │ article1.pdf │
│ 1         │ int    │ てすと記事2 │ article2.pdf │
│ 1         │ string │ てすと記事3 │ article3.pdf │
└───────────┴────────┴─────────────┴──────────────┘
```

この`reportData.csv`の行の順番を変更することで結合したときの記事の順番も変更することが出来ます。

`FrontCover`に表紙のpdfを、`TableOfContents`に目次のpdfを、`BackCover`に裏表紙のpdfを一枚入れた状態で

```shell
$ ./report-binding b
```

とすることで`BoundReport`ディレクトリ内に製本されたpdfが作られます。なお、製本されたpdfはページ数が4の倍数となるように調整されます。

```shell
$ lsd --tree
 .
├──  article1.pdf
├──  article2.pdf
├──  article3.pdf
├──  BackCover
│   └──  backCover.pdf
├──  FrontCover
│   └──  frontCover.pdf
├──  report-binding
├──  reportData.csv
├──  TableOfContents
│   └──  tableOfContents.pdf
└──  UnitedReport

$ ./report-binding b

$ lsd --tree
 .
├──  article1.pdf
├──  article2.pdf
├──  article3.pdf
├──  BackCover
│   └──  backCover.pdf
├──  BoundReport
│   └──  boundReport.pdf
├──  FrontCover
│   └──  frontCover.pdf
├──  report-binding
├──  reportData.csv
├──  TableOfContents
│   └──  tableOfContents.pdf
└──  UnitedReport
    └──  unitedReport.pdf
```

### 目次生成機能
`reportData.csv`のデータをもとに目次pdfを生成できます。データに空の部分があってはいけません。
```
┌───────────┬────────┬─────────────┬──────────────┐
│ PageCount │ Author │ Title       │ Filename     │
├───────────┼────────┼─────────────┼──────────────┤
│ 1         │ void   │ てすと記事1 │ article1.pdf │
│ 1         │ int    │ てすと記事2 │ article2.pdf │
│ 1         │ string │ てすと記事3 │ article3.pdf │
└───────────┴────────┴─────────────┴──────────────┘
```
以下のコマンドを記事と`reportData.csv`があるディレクトリで実行することで、`FrontCover/frontCover.pdf`が作成されます。

```
$ ./report-binding t
```

### ヘルプ
引数などをつけずに実行することでヘルプを見れます。

## font
このプログラムでは目次pdf作成時にIPAフォントを使用しています。IPAフォントのライセンスはfontディレクトリ内にあります。

