# hicli
hicli is an open source, community-driven command line client for managing hi devops servers.

hicli is writen in Go, it can be run cross Windows, Mac, Linux.

```bash
hi cicd run --app=my-app --project=my-project --profile=test
```

## Git workflow

Below, we outline one of the more common Git workflows that core developers use. Other Git workflows are also valid.

### Fork the main repository

* Go to https://github.com/hidevopsio/hicli
* Click the "Fork" button (at the top right)

### Clone your fork

The commands below require that you have $GOPATH set ($GOPATH docs). We highly recommend you put Istio's code into your GOPATH. Note: the commands below will not work if there is more than one directory in your $GOPATH.

```bash
export GITHUB_USER=your-github-username
mkdir -p $GOPATH/src/github.com/hidevopsio
cd $GOPATH/src/github.com/hidevopsio
git clone https://github.com/$GITHUB_USER/hicli
cd hicli
git remote add upstream 'https://github.com/hidevopsio/hicli'
git config --global --add http.followRedirects 1
```
