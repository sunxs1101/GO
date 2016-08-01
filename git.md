git学习

## git pull https://github.com/k8sp/auto-install.git master
  
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
 
