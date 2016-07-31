## The Problem

If Go code is in a private Github repo, we wouldn't be able to run `go get github.com/account/repo` to retrieve the code. It is true that there are workarounds like those use Github SSH keys, but which would introduce other problems like `go get -u` doesn't work. 

## The Solution

The easiest solution is to

1. generate a personal token from your Github account which has the right to access the private repo, and
2. tell git, which is called by `go get`, to access the repo via `https://<token>:github.com/account/repo` instead of via `https://github.com/account/repo`.

### Generate Personal Token

1. Go to the Web page of Github, login, and go to Github settings page.  At the "personal access tokens" tab, click the button "generate personal access token".

    <img width="1147" alt="screen shot 2016-07-26 at 12 42 34 pm" src="https://cloud.githubusercontent.com/assets/1548775/17153116/cea6c938-532f-11e6-9ac1-888c848d88b2.png">

1. Select "repo" in "scopes":

    <img width="1147" alt="screen shot 2016-07-26 at 12 44 40 pm" src="https://cloud.githubusercontent.com/assets/1548775/17153163/18d983f6-5330-11e6-80e4-c344e80bdf5a.png">

1. Copy and paste the generated token number, and save it somewhere you won't forget.  An unsafe and reasonable place is your email box.

### Tell Git About Your Token

Run the following command 

```
git config --global url."https://c61axxxxxxxxxxxxxxx:x-oauth-basic@github.com/".insteadOf "https://github.com/"
```

where `c61axxxxxxxxxxxxxxx` is your Github personal access token.

Then you will find something new was added to your `~/.gitconfig` file:

```
[url "https://c61axxxxxxxxxxxxxxx:x-oauth-basic@github.com/"]
	insteadOf = https://github.com/
```

### Test It!

Now, run `go get github.com/k8sp/auto-install/cloud-config-server`.
