import 'package:app/features/chat/controllers/chat_controller.dart';
import 'package:app/features/chat/widgets/chat_widgets.dart';
import 'package:app/models/send/send.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class ChatPage extends StatelessWidget {
  ChatPage({super.key});

  final ChatController chatController = Get.find<ChatController>();
  final TextEditingController messageController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        title: const Text("Chat Page"),
        centerTitle: true,
        backgroundColor: Colors.black,
      ),
      body: GestureDetector(
        onTap: () => FocusScope.of(context).unfocus(),
        child: Column(
          children: [
            Expanded(
              child: Container(
                decoration: BoxDecoration(
                  color: Colors.grey.shade200,
                  borderRadius: const BorderRadius.only(
                    topLeft: Radius.circular(30),
                    topRight: Radius.circular(30),
                  ),
                ),
                child: ClipRRect(
                  borderRadius: const BorderRadius.only(
                    topLeft: Radius.circular(30),
                    topRight: Radius.circular(30),
                  ),
                  child: Obx(
                    () => ListView.builder(
                      reverse: true,
                      padding: const EdgeInsets.only(top: 15),
                      itemCount: chatController.messages.length,
                      itemBuilder: (BuildContext context, int index) {
                        final message = chatController.messages[
                            chatController.messages.length - 1 - index];
                        final bool isUser = true;
                        return ChatWidgets.buildMessage(message, isUser);
                      },
                    ),
                  ),
                ),
              ),
            ),
            ChatWidgets.buildMessageComposer(messageController, chatController),
          ],
        ),
      ),
    );
  }
}
