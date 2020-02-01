#!/bin/sh

DIR=`dirname "$0"`
FILE=bundled.go
BIN=`go env GOPATH`/bin

cd $DIR
rm $FILE

$BIN/fyne bundle -package data -name pencil pencil.svg > $FILE
$BIN/fyne bundle -package data -append -name dropper dropper.svg >> $FILE
