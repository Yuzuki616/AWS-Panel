mkdir tmp
CGO_ENABLED=1 go build || exit 1
mv Aws-Panel tmp/
cp ./example/config.json tmp/
cp LICENSE tmp/
cd ./web/ || exit 1
npm install || exit 1
npm run build || exit 1
mv dist/ ../tmp/web
cd ../tmp/ || exit 1
zip -r Aws-Panel.zip ./*
mv Aws-Panel.zip ../
cd ../
rm -rf tmp
