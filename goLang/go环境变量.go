/*

The GOPATH environment variable lists places to look for Go code. On Unix, the value is a colon-separated string. On Windows, the value is a semicolon-separated string. On Plan 9, the value is a list.

colon-separated string 代表:
semicolon-separated string 代表;

GOPATH must be set to get, build and install packages outside the standard Go tree.
当时看到gopath 我就很快的想到了class_path 因为他们真的比较像
尤其是当初在windows设置class_path环境变量的时候，也是用;作为分隔符



可以把这个目录理解为工作目录
GOPATH is a variable that defines the root of your workspace. By default, the workspace directory is a directory that is named go within your user home directory (~/go for Linux and MacOS, %USERPROFILE%/go for Windows). GOPATH stores your code base and all the files that are necessary for your development. You can use another directory as your workspace by configuring GOPATH for different scopes. GOPATH is the root of your workspace and contains the following folders:

src/: location of Go source code (for example, .go, .c, .g, .s).

pkg/: location of compiled package code (for example, .a).

bin/: location of compiled executable programs built by Go.

*/

/**

GOROOT is a variable that defines where your Go SDK is located. You do not need to change this variable, unless you plan to use different Go versions.

**/


/**
go env查看go的环境变量
**/

main() {
	print("hello")
}