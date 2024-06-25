# Get the version.
version=`git describe --tags --abbrev=0`
# Write out the package.
cat << EOF > version.go
package main
//go:generate bash ./get_version.sh
var version = "$version"
EOF