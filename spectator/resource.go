package spectator

import (
	"encoding/json"
	"fmt"
	"os/user"
	"strconv"
)

type ChunkInfo struct {
	ChunkId            int `json:"chunkId"`
	AvailableSince     int `json:"availableSince"`
	NextAvailableChunk int `json:"nextAvailableChunk"`
	KeyFrameId         int `json:"keyFrameId"`
	NextChunkId        int `json:"nextChunkId"`
	EndStartupChunkId  int `json:"endStartupChunkId"`
	StartGameChunkId   int `json:"startGameChunkId"`
	EndGameChunkId     int `json:"endGameChunkId"`
	Duration           int `json:"duration"`
}

type MetaData struct {
	GameKey struct {
		GameId     int    `json:"gameId"`
		PlatformId string `json:"platformId"`
	} `json:"gameKey"`
	GameServerAddress         string `json:"gameServerAddress"`
	Port                      int    `json:"port"`
	EncryptionKey             string `json:"encryptionKey"`
	ChunkTimeInterval         int    `json:"chunkTimeInterval"`
	StartTime                 string `json:"startTime"`
	GameEnded                 bool   `json:"gameEnded"`
	LastChunkId               int    `json:"lastChunkId"`
	LastKeyFrameId            int    `json:"lastKeyFrameId"`
	EndStartupChunkId         int    `json:"endStartupChunkId"`
	DelayTime                 int    `json:"delayTime"`
	PendingAvailableChunkInfo []struct {
		Duration     int    `json:"duration"`
		Id           int    `json:"id"`
		ReceivedTime string `json:"receivedTime"`
	} `json:"pendingAvailableChunkInfo"`
	PendingAvailableKeyFrameInfo []struct {
		Id           int    `json:"id"`
		ReceivedTime string `json:"receivedTime"`
		NextChunkId  int    `json:"nextChunkId"`
	} `json:"pendingAvailableKeyFrameInfo"`
	KeyFrameTimeInterval      int    `json:"keyFrameTimeInterval"`
	DecodedEncryptionKey      string `json:"decodedEncryptionKey"`
	StartGameChunkId          int    `json:"startGameChunkId"`
	GameLength                int    `json:"gameLength"`
	ClientAddedLag            int    `json:"clientAddedLag"`
	ClientBackFetchingEnabled bool   `json:"clientBackFetchingEnabled"`
	ClientBackFetchingFreq    int    `json:"clientBackFetchingFreq"`
	InterestScore             int    `json:"interestScore"`
	FeaturedGame              bool   `json:"featuredGame"`
	CreateTime                string `json:"createTime"`
	EndGameChunkId            int    `json:"endGameChunkId"`
	EndGameKeyFrameId         int    `json:"endGameKeyFrameId"`
	FirstChunkId              int    `json:"firstChunkId"`
}

func makePath(platformId, gameId string) string {
	usr, _ := user.Current()
	return fmt.Sprintf("%s/replays/%s/%s%s/%s/%s/%s", usr.HomeDir, platformId, gameId[0:1], gameId[1:2], gameId[2:3], gameId[3:4], gameId)
}

func group(byteChan <-chan []byte) []byte {
	result := []byte{}
	for chunk := range byteChan {
		result = append(result, chunk...)
	}

	return result
}

