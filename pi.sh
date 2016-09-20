#!/usr/bin/env bash
set -e

echo " - Compiling..."
(cd client; npm install && gulp && rm -rf ../bin/dist && mv -f dist/ ../bin/) &> /dev/null

echo " - Packaging..."
mkdir -p pi pi/server
cp -a bin/. pi/bin
cp server/main.go pi/server/main.go
cp residents.yaml pi/residents.yaml
echo "#!/usr/bin/env bash\n./bin/server" > ./pi/run.sh

echo " - Compressing..."
tar -czf pi.tar.gz ./pi
rm -r ./pi
