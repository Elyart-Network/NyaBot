package cqcode

func Face(data FaceData) string {
	return "[CQ:face,id=" + data.ID + "]"
}

func Record(data RecordData) string {
	return "[CQ:record,file=" + data.File + ",magic=" + data.Magic + ",cache=" + data.Cache + ",proxy=" + data.Proxy + ",timeout=" + data.Timeout + "]"
}

func Video(data VideoData) string {
	return "[CQ:video,file=" + data.File + "]"
}

func VideoFull(data VideoData) string {
	return "[CQ:video,file=" + data.File + ",cover=" + data.Cover + ",c=" + data.Thread + "]"
}

func At(data AtData) string {
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

func Share(data ShareData) string {
	return "[CQ:share,url=" + data.URL + ",title=" + data.Title + "]"
}

func ShareFull(data ShareData) string {
	return "[CQ:share,url=" + data.URL + ",title=" + data.Title + ",content=" + data.Content + ",image=" + data.Image + "]"
}

func Contact(data ContactData) string {
	return "[CQ:contact,type=" + data.Type + ",id=" + data.ID + "]"
}

func Location(data LocationData) string {
	return "[CQ:location,lat=" + data.Lat + ",lon=" + data.Lon + "]"
}

func LocationFull(data LocationData) string {
	return "[CQ:location,lat=" + data.Lat + ",lon=" + data.Lon + ",title=" + data.Title + ",content=" + data.Content + "]"
}

func Music(data MusicData) string {
	return "[CQ:music,type=" + data.Type + ",id=" + data.ID + "]"
}

func MusicFull(data MusicData) string {
	return "[CQ:music,type=" + data.Type + ",id=" + data.ID + ",audio=" + data.Audio + ",title=" + data.Title + ",content=" + data.Content + ",image=" + data.Image + "]"
}

func Image(data ImageData) string {
	return "[CQ:image,file=" + data.File + "url" + data.Url + "]"
}

func ImageFull(data ImageData) string {
	return "[CQ:image,file=" + data.File + ",type=" + data.Type + ",subType=" + data.SubType + ",url=" + data.Url + ",cache=" + data.Cache + ",id=" + data.ID + ",c=" + data.Thread + "]"
}

func Reply(data ReplyData) string {
	return "[CQ:reply,id=" + data.ID + "]"
}

func ReplyFull(data ReplyData) string {
	return "[CQ:reply,id=" + data.ID + ",text=" + data.Text + ",qq=" + data.QQ + ",time=" + data.Time + ",seq=" + data.Seq + "]"
}

func RedBag(data RedBagData) string {
	return "[CQ:redbag,title=" + data.Title + "]"
}

func Poke(data PokeData) string {
	return "[CQ:poke,qq=" + data.QQ + "]"
}

func Gift(data GiftData) string {
	return "[CQ:gift,qq=" + data.QQ + ",id=" + data.ID + "]"
}

func Forward(data ForwardData) string {
	return "[CQ:forward,id=" + data.ID + "]"
}

func Xml(data XmlData) string {
	return "[CQ:xml,data=" + data.Data + ",resid=" + data.ResID + "]"
}

func Json(data JsonData) string {
	return "[CQ:json,data=" + data.Data + "]"
}

func JsonRich(data JsonData) string {
	return "[CQ:json,data=" + data.Data + ",resid=" + data.ResID + "]"
}

func CardImage(data CardImageData) string {
	return "[CQ:cardimage,file=" + data.File + "]"
}

func CardImageFull(data CardImageData) string {
	return "[CQ:cardimage,file=" + data.File + ",minwidth=" + data.MinWidth + ",minheight=" + data.MinHeight + ",maxwidth=" + data.MaxWidth + ",maxheight=" + data.MaxHeight + ",source=" + data.Source + ",icon=" + data.Icon + "]"
}

func Tts(data TtsData) string {
	return "[CQ:tts,text=" + data.Text + "]"
}
