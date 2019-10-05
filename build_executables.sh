#!/usr/bin/env bash

package="github.com/mutl3y/PRTG_XMR-STAK"
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("windows/amd64" "windows/386" "darwin/amd64" "linux/amd64")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=releases/$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution…'
        exit 1
    fi
done

cp releases/PRTG_XMR-STAK-windows-amd64.exe "/c/Program Files (x86)/PRTG Network Monitor/Custom Sensors/EXEXML/"

function myscp() {
     scp -i /c/Users/mark/.ssh/mark releases/PRTG_XMR-STAK-linux-amd64 mark@linuxserver:/var/prtg/scriptsxml
 }
 tries=0
#ssh -i /c/Users/mark/.ssh/mark mark@linuxserver sudo rm /var/prtg/scriptsxml/PRTG_XMR-STAK-linux-amd64

myscp; while [ $? -ne 0 ]
 do
 sleep 0.5
 myscp
 let tries=tries+1
if  [ ${tries} -ge 20 ]; then
    break
fi
 done