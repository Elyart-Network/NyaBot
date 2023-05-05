package types

type Message struct {
	Sender    Sender      `bson:"sender"`
	Content   string      `bson:"content"`
	TimeStamp int64       `bson:"timestamp"`
	RawMsg    string      `bson:"raw_msg"`
	Raw       interface{} `bson:"raw"`
}

type Sender struct {
	UserIdentity string `bson:"user_identity"`
	UserName     string `bson:"user_name,omitempty"`
	UserNick     string `bson:"user_nick,omitempty"`
	UserAddition string `bson:"user_addition,omitempty"`
}
