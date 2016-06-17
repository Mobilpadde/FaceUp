# FaceUp

Want to see a [demo](https://gfycat.com/LoathsomePoorErmine)?

Let a specific hue light up whenever someone writes you at [Messenger](http://messenger.com).

It's pretty easy to use: Simply fill out `config.template.json`, save it as `config.json`, build and run `main`.

Alternatively, if you do not have go, you can download the [w64 version](https://github.com/Mobilpadde/FaceUp/releases), but still you'll have to fill out the `config.json`.

As of [v0.0.3](https://github.com/Mobilpadde/FaceUp/releases) you don't even have to fill out the config; simply use the `g`-flag and type in your information :D

## config.json

An explanation of `config.json` is of course in order:

 * `UserId`: Your facebook user id; used to not make it blink at your own messages
 * `Mail`: Your facebook e-mail
 * `Pass`: Your facebook password (Don't worry, it's only used to login to your facebook-account)
 * `Name`: Name of the hue you want to light up
 * `Speed`: How fast the light should blink
 * `Hueser`: The user of the bridge

If you'd like to check what have changed during the versions, feel free to checkout the [changelog](changelog.md).