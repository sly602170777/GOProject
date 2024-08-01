# 20240801开始学习git--20240801和0802

分支管理:![ ](/Image/GIT分支管理.jpg)

## 当文件提交有冲突
提交解决冲突:![ ](/Image/git提交冲突.jpg)
1. git checkout -- file
2. git add file
3. git commit -m "message"

## 当文件相同的地方有冲突，需要手动解决冲突
手动解决冲突:![ ](/Image/手动解决冲突.jpg)
解决方法：当本地文件已经修改的文件有冲突时，将本文件copy一个副本，
然后将源文件内容revert到最新，然后打开副本文件，将冲突部分手动解决，然后add和commit

## git log
git log --pretty=oneline

## git reflog
git reflog

## git reset --hard HEAD^
git reset --hard HEAD~1
git reset --hard commit_id

## git pull
git pull = git fetch + git merge

## git fetch
git fetch origin master:tmp
git diff tmp
git merge tmp

## git stash
git stash
git stash pop

## git cherry-pick
git cherry-pick commit_id

## git reset
git reset --hard HEAD^
git reset --hard commit_id

## git revert
git revert commit_id

## git reflog
git reflog

## git branch
git branch branch_name