# GoTV - watching Weiqi like TV

currently this is just a brain dump of implementation parts

[![Build Status](https://travis-ci.org/dgf/gotv.svg)](https://travis-ci.org/dgf/gotv)

## model

contains basic Go model types

### Board

supports a straight forward Go game with capturing and KO + suicide detection

## sgf

parse SGF into generic game tree collection

lexer is generated with [Nex], see also [Structural Regular Expressions]

```sh
go get github.com/blynn/nex
go generate
```

[Nex](https://github.com/blynn/nex)
[Structural Regular Expressions](http://doc.cat-v.org/bell_labs/structural_regexps/se.pdf)
