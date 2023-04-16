// ignore_for_file: avoid_print

import 'dart:convert';

import 'package:app/models/send/send.dart';
import 'package:fluttertoast/fluttertoast.dart';
import 'package:get/get.dart';
import 'package:socket_io_client/socket_io_client.dart' as io;

class ChatController extends GetxController {
  late io.Socket socket;
  RxBool loading = false.obs;
  RxList<String> messages = <String>[].obs;

  @override
  void onInit() {
    super.onInit();
    testermessages();
    connectSocket();
  }

  void testermessages() {
    for (var i = 0; i < 20; i++) {
      messages.add("Test $i");
    }
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
        connectToRoom("1234");
      });

      socket.on('message', (data) {
        print(data);
        messages.add(data.toString());
      });
    } on Exception catch (e) {
      print(e);
      Fluttertoast.showToast(msg: "Error connecting to socket");
    }

    toogleLoading();
  }

  void connectToRoom(String room) {
    socket.emit('join', room);
  }

  void sendMessage(Send send) {
    try {
      socket.emit('send', jsonEncode(send));
    } on Exception catch (e) {
      print(e);
    }
  }

  void toogleLoading() {
    loading.value = !loading.value;
  }
}
