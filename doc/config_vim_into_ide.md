# Config vim to Golang-IDE in linux 

https://my.oschina.net/u/3287304/blog/1559952


1. upgrade vim  to 8.0 latest

[root@cloud-sz-kolla-b13-01 yum.repos.d]# cat vim.repo 

[vim-latest]
name=vim 
baseurl=http://mirror.ghettoforge.org/distributions/gf/el/7/plus/x86_64/
enabled=1
skip_if_unavailable=1
gpgcheck=0

yum install -y cmake vim 


2. modify .vimrc for plugin  


3. 





# Issue 

https://github.com/todaygood/go_repo/issues/6

https://github.com/todaygood/go_repo/issues/5




