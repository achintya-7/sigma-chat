import 'package:app/features/chat/bindings/chat_bindings.dart';
import 'package:app/features/chat/views/chat_page.dart';
import 'package:app/features/home/views/home_page.dart';
import 'package:flutter/material.dart';
import 'package:get/route_manager.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return GetMaterialApp(
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      initialRoute: '/chat',
      getPages: [
        GetPage(name: '/home', page: () => HomePage()),
        GetPage(name: '/chat', page: () => ChatPage(), binding: ChatBindings()),
      ],
    );
  }
}
