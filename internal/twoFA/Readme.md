# 2FA

2fa is a two-factor authentication agent.

Inspired from: https://github.com/rsc/2fa/tree/master

`sidekick 2fa --add <name>` adds a new key to the 2fa keychain with the given name. 
It prints a prompt to standard error and reads a two-factor key from standard input. 
Two-factor keys are short case-insensitive strings of letters A-Z and digits 2-7.

By default, the new key generates time-based (TOTP) authentication codes
By default, the new key generates 6-digit codes

`sidekick 2fa --list` lists the names of all the keys in the keychain.

`sidekick 2fa <name>` prints a two-factor authentication code from the key with the given name.

`sidekick 2fa` print two-factor authentication codes for all keys registered

The default time-based authentication codes are derived from a hash of the key and the current time, so it is important that the system clock have at least one-minute accuracy.

The keychain is stored unencrypted in the text file $HOME/.2fa.

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