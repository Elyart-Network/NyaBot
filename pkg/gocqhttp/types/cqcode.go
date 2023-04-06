package types

type FaceData struct {
	ID string
}

type RecordData struct {
	File    string
	Magic   string
	Cache   string
	Proxy   string
	Timeout string
}

type VideoData struct {
	File   string
	Cover  string
	Thread string
}

type AtData struct {
	QQ   string
	Name string
}

type ShareData struct {
	URL     string
	Title   string
	Content string
	Image   string
}

type ContactData struct {
	Type string
	ID   string
}

type LocationData struct {
	Lat     string
	Lon     string
	Title   string
	Content string
}

type MusicData struct {
	Type    string
	ID      string
	Audio   string
	Title   string
	Content string
	Image   string
}

type ImageData struct {
	File    string
	Type    string
	SubType string
	Url     string
	Cache   string
	ID      string
	Thread  string
}

type ReplyData struct {
	ID   string
	Text string
	QQ   string
	Time string
	Seq  string
}

type RedBagData struct {
	Title string
}

type PokeData struct {
	QQ string
}

type GiftData struct {
	QQ string
	ID string
}

type ForwardData struct {
	ID string
}

type XmlData struct {
	Data  string
	ResID string
}

type JsonData struct {
	Data  string
	ResID string
}

type CardImageData struct {
	File      string
	MinWidth  string
	MaxWidth  string
	MinHeight string
	MaxHeight string
	Source    string
	Icon      string
}

type TtsData struct {
	Text string
}
