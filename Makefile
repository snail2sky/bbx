mac:
		export GOOS=darwin && go build

linux:
		export GOOS=linux && go build

windows:
		export GOOS=windows && go build

help:
		echo "make [ mac | linux | windows ]"