// Lists the 10 featured games for the regions supported by this server.
func featured() []byte {
	return []byte(`{"gameList":[{"gameId":1230560002,"mapId":11,"gameMode":"CLASSIC","gameType":"MATCHED_GAME","gameQueueConfigId":4,"participants":[{"teamId":100,"spell1Id":12,"spell2Id":14,"championId":105,"skinIndex":8,"profileIconId":767,"summonerName":"Jurrog","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":14,"championId":26,"skinIndex":0,"profileIconId":907,"summonerName":"HungrypandaGG","bot":false},{"teamId":100,"spell1Id":11,"spell2Id":4,"championId":64,"skinIndex":0,"profileIconId":711,"summonerName":"Naande","bot":false},{"teamId":100,"spell1Id":7,"spell2Id":4,"championId":18,"skinIndex":2,"profileIconId":758,"summonerName":"Bananitoo","bot":false},{"teamId":100,"spell1Id":14,"spell2Id":4,"championId":38,"skinIndex":4,"profileIconId":7,"summonerName":"MrMgr","bot":false},{"teamId":200,"spell1Id":7,"spell2Id":4,"championId":42,"skinIndex":6,"profileIconId":552,"summonerName":"niszczyciel","bot":false},{"teamId":200,"spell1Id":14,"spell2Id":4,"championId":39,"skinIndex":3,"profileIconId":750,"summonerName":"Wolszczuk","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":14,"championId":1,"skinIndex":1,"profileIconId":682,"summonerName":"Air4","bot":false},{"teamId":200,"spell1Id":14,"spell2Id":4,"championId":412,"skinIndex":1,"profileIconId":873,"summonerName":"eS TX","bot":false},{"teamId":200,"spell1Id":11,"spell2Id":4,"championId":33,"skinIndex":6,"profileIconId":0,"summonerName":"SoT Drago","bot":false}],"observers":{"encryptionKey":"x4llxUH96TP6VUZHMbxZwYotxVWB1MJk"},"platformId":"EUN1","gameTypeConfigId":2,"bannedChampions":[{"championId":112,"teamId":100,"pickTurn":1},{"championId":107,"teamId":200,"pickTurn":2},{"championId":114,"teamId":100,"pickTurn":3},{"championId":28,"teamId":200,"pickTurn":4},{"championId":72,"teamId":100,"pickTurn":5},{"championId":131,"teamId":200,"pickTurn":6}],"gameStartTime":1440434731522,"gameLength":23},{"gameId":1230550248,"mapId":11,"gameMode":"CLASSIC","gameType":"MATCHED_GAME","gameQueueConfigId":4,"participants":[{"teamId":100,"spell1Id":4,"spell2Id":11,"championId":64,"skinIndex":4,"profileIconId":618,"summonerName":"Thir13en Sins","bot":false},{"teamId":100,"spell1Id":14,"spell2Id":12,"championId":105,"skinIndex":0,"profileIconId":0,"summonerName":"hello im funzkai","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":14,"championId":412,"skinIndex":0,"profileIconId":561,"summonerName":"DingzzZ","bot":false},{"teamId":100,"spell1Id":7,"spell2Id":4,"championId":67,"skinIndex":8,"profileIconId":685,"summonerName":"Gabend","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":1,"championId":268,"skinIndex":0,"profileIconId":592,"summonerName":"Simpli","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":12,"championId":41,"skinIndex":7,"profileIconId":509,"summonerName":"TAP xFoRkAn","bot":false},{"teamId":200,"spell1Id":7,"spell2Id":4,"championId":222,"skinIndex":0,"profileIconId":719,"summonerName":"Dark King 87","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":14,"championId":50,"skinIndex":0,"profileIconId":575,"summonerName":"Alexei","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":11,"championId":11,"skinIndex":4,"profileIconId":898,"summonerName":"SevÃ©ns","bot":false},{"teamId":200,"spell1Id":14,"spell2Id":4,"championId":111,"skinIndex":0,"profileIconId":898,"summonerName":"kokichat","bot":false}],"observers":{"encryptionKey":"4xmhJQXsS0rJeRGnYWOWBL8UfqxubFQo"},"platformId":"EUN1","gameTypeConfigId":2,"bannedChampions":[{"championId":39,"teamId":100,"pickTurn":1},{"championId":72,"teamId":200,"pickTurn":2},{"championId":107,"teamId":100,"pickTurn":3},{"championId":82,"teamId":200,"pickTurn":4},{"championId":112,"teamId":100,"pickTurn":5},{"championId":10,"teamId":200,"pickTurn":6}],"gameStartTime":1440434414868,"gameLength":340},{"gameId":1230559292,"mapId":11,"gameMode":"CLASSIC","gameType":"MATCHED_GAME","gameQueueConfigId":4,"participants":[{"teamId":100,"spell1Id":11,"spell2Id":4,"championId":60,"skinIndex":2,"profileIconId":786,"summonerName":"strycek Redovski","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":7,"championId":18,"skinIndex":6,"profileIconId":685,"summonerName":"PotÄ™Å¼ny Aven","bot":false},{"teamId":100,"spell1Id":3,"spell2Id":4,"championId":157,"skinIndex":2,"profileIconId":666,"summonerName":"Extra Salty","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":14,"championId":412,"skinIndex":3,"profileIconId":595,"summonerName":"NLB Yen","bot":false},{"teamId":100,"spell1Id":12,"spell2Id":4,"championId":150,"skinIndex":1,"profileIconId":745,"summonerName":"Doxxy The Cat","bot":false},{"teamId":200,"spell1Id":14,"spell2Id":4,"championId":111,"skinIndex":0,"profileIconId":782,"summonerName":"Natsumi","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":12,"championId":84,"skinIndex":7,"profileIconId":508,"summonerName":"InnomÃ­nate","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":3,"championId":268,"skinIndex":1,"profileIconId":758,"summonerName":"AG Lee Paulixs","bot":false},{"teamId":200,"spell1Id":11,"spell2Id":4,"championId":421,"skinIndex":0,"profileIconId":539,"summonerName":"xxLOKIx","bot":false},{"teamId":200,"spell1Id":7,"spell2Id":4,"championId":236,"skinIndex":0,"profileIconId":781,"summonerName":"SÃ©LL","bot":false}],"observers":{"encryptionKey":"oih8na1dCD0vfGBcg6SCsifi9FXAKPEs"},"platformId":"EUN1","gameTypeConfigId":2,"bannedChampions":[{"championId":107,"teamId":100,"pickTurn":1},{"championId":105,"teamId":200,"pickTurn":2},{"championId":120,"teamId":100,"pickTurn":3},{"championId":72,"teamId":200,"pickTurn":4},{"championId":102,"teamId":100,"pickTurn":5},{"championId":39,"teamId":200,"pickTurn":6}],"gameStartTime":1440434545824,"gameLength":209},{"gameId":1230543473,"mapId":11,"gameMode":"CLASSIC","gameType":"MATCHED_GAME","gameQueueConfigId":4,"participants":[{"teamId":100,"spell1Id":4,"spell2Id":14,"championId":1,"skinIndex":8,"profileIconId":682,"summonerName":"lukezor","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":12,"championId":54,"skinIndex":0,"profileIconId":545,"summonerName":"MrG0dlik3","bot":false},{"teamId":100,"spell1Id":11,"spell2Id":4,"championId":102,"skinIndex":0,"profileIconId":22,"summonerName":"PÃ¶silÃ¶","bot":false},{"teamId":100,"spell1Id":7,"spell2Id":4,"championId":21,"skinIndex":0,"profileIconId":588,"summonerName":"Deepse","bot":false},{"teamId":100,"spell1Id":3,"spell2Id":4,"championId":16,"skinIndex":0,"profileIconId":7,"summonerName":"LeoTKO","bot":false},{"teamId":200,"spell1Id":7,"spell2Id":4,"championId":429,"skinIndex":1,"profileIconId":767,"summonerName":"pancakes yumyum","bot":false},{"teamId":200,"spell1Id":12,"spell2Id":4,"championId":86,"skinIndex":0,"profileIconId":903,"summonerName":"LuG0","bot":false},{"teamId":200,"spell1Id":14,"spell2Id":4,"championId":412,"skinIndex":3,"profileIconId":767,"summonerName":"Pyl","bot":false},{"teamId":200,"spell1Id":14,"spell2Id":4,"championId":117,"skinIndex":3,"profileIconId":508,"summonerName":"ReaI SIim Shady","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":11,"championId":76,"skinIndex":6,"profileIconId":577,"summonerName":"Thanee","bot":false}],"observers":{"encryptionKey":"H1aGTzXCCWGA10Lui8eankXkF7c1MCJw"},"platformId":"EUN1","gameTypeConfigId":2,"bannedChampions":[{"championId":72,"teamId":100,"pickTurn":1},{"championId":60,"teamId":200,"pickTurn":2},{"championId":107,"teamId":100,"pickTurn":3},{"championId":82,"teamId":200,"pickTurn":4},{"championId":92,"teamId":100,"pickTurn":5},{"championId":41,"teamId":200,"pickTurn":6}],"gameStartTime":1440434724262,"gameLength":30},{"gameId":1230543686,"mapId":11,"gameMode":"CLASSIC","gameType":"MATCHED_GAME","gameQueueConfigId":4,"participants":[{"teamId":100,"spell1Id":3,"spell2Id":4,"championId":111,"skinIndex":0,"profileIconId":767,"summonerName":"flastyBiceps","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":12,"championId":41,"skinIndex":0,"profileIconId":605,"summonerName":"Kryndor","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":7,"championId":104,"skinIndex":3,"profileIconId":682,"summonerName":"A F Y","bot":false},{"teamId":100,"spell1Id":4,"spell2Id":11,"championId":56,"skinIndex":0,"profileIconId":7,"summonerName":"Prokyon","bot":false},{"teamId":100,"spell1Id":14,"spell2Id":4,"championId":34,"skinIndex":5,"profileIconId":745,"summonerName":"Kaanv","bot":false},{"teamId":200,"spell1Id":7,"spell2Id":4,"championId":236,"skinIndex":0,"profileIconId":621,"summonerName":"Impure Foolz","bot":false},{"teamId":200,"spell1Id":14,"spell2Id":4,"championId":157,"skinIndex":2,"profileIconId":898,"summonerName":"Behold","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":12,"championId":54,"skinIndex":5,"profileIconId":580,"summonerName":"Hemi","bot":false},{"teamId":200,"spell1Id":4,"spell2Id":14,"championId":25,"skinIndex":4,"profileIconId":548,"summonerName":"HellCat Boldicat","bot":false},{"teamId":200,"spell1Id":11,"spell2Id":4,"championId":60,"skinIndex":1,"profileIconId":761,"summonerName":"Mitocanu","bot":false}],"observers":{"encryptionKey":"1xamnxRRPCoaIwNStBbBNVWXKVRijem3"},"platformId":"EUN1","gameTypeConfigId":2,"bannedChampions":[{"championId":72,"teamId":100,"pickTurn":1},{"championId":39,"teamId":200,"pickTurn":2},{"championId":82,"teamId":100,"pickTurn":3},{"championId":28,"teamId":200,"pickTurn":4},{"championId":105,"teamId":100,"pickTurn":5},{"championId":107,"teamId":200,"pickTurn":6}],"gameStartTime":1440434596657,"gameLength":158}],"clientRefreshInterval":300}`)
}

