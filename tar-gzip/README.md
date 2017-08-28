アーカイブ・圧縮・解凍をする
====

## tar.gzについて

圧縮ファイルに`tar.gz`がある。`tar`,`gz`はそれぞれ意味がことなる。

### tar とは

> TARとは、複数のファイルを一つにまとめて格納するアーカイブ(書庫)ファイルの形式の一つ。また、同形式のファイルの操作を行うためのUNIX系OS標準のコマンドおよびプログラムの名前。TAR形式の書庫ファイルは「tarball」(ターボール)とも呼ばれ、ファイル名末尾の標準の拡張子は「.tar」。
> http://e-words.jp/w/TAR.html

**複数のファイルを1つにまとめる機能**であり、圧縮機能はない。


### gzip とは

> GZIPとは、汎用のデータ圧縮方式および圧縮ファイル形式の一つ。ファイル名の標準の拡張子は「.gzip」または「.gz」。また、主にUNIX系OSでよく用いられる、同形式によるファイルの圧縮・伸張を行うプログラムおよびコマンド名。
> http://e-words.jp/w/GZIP.html

**ファイルを圧縮する機能**であり、アーカイブ（書庫）機能はない。


### tar.gz とは

tarでファイルを1つにまとめ、gzで圧縮したファイルということ。


## アーカイブ・圧縮

### tarの仕組み

tarは **TA**pe **A**rchiveの略。

tarアーカイブファイルは、headerブロックとcontentブロックからなる。複数ファイルをまとめるときにそれぞれのファイルにheaderを付けてアーカイブする。

|ブロック|サイズ(byte)|
|:---|---:|
|header1|512|
|content1|...|
|header2|512|
|content2|...|
|...|...|

headerブロックには以下のフィールドがある。

* name
* mode
   * setUID, setGID, TSVTXのモードがある
* uid
   * ユーザID
* gid
   * グループID
* size
   * ファイルサイズ
* mtime
   * 最終更新日
* chksum
* flagtype
   * アーカイブの種類を決定する
* linkname
* magic
* version
* uname
* gname
* devmajor
* devminor 
* prefix

[tar の構造](http://www.redout.net/data/tar.html)



## 展開・解凍