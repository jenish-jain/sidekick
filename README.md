# Your own sidekick
Like Jarvis is to Ironman, robin is to batman.
This is your trusty sidekick here to handle your itty-bitty chores and
sprinkle in a bit of entertainment whenever you fancy.
You've got yourself a dynamic duo, honey! 💁‍♀️️💥

### how to run it locally
get random dad joke 
````shell 
go run main.go dadjoke
````


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

tagging and release are automated 🚀
* follow these committing [guidelines](https://github.com/anothrNick/github-tag-action#options) 

