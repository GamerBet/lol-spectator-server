#!/usr/bin/env bash

echo -----------------------
echo Spectate by op.gg
echo -----------------------


cd "/Applications/League of Legends.app/Contents/LoL/RADS/projects/lol_air_client/releases/"

for CLIENT_PATH in `ls -v`
do
if test -d $CLIENT_PATH
then
break
fi
done

echo "$CLIENT_PATH"


cd "/Applications/League of Legends.app/Contents/LoL/RADS/solutions/lol_game_client_sln/releases/"

for LAUNCHER_PATH in `ls -v`
do
if test -d $LAUNCHER_PATH
then
break
fi
done

echo "$LAUNCHER_PATH"


cd "$LAUNCHER_PATH/deploy/LeagueOfLegends.app/Contents/MacOS"

	riot_launched=true "/Applications/League of Legends.app/Contents/LoL/RADS/solutions/lol_game_client_sln/releases/$LAUNCHER_PATH/deploy/LeagueOfLegends.app/Contents/MacOS/LeagueofLegends" 8394 LoLLauncher "/Applications/League of Legends.app/Contents/LoL/RADS/projects/lol_air_client/releases/$CLIENT_PATH/deploy/bin/LolClient" "spectator {{YOUR_SERVER}} {{GAME_ENCRYPTION_KEY}} {{GAME_ID}} {{GAME_REGION}}"
