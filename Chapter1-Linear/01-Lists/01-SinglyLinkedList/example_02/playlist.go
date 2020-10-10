package main

import "fmt"

// 歌曲的结构体
// 包含歌曲名 歌手 下一首
type song struct {
	name   string
	artist string
	next   *song
}

// 歌单播放列表
// 列表名 列表里的第一首歌曲 正在播放的歌曲
type playlist struct {
	name       string
	head       *song
	nowPlaying *song
}

// 根据名称创建一个播放列表
func createPlaylist(name string) *playlist {
	return &playlist{
		name: name,
	}
}

// 添加歌曲 根据歌名 歌手
func (p *playlist) addSong(name, artist string) error {
	// 创建歌曲
	s := &song{
		name:   name,
		artist: artist,
	}
	// 播放列表的第一个节点等于nil代表这个播放列表单链表是空的
	if p.head == nil {
		// 直接添加歌曲到播放列表的第一个节点
		p.head = s
	} else {
		// 遍历播放列表找到最后一个歌曲
		currentNode := p.head
		// 当这个currentNode 也就是当前歌曲的指向下一个歌曲的指针是空的
		// 这个currentNode就是播放列表的最后一个节点
		for currentNode.next != nil {
			// 将下一个节点赋值到当前结点 进行下一次for循环
			currentNode = currentNode.next
		}
		// 跳出for循环的这个currentNode就是当前播放列表的最后一个节点
		// 也就是最后一首歌曲 将这个新歌曲赋值到它的next指向
		currentNode.next = s
	}
	return nil
}

// 显示所有的歌曲
func (p *playlist) showAllSongs() error {
	currentNode := p.head
	// 如果当前的首节点是空的这个播放列表是空的
	if currentNode == nil {
		fmt.Println("Playlist is empty.")
		return nil
	}
	fmt.Printf("%+v\n", *currentNode)
	// 遍历显示整个播放列表单链表
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}

	return nil
}

// 开始播放
func (p *playlist) startPlaying() *song {
	// 从头开始播放
	p.nowPlaying = p.head
	return p.nowPlaying
}

// 下一首
func (p *playlist) nextSong() *song {
	// 将当前播放的歌曲的下一个节点赋值给nowPlaying
	p.nowPlaying = p.nowPlaying.next
	return p.nowPlaying
}

func main() {
	playlistName := "myplaylist"
	myPlaylist := createPlaylist(playlistName)
	fmt.Println("Created playlist")
	fmt.Println()

	fmt.Print("Adding songs to the playlist...\n\n")
	myPlaylist.addSong("Ophelia", "The Lumineers")
	myPlaylist.addSong("Shape of you", "Ed Sheeran")
	myPlaylist.addSong("Stubborn Love", "The Lumineers")
	myPlaylist.addSong("Feels", "Calvin Harris")
	fmt.Println("Showing all songs in playlist...")
	myPlaylist.showAllSongs()
	fmt.Println()

	myPlaylist.startPlaying()
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()

	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)
	fmt.Println()
	myPlaylist.nextSong()
	fmt.Println("Changing next song...")
	fmt.Printf("Now playing: %s by %s\n", myPlaylist.nowPlaying.name, myPlaylist.nowPlaying.artist)

}
