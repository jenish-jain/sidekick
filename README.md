# Your own sidekick

### how to run it locally
> get random dad joke 
`` go run main.go dadjoke``


### how to use the tool from homebrew
````shell
brew tap jenish-jain/tap
brew install jenish-jain/tap/sidekick
````

### what can your sidekick do?

* Tell you a dad joke. 😉
```shell
sidekick dadjoke
```
* manage your 2FAs 💁‍ > [doc](internal/twoFA/Readme.md)
```shell
sidekick 2fa --list
```

### see the tool in action
![sidekick-demo](assets/sidekick-demo.gif)

## release steps

```shell
git tag -a x.x.x commitId -m "Message"
git push origin x.x.x
```
