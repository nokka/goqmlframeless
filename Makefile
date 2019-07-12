# Use run if there's only changes to the QML layer.
run:
	./deploy/darwin/widget-test.app/Contents/MacOS/widget-test

# Rebuild the app if you made changes to the Go layer.
build:
	qtdeploy build