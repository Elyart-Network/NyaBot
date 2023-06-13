package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/cqcode"
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/types"
	lua "github.com/yuin/gopher-lua"
	luar "layeh.com/gopher-luar"
)

type CqCodeFunc struct{}

func (c CqCodeFunc) Find(msg string) []string {
	return cqcode.Find(msg)
}

func (c CqCodeFunc) Decode(str string) any {
	return cqcode.Decode(str)
}

func (c CqCodeFunc) DecodeOne(code string, key string) string {
	return cqcode.Decode(code)[key]
}

func (c CqCodeFunc) Face(id string) string {
	data := types.FaceData{ID: id}
	return cqcode.Face(data)
}

func (c CqCodeFunc) Record(File string, Magic string, Cache string, Proxy string, Timeout string) string {
	data := types.RecordData{
		File:    File,
		Magic:   Magic,
		Cache:   Cache,
		Proxy:   Proxy,
		Timeout: Timeout,
	}
	return cqcode.Record(data)
}

func (c CqCodeFunc) Video(File string) string {
	data := types.VideoData{
		File: File,
	}
	return cqcode.Video(data)
}

func (c CqCodeFunc) VideoFull(File string, Cover string, Thread string) string {
	data := types.VideoData{
		File:   File,
		Cover:  Cover,
		Thread: Thread,
	}
	return cqcode.VideoFull(data)
}

func (c CqCodeFunc) At(QQ string, Name string) string {
	data := types.AtData{
		QQ:   QQ,
		Name: Name,
	}
	return cqcode.At(data)
}

func (c CqCodeFunc) Rps() string {
	return cqcode.Rps()
}

func (c CqCodeFunc) Dice() string {
	return cqcode.Dice()
}

func (c CqCodeFunc) Shake() string {
	return cqcode.Shake()
}

func (c CqCodeFunc) Anonymous() string {
	return cqcode.Anonymous()
}

func (c CqCodeFunc) Share(URL string, Title string) string {
	data := types.ShareData{
		URL:   URL,
		Title: Title,
	}
	return cqcode.Share(data)
}

func (c CqCodeFunc) ShareFull(URL string, Title string, Content string, Image string) string {
	data := types.ShareData{
		URL:     URL,
		Title:   Title,
		Content: Content,
		Image:   Image,
	}
	return cqcode.ShareFull(data)
}

func (c CqCodeFunc) Contact(Type string, ID string) string {
	data := types.ContactData{
		Type: Type,
		ID:   ID,
	}
	return cqcode.Contact(data)
}

func (c CqCodeFunc) Location(Lat string, Lon string) string {
	data := types.LocationData{
		Lat: Lat,
		Lon: Lon,
	}
	return cqcode.Location(data)
}

func (c CqCodeFunc) LocationFull(Lat string, Lon string, Title string, Content string) string {
	data := types.LocationData{
		Lat:     Lat,
		Lon:     Lon,
		Title:   Title,
		Content: Content,
	}
	return cqcode.LocationFull(data)
}

func (c CqCodeFunc) Music(Type string, ID string) string {
	data := types.MusicData{Type: Type, ID: ID}
	return cqcode.Music(data)
}

func (c CqCodeFunc) MusicFull(Type string, ID string, Audio string, Title string, Content string, Image string) string {
	data := types.MusicData{
		Type:    Type,
		ID:      ID,
		Audio:   Audio,
		Title:   Title,
		Content: Content,
		Image:   Image,
	}
	return cqcode.MusicFull(data)
}

func (c CqCodeFunc) Image(File string, URL string) string {
	data := types.ImageData{File: File, Url: URL}
	return cqcode.Image(data)
}

func (c CqCodeFunc) ImageFull(File string, Type string, SubType string, URL string, Cache string, ID string, Thread string) string {
	data := types.ImageData{
		File:    File,
		Type:    Type,
		SubType: SubType,
		Url:     URL,
		Cache:   Cache,
		ID:      ID,
		Thread:  Thread,
	}
	return cqcode.ImageFull(data)
}

func (c CqCodeFunc) Reply(ID string) string {
	data := types.ReplyData{ID: ID}
	return cqcode.Reply(data)
}

func (c CqCodeFunc) ReplyFull(ID string, Text string, QQ string, Time string, Seq string) string {
	data := types.ReplyData{
		ID:   ID,
		Text: Text,
		QQ:   QQ,
		Time: Time,
		Seq:  Seq,
	}
	return cqcode.ReplyFull(data)
}

func (c CqCodeFunc) RedBag(Title string) string {
	data := types.RedBagData{Title: Title}
	return cqcode.RedBag(data)
}

func (c CqCodeFunc) Poke(QQ string) string {
	data := types.PokeData{QQ: QQ}
	return cqcode.Poke(data)
}

func (c CqCodeFunc) Gift(QQ string, ID string) string {
	data := types.GiftData{
		QQ: QQ,
		ID: ID,
	}
	return cqcode.Gift(data)
}

func (c CqCodeFunc) Forward(ID string) string {
	data := types.ForwardData{ID: ID}
	return cqcode.Forward(data)
}

func (c CqCodeFunc) Xml(Data string, ResID string) string {
	data := types.XmlData{
		Data:  Data,
		ResID: ResID,
	}
	return cqcode.Xml(data)
}

func (c CqCodeFunc) Json(Data string) string {
	data := types.JsonData{Data: Data}
	return cqcode.Json(data)
}

func (c CqCodeFunc) JsonRich(Data string, ResID string) string {
	data := types.JsonData{
		Data:  Data,
		ResID: ResID,
	}
	return cqcode.JsonRich(data)
}

func (c CqCodeFunc) CardImage(File string) string {
	data := types.CardImageData{File: File}
	return cqcode.CardImage(data)
}

func (c CqCodeFunc) CardImageFull(File string, MinWidth string, MaxWidth string, MinHeight string, MaxHeight string, Source string, Icon string) string {
	data := types.CardImageData{
		File:      File,
		MinWidth:  MinWidth,
		MaxWidth:  MaxWidth,
		MinHeight: MinHeight,
		MaxHeight: MaxHeight,
		Source:    Source,
		Icon:      Icon,
	}
	return cqcode.CardImageFull(data)
}

func (c CqCodeFunc) Tts(Text string) string {
	data := types.TtsData{Text: Text}
	return cqcode.Tts(data)
}

func (c CqCodeFunc) CqCode(L *lua.LState) {
	var CqCodeFunc = &CqCodeFunc{}
	L.SetGlobal("CqCode", luar.New(L, CqCodeFunc))
}
