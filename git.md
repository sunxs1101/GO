

## git学习

git pull https://github.com/k8sp/auto-install.git master
  
 ```
 git clone https://github.com/k8sp/auto-install.git
 git branch
 git checkout tftp_conf不可行
 解决方法：
 git remote -v
 git status
 然后git checkout tftp_conf
 cp /home/crawler/work/src/tftp/tftp
 git status
 git add *
 git commit -m "update tftp"
 git config --global user.email "you@example.com"
 git config --global user.name "Your Name"
 git commit -m "update tftp"
 git status
 git merge master
 git commit -m "mv to  bootstrapper/tftp/"
 git push
 ```
 步骤：
 ```
 pull master
 新建分支：checkout -a
 在相应目录下操作
 git commit
 git push
 ```
 ```
 进入git clone的目录，git pull获取最新的状态
 ```
 ```
 git checkout -b wangkuiyi-master master
 git branch -a
 git checkout master
 get merge --no-ff wangkuiyi-master
 git merge --no-ff wangkuiyi-master
 ```
 ### git merge与git rebase
 这两者是一回事，
 ```
 将master分支合并到feature分支
 git checkout feature
 git merge master
 可以压缩到一起
 git merge master feature
 ```
 ```
 git checkout feature
 git rebase master
 ```
 
 
 
