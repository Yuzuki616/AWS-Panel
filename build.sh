mkdir tmp
CGO_ENABLED=1 go build
mv Aws-Panel tmp/
cd ./web/
npm install
npm run build
mv dist/ ../tmp/web
cd ../tmp/
zip -r Aws-Panel.zip *
mv Aws-Panel.zip ../
cd ../
rm -rf tmp
