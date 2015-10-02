# lol-spectator-server

REST Server for League Of Legends spectator

### Usage

```
$ go run server.go --port=5000
```

or

```
$ go build
$ ./lol-spectator-server --port=5000
```

### Specification

Currently this [gist](https://gist.github.com/themasch/8375971) is the best written specification we've seen for the Riot Spectator API.


### Running the League of Legends client in spectator mode

- [Windows](https://developer.riotgames.com/docs/spectating-games)

```
$ "C:\Riot Games\League of Legends\RADS\solutions\lol_game_client_sln\releases\{{CURRENT_RELEASE}}\
deploy\League of Legends.exe" "8394" "LoLLauncher.exe" "" "spectator
{{YOUR_SERVER}} {{GAME_ENCRYPTION_KEY}} {{GAME_ID}} {{GAME_REGION}}"
```

- [Mac](https://github.com/GamerBet/lol-observer-server/blob/master/clientfiles/mac.command)

```
"/Applications/League of Legends.app/Contents/LoL/RADS/solutions/lol_game_client_sln/releases/$LAUNCHER_PATH/deploy/LeagueOfLegends.app/Contents/MacOS/LeagueofLegends" 8394 LoLLauncher "/Applications/League of Legends.app/Contents/LoL/RADS/projects/lol_air_client/releases/$CLIENT_PATH/deploy/bin/LolClient" "spectator {{YOUR_SERVER}} {{GAME_ENCRYPTION_KEY}} {{GAME_ID}} {{GAME_REGION}}"
```

- [Linux](https://www.playonlinux.com/en/topic-12991-Arguments_too_long__Help.html)

```
$ playonlinux-bash ~/.PlayOnLinux/shortcuts/"LoL Spectator" "8394" "LoLLauncher.exe" "" "spectator {{YOUR_SERVER}} {{GAME_ENCRYPTION_KEY}} {{GAME_ID}} {{GAME_REGION}}"
```



