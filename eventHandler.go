package main

import (
	"context"
	"errors"
	"fmt"
	server "github.com/rolancia/thing"
	"time"
)

type PostHandler struct {
}

type EventHandler struct {
}

func (h *EventHandler) OnConnected(conn *server.UserConn) (context.Context, server.PostAction) {
	//No Deadline set
	conn.Conn().SetDeadline(time.Time{})
	conn.Async = false
	//log.WithFields(logrus.Fields{
	//	"remote": conn.Conn().RemoteAddr(),
	//})

	return conn.Context(), server.PostActionNone
}

// called first packet received
func (h *EventHandler) OnJoin(conn *server.UserConn, firstMP *server.DefaultMessagePacket) (context.Context, server.PostAction) {
	// do login or ...
	//log.WithField("Protocol A, First payload received from [%s]", conn.Conn().RemoteAddr())

	h.OnMessage(conn, firstMP)
	return conn.Context(), server.PostActionNone
}

// called a packet received except first that
func (h *EventHandler) OnMessage(conn *server.UserConn, mp *server.DefaultMessagePacket) (context.Context, server.PostAction) {
	payload := mp.Payload

	if payload != nil {
		// do dispatch or ...
		//log.WithField("OnMessage:", string(mp.Payload))
		v := conn.Context().Value(payloadChannelContextKey)
		ch, ok := v.(chan Payload)
		if !ok {
			server.GetLandfill(conn.Context()) <- server.NewFatError(
				errors.New("Assert at OnMessage, channel has not been registered"), server.ErrorActionPrint, conn)
		} else {
			ch <- payload
		}
		return conn.Context(), server.PostActionNone
	} else {
		// this will call 'OnErrorPrint'
		server.GetLandfill(conn.Context()) <- server.NewFatError(errors.New("err!"), server.ErrorActionPrint, conn)
	}
	return conn.Context(), server.PostActionNone
}

// called before closing connection
func (h *EventHandler) OnBeforeClose(conn *server.UserConn) {
}

// called if the landfill consumes a fat error with ErrorActionPrint
func (h *EventHandler) OnErrorPrint(serverCtx context.Context, err *server.FatError) {
	fmt.Println(err.Error())
}

// called if the landfill consumes a fat error with ErrorActionSave
func (h *EventHandler) OnErrorSave(serverCtx context.Context, err *server.FatError) {
}

// called if deserialization failed
func (h *EventHandler) OnParsingFailed(conn *server.UserConn, data []byte) {
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// post action handler

func (h *PostHandler) OnPostAction(act server.PostAction, conn *server.UserConn) {
}
