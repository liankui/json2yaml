# json2yaml

## how to use
1. in terminal, use the command: `go run .`
2. open a browser, type in http://127.0.0.1:21111 and press enter
3. enjoy it

## use in docker
1. in terminal, use this command: `sh build.sh`, and done!

### in this command
You will get a docker image: json2yaml:v1 and a running container.

You can use `Ctrl + C` to shutdown the running container.

Or use `docker run -it --rm -p 21111:21111 json2yaml:v1` to run it again.

### in different platform
If you use it in different platform, you can edit build.sh, 

you can use --platform to specify different platform, 

such as --platform=linux/amd64(linux/arm64,darwin/arm64,windows/amd64)
