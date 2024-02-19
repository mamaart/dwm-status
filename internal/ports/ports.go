package ports

type UniStreamer interface {
	Reader() <-chan []byte
}
