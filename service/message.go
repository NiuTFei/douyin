package service

import (
	"douyin/repository"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
)

var chatConnMap = sync.Map{}

func RunMessageServer() {
	listen, err := net.Listen("tcp", "10.180.22.64:9090")

	if err != nil {
		fmt.Printf("Run message sever failed: %v\n", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("Accept conn failed: %v\n", err)
			continue
		}

		go process(conn)
	}
}

func process(conn net.Conn) {
	defer conn.Close()

	var buf [256]byte
	for {
		n, err := conn.Read(buf[:])
		if n == 0 {
			if err == io.EOF {
				break
			}
			fmt.Printf("Read message failed: %v\n", err)
			continue
		}

		var event = repository.MessageSendEvent{}
		_ = json.Unmarshal(buf[:n], &event)
		fmt.Printf("Receive Message：%+v\n", event)

		fromChatKey := fmt.Sprintf("%d_%d", event.UserId, event.ToUserId)
		if len(event.MsgContent) == 0 {
			chatConnMap.Store(fromChatKey, conn)
			continue
		}

		toChatKey := fmt.Sprintf("%d_%d", event.ToUserId, event.UserId)
		writeConn, exist := chatConnMap.Load(toChatKey)
		if !exist {
			fmt.Printf("User %d offline\n", event.ToUserId)
			continue
		}

		pushEvent := repository.MessagePushEvent{
			FromUserId: event.UserId,
			MsgContent: event.MsgContent,
		}
		pushData, _ := json.Marshal(pushEvent)
		_, err = writeConn.(net.Conn).Write(pushData)
		if err != nil {
			fmt.Printf("Push message failed: %v\n", err)
		}
	}
}

func MessageClient(token string, toUserId string, content string) {
	userIdA := CheckToken(token)
	if userIdA == 0 || userIdA == -1 {
		return
	}
	//userIdA, _ := strconv.ParseInt(userId, 10, 64)
	userIdB, _ := strconv.ParseInt(toUserId, 10, 64)
	message := repository.MessageDao{FromUserId: userIdA, ToUserId: userIdB, Content: content}

	conn, err := net.Dial("tcp", "10.180.22.64:9090")

	if err != nil {
		fmt.Printf("Message client  failed: %v\n", err)
		return
	}
	defer conn.Close()

	toChatKey := fmt.Sprintf("%d_%d", userIdB, userIdA)
	chatConnMap.Store(toChatKey, conn)

	sendEvent := repository.MessageSendEvent{
		UserId:     userIdA,
		ToUserId:   userIdB,
		MsgContent: content,
	}
	writeBuf, _ := json.Marshal(sendEvent)
	_, err = conn.Write(writeBuf)
	if err != nil {
		fmt.Printf("Send message failed: %v\n", err)
		return
	}
	db.Create(&message)

}

func MessageChat(token string, toUserIdStr string) repository.MessageListResponse {
	fromUserId := CheckToken(token)
	if fromUserId == 0 || fromUserId == -1 {
		return repository.MessageListResponse{
			Response: repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"},
		}
	}
	toUserId, _ := strconv.ParseInt(toUserIdStr, 10, 64)

	var messageDaoList []repository.MessageDao
	db.Where("from_user_id = ? and to_user_id = ?", fromUserId, toUserId).Find(&messageDaoList)

	return repository.MessageListResponse{
		Response:    repository.Response{StatusCode: 1, StatusMsg: "用户未登陆或不存在"},
		MessageList: MessageDaoListToMessageList(messageDaoList),
	}
}
