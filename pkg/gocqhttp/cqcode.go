package gocqhttp

import (
	"github.com/Elyart-Network/NyaBot/pkg/gocqhttp/models"
)

func Face(data models.FaceData) string {
	return "[CQ:face,id=" + data.ID + "]"
}

func Record(data models.RecordData) string {
	return "[CQ:record,file=" + data.File + ",magic=" + data.Magic + ",cache=" + data.Cache + ",proxy=" + data.Proxy + ",timeout=" + data.Timeout + "]"
}

func Video(data models.VideoData) string {
	return "[CQ:video,file=" + data.File + "]"
}

func VideoFull(data models.VideoData) string {
	return "[CQ:video,file=" + data.File + ",cover=" + data.Cover + ",c=" + data.Thread + "]"
}

func At(data models.AtData) string {
	return "[CQ:at,qq=" + data.QQ + ",name=" + data.Name + "]"
}

func Rps() string {
	return "[CQ:rps]"
}

func Dice() string {
	return "[CQ:dice]"
}

func Shake() string {
	return "[CQ:shake]"
}

func Anonymous() string {
	return "[CQ:anonymous]"
}

func Share(data models.ShareData) string {
	return "[CQ:share,url=" + data.URL + ",title=" + data.Title + "]"
}

func ShareFull(data models.ShareData) string {
	return "[CQ:share,url=" + data.URL + ",title=" + data.Title + ",content=" + data.Content + ",image=" + data.Image + "]"
}

func Contact(data models.ContactData) string {
	return "[CQ:contact,type=" + data.Type + ",id=" + data.ID + "]"
}

func Location(data models.LocationData) string {
	return "[CQ:location,lat=" + data.Lat + ",lon=" + data.Lon + "]"
}

func LocationFull(data models.LocationData) string {
	return "[CQ:location,lat=" + data.Lat + ",lon=" + data.Lon + ",title=" + data.Title + ",content=" + data.Content + "]"
}

func Music(data models.MusicData) string {
	return "[CQ:music,type=" + data.Type + ",id=" + data.ID + "]"
}

func MusicFull(data models.MusicData) string {
	return "[CQ:music,type=" + data.Type + ",id=" + data.ID + ",audio=" + data.Audio + ",title=" + data.Title + ",content=" + data.Content + ",image=" + data.Image + "]"
}

func Image(data models.ImageData) string {
	return "[CQ:image,file=" + data.File + ",url=" + data.Url + "]"
}

func ImageFull(data models.ImageData) string {
	return "[CQ:image,file=" + data.File + ",type=" + data.Type + ",subType=" + data.SubType + ",url=" + data.Url + ",cache=" + data.Cache + ",id=" + data.ID + ",c=" + data.Thread + "]"
}

func Reply(data models.ReplyData) string {
	return "[CQ:reply,id=" + data.ID + "]"
}

func ReplyFull(data models.ReplyData) string {
	return "[CQ:reply,id=" + data.ID + ",text=" + data.Text + ",qq=" + data.QQ + ",time=" + data.Time + ",seq=" + data.Seq + "]"
}

func RedBag(data models.RedBagData) string {
	return "[CQ:redbag,title=" + data.Title + "]"
}

func Poke(data models.PokeData) string {
	return "[CQ:poke,qq=" + data.QQ + "]"
}

func Gift(data models.GiftData) string {
	return "[CQ:gift,qq=" + data.QQ + ",id=" + data.ID + "]"
}

func Forward(data models.ForwardData) string {
	return "[CQ:forward,id=" + data.ID + "]"
}

func Xml(data models.XmlData) string {
	return "[CQ:xml,data=" + data.Data + ",resid=" + data.ResID + "]"
}

func Json(data models.JsonData) string {
	return "[CQ:json,data=" + data.Data + "]"
}

func JsonRich(data models.JsonData) string {
	return "[CQ:json,data=" + data.Data + ",resid=" + data.ResID + "]"
}

func CardImage(data models.CardImageData) string {
	return "[CQ:cardimage,file=" + data.File + "]"
}

func CardImageFull(data models.CardImageData) string {
	return "[CQ:cardimage,file=" + data.File + ",minwidth=" + data.MinWidth + ",minheight=" + data.MinHeight + ",maxwidth=" + data.MaxWidth + ",maxheight=" + data.MaxHeight + ",source=" + data.Source + ",icon=" + data.Icon + "]"
}

func Tts(data models.TtsData) string {
	return "[CQ:tts,text=" + data.Text + "]"
}
