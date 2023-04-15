// ignore_for_file: avoid_print

import 'package:get/get.dart';
import 'package:socket_io_client/socket_io_client.dart' as io;

class ChatController extends GetxController {
  late io.Socket socket;

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
        'force new connection': true,
      });

      print("Connecting to socket");
      socket.connect();

      socket.onConnect((_) {
        print('Connected to socket');
      });

    } on Exception catch (e) {
      print(e);
    }
  }
}
