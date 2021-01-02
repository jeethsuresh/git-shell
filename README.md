# git-shell

This is a toy project to get a git shell running on a VM; I haven't figured out how to bootstrap the server yet, so for now we'll lean on github as a crutch. 

This project borrows heavily from go-git: https://github.com/go-git/go-git and its included cli/server application. However, that application is non-functional as provided and so I made an array of fixes. During adoption you'll want to create a folder in `/` called `git` for debug logs, since this thing runs as a shell we can't exactly print to the screen when something goes wrong. Remember to give the git user access to that folder, because if you do things right your ssh shell won't be running as root!

You can use this as a drop-in replacement for the original git-shell to clone/push, and use the official git-shell guide: https://git-scm.com/docs/git-shell to install it (simply clone and build this repo, move it somewhere the git user can access it, and then add it to the list of shells before following the rest of the guide).

One innovation I added, specifically for my own benefit, is the ability to initialize a new repo at the time of your first push. I have no desire to SSH into a VM every time I want to store a new codebase on my server, and I trust myself to stay disciplined about naming/duplicates. The changes required for that are in receive_pack.go if you want to revert them. 

TODO:
- something to replace gitlab's awesome CI/CD system
- Authn/Authz to support multi-tenancy
- Abstracting the repo name/location, also for multi-tenancy
- HTTP support 
- a web UI
- k8s