// Contains the current version for this Region.
func version() string {
	return "1.82.89"
}

// URL: .../consumer/getGameMetaData/{platformId}/{gameID}/1/token
// Returns information about the given game. This contains the games type and map, summoners involved, champions picked & banned, start time of the game and the encryption key required to read the replay data.
func getGameMetaData(platformId, gameId string) (*MetaData, error) {
	filepath := makePath(platformId, gameId) + "/metas.json"
	byteArr, err := BufferedRead(filepath)

	if err != nil {
		return nil, err
	}

	var meta MetaData
	if err = json.Unmarshal(byteArr, &meta); err != nil {
		return nil, err
	}

	return &meta, nil
}

// URL: .../consumer/getLastChunkInfo/{platformId}/{gameID}/{param}/token
// Return some information about the last available chunk:
func getLastChunkInfo(platformId, gameId string, param string) (*ChunkInfo, error) {
	meta, err := getGameMetaData(platformId, gameId)
	if err != nil {
		return nil, err
	}

	result := ChunkInfo{
		ChunkId:            meta.StartGameChunkId,
		AvailableSince:     30000,
		NextAvailableChunk: 30000,
		KeyFrameId:         1,
		NextChunkId:        meta.StartGameChunkId,
		EndStartupChunkId:  meta.EndStartupChunkId,
		StartGameChunkId:   meta.StartGameChunkId,
		EndGameChunkId:     0,
		Duration:           30000,
	}

	//0 Param requires end data for client to know to stream the rest
	if tmp, _ := strconv.Atoi(param); tmp == 0 {
		result = ChunkInfo{
			ChunkId:            meta.LastChunkId,
			AvailableSince:     30000,
			NextAvailableChunk: 30000,
			KeyFrameId:         meta.LastKeyFrameId,
			NextChunkId:        meta.LastChunkId - 1,
			EndStartupChunkId:  meta.EndStartupChunkId,
			StartGameChunkId:   meta.StartGameChunkId,
			EndGameChunkId:     meta.LastChunkId,
			Duration:           30000,
		}
	}

	return &result, nil
}

// URL: .../consumer/getGameDataChunk/{platformId}/{gameID}/{chunkId}/token
// Retrieves a chunk of data for the given game.
func getGameDataChunk(platformId, gameId, chunkId string) ([]byte, error) {
	filepath := makePath(platformId, gameId) + fmt.Sprintf("/chunks/%s", chunkId)
	return BufferedRead(filepath)
}

// URL: .../consumer/getKeyFrame/{platformId}/{gameID}/{keyFrameId}/token
// Retrieves a key frame for the given game.
func getKeyFrame(platformId, gameId, keyFrameId string) ([]byte, error) {
	filepath := makePath(platformId, gameId) + fmt.Sprintf("/keyframes/%s", keyFrameId)
	return BufferedRead(filepath)
}

// INCOMPLETE
// URL: .../consumer/getLastChunkInfo/{platformId}/{gameID}/null (!)
// Contains data used for the statistics screen after a game.
func endOfGameStats(platformId, gameId string) []byte {
	return nil
}
