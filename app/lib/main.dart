import 'package:app/features/auth/views/sign_in_page.dart';
import 'package:app/features/auth/views/sign_up_page.dart';
import 'package:app/features/chat/bindings/chat_bindings.dart';
import 'package:app/features/chat/views/chat_page.dart';
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
      initialRoute: '/signIn',
      getPages: [
        GetPage(name: '/signIn', page: () => AuthPage()),
        GetPage(name: '/signUp', page: () => SignUpPage()),
        GetPage(name: '/chat', page: () => ChatPage(), binding: ChatBindings()),
      ],
    );
  }
}
