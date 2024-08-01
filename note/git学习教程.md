# 20240801开始学习git--0802

分支管理:![ ](/Image/GIT分支管理.jpg)

## 当文件提交有冲突
分支管理:![ ](/Image/git提交冲突.jpg)
1. git checkout -- file
2. git add file
3. git commit -m "message"

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