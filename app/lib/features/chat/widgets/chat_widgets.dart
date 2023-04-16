import 'package:app/features/chat/controllers/chat_controller.dart';
import 'package:app/models/send/send.dart';
import 'package:flutter/material.dart';

class ChatWidgets{
  static buildMessage(String msg, bool isUser) {
    return Container(
      alignment: isUser ? Alignment.centerRight : Alignment.centerLeft,
      padding: EdgeInsets.only(
        top: 20,
        bottom: 20,
        left: isUser ? 0 : 15,
        right: isUser ? 15 : 0,
      ),
      margin: isUser
          ? const EdgeInsets.only(
              top: 8,
              bottom: 8,
              left: 80,
            )
          : const EdgeInsets.only(
              top: 8,
              bottom: 8,
              right: 80,
            ),
      decoration: BoxDecoration(
        color: isUser ? Colors.black : Colors.white,
        borderRadius: isUser
            ? const BorderRadius.only(
                topLeft: Radius.circular(15),
                bottomLeft: Radius.circular(15),
              )
            : const BorderRadius.only(
                topRight: Radius.circular(15),
                bottomRight: Radius.circular(15),
              ),
      ),
      child: Text(msg,
          style: TextStyle(color: isUser ? Colors.white : Colors.black)),
    );
  }

    static buildMessageComposer(TextEditingController messageController, ChatController chatController) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 8.0),
      height: 70,
      color: Colors.white,
      child: Row(
        children: [
          Expanded(
            child: TextField(
              controller: messageController,
              textCapitalization: TextCapitalization.sentences,
              onChanged: (value) {},
              decoration: const InputDecoration.collapsed(
                hintText: null,
              ),
            ),
          ),
          IconButton(
            icon: const Icon(Icons.send),
            iconSize: 30,
            color: Colors.black,
            onPressed: () {
              chatController.sendMessage(
                Send(
                  room: "1234",
                  msg: "Flutter is messaging",
                ),
              );
            },
          ),
        ],
      ),
    );
  }
}