#!/bin/bash

baseDir=$(pwd)
cd $baseDir/../cmd
go build main.go
mv main tgBot
mv tgBot $baseDir/../
