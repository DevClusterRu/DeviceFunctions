git pull origin master
git add .
git commit -m "AutoBuild"
git push origin master

git tag $1
git push $1