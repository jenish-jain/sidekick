# 2FA

2fa is a two-factor authentication agent.

Inspired from: https://github.com/rsc/2fa/tree/master

Learn more about this: https://www.rfc-editor.org/rfc/rfc4226

| Flag Name | Short Hand | Description                                                                                                                                                                                                                                                                                                                                                         | Usage                                                                                        |
|:----------|:----------:|:--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|:---------------------------------------------------------------------------------------------|
| --add     |     -a     | adds a new key to the 2fa keychain with the given name.<br>It prints a prompt to standard error and reads a two-factor key from standard input.Two-factor keys are short case-insensitive strings of letters A-Z and digits 2-7. <br>By default, the new key generates time-based (TOTP) authentication codes . <br>By default, the new key generates 6-digit codes | <ul><li>`sidekick 2fa --add <name>`</li> <li>`sidekick 2fa -a <name>`</li></ul>              |
| --hotp    |     -H     | this makes the new key generated counter-based (HOTP) codes instead.                                                                                                                                                                                                                                                                                                | <ul><li>`sidekick 2fa --add --hotp <name>` </li> <li> `sidekick 2fa -a -H <name>`</li></ul>  |
| --seven   |     -7     | the -7 flag select 7-digit codes instead.                                                                                                                                                                                                                                                                                                                           | <ul><li>`sidekick 2fa --add --seven <name>` </li> <li> `sidekick 2fa -a -7 <name>`</li></ul> |
| --eight   |     -8     | the -8 flag select 8-digit codes instead                                                                                                                                                                                                                                                                                                                            | <ul><li>`sidekick 2fa --add --eight <name>` </li> <li> `sidekick 2fa -a -8 <name>`</li></ul> |
| --list    |     -l     | lists the names of all the keys in the keychain.                                                                                                                                                                                                                                                                                                                    | <ul><li>`sidekick 2fa --list`</li><li>`sidekick 2fa -l`</li></ul>                            |
| --clip    |     -c     | copies two-factor authentication for the specified name to your clipboard                                                                                                                                                                                                                                                                                           | <ul><li>`sidekick 2fa --clip <name>`</li><li>`sidekick 2fa -c <name>`</li></ul>              |
| NONE      |     NA     | prints a two-factor authentication code for all key or specified name.                                                                                                                                                                                                                                                                                              | <ul><li>`sidekick 2fa <name>`</li><li>`sidekick 2fa`</li></ul>                               |   

* The default time-based authentication codes are derived from a hash of the key and the current time, so it is important that the system clock have at least one-minute accuracy.
* while displaying the code, it also displays the sec code can be refreshed next in, -1 indicated no expiry till next usage
* The keychain is stored unencrypted in the text file $HOME/.2fa.

## Example

During GitHub 2FA setup, at the “Scan this barcode with your app” step,
click the “enter this text code instead” link.
A window pops up showing “your two-factor secret,” a short string of letters and digits.

Add it to 2FA under the name github, typing the secret at the prompt:
```shell
sidekick 2fa --add github
2fa key for github: nzxxiidbebvwk6jb
```

Then whenever GitHub prompts for a 2FA code, run 2fa to obtain one:

```shell 
sidekick 2fa github
268346
``` 