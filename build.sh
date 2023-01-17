mkdir tmp
CGO_ENABLED=1 go build
mv Aws-Panel tmp/
cp ./example/config.json tmp/
cp LICENSE tmp/
cd ./web/
npm install
npm run build
mv dist/ ../tmp/web
cd ../tmp/
zip -r Aws-Panel.zip *
mv Aws-Panel.zip ../
cd ../
rm -rf tmp
