// ignore_for_file: avoid_print

import 'package:app/models/chats/chats.dart';
import 'package:app/models/send/send.dart';
import 'package:get/get.dart';
import 'package:socket_io_client/socket_io_client.dart' as io;

class ChatController extends GetxController {
  late io.Socket socket;

  RxList<Chat> messages = <Chat>[].obs;

  @override
  void onInit() {
    super.onInit();
    connectSocket();
  }

  void connectSocket() {
    try {
      socket = io.io('http://localhost:4001', <String, dynamic>{
        'transports': ['websocket'],
        'autoConnect': false,
      });

      print("Connecting to socket");
      socket.connect();

      socket.onConnect((_) {
        print('Connected to socket');
        socket.emit('join', "1234");
      });

      socket.on('message', (data) {
        print(data);
      });

    } on Exception catch (e) {
      print(e);
    }
  }

  void sendMessage(Send send) {
    socket.emit('send', send.toJson().toString());
  }


}

