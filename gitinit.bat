@echo off
git init -b main
git config user.name "mechiko"
git config user.email "kbprime@gmail.com"
git config core.filemode false
git config core.autocrlf false
git config --global push.autoSetupRemote true
git branch -M main
rem Add origin if missing; otherwise update it
git remote add origin git@github.com:mechiko/kituwalk.git
