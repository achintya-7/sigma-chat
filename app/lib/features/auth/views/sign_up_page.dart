import 'package:app/features/auth/widgets/text_field.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';

class SignUpPage extends StatelessWidget {
  SignUpPage({super.key});

  final TextEditingController userNameController = TextEditingController();
  final TextEditingController emailController = TextEditingController();
  final TextEditingController passwordController = TextEditingController();
  final TextEditingController confirmPasswordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.grey[300],
      body: SafeArea(
          child: Center(
        child: Column(
          children: [
            const SizedBox(height: 50),
            Image.asset('lib/images/generatedtext.png'),
            const Icon(Icons.person_2, size: 100),
            const SizedBox(height: 40),
            Text(
              "Lets get started",
              style: TextStyle(
                fontSize: 16,
                color: Colors.grey[700],
              ),
            ),
            const SizedBox(height: 25),
            MyInputField(
              labelText: 'Username',
              obscureText: false,
              controller: userNameController,
            ),
            const SizedBox(height: 15),
            MyInputField(
              labelText: 'Email',
              obscureText: false,
              controller: emailController,
            ),
            const SizedBox(height: 15),
            MyInputField(
              labelText: 'Password',
              obscureText: true,
              controller: passwordController,
            ),
            const SizedBox(height: 15),
            MyInputField(
              labelText: 'Confirm Password',
              obscureText: true,
              controller: confirmPasswordController,
            ),
            const SizedBox(height: 25),
            ElevatedButton(
              onPressed: () {
                Get.offNamedUntil('/chat', (route) => false);
              },
              style: ElevatedButton.styleFrom(
                backgroundColor: Colors.black,
                elevation: 5,
                padding: const EdgeInsets.symmetric(
                  horizontal: 100,
                  vertical: 15,
                ),
              ),
              child: const Text('Sign Up'),
            ),
            const Spacer(),
            Row(
              mainAxisAlignment: MainAxisAlignment.center,
              children: [
                const Text('Already have an account?'),
                const SizedBox(width: 5),
                GestureDetector(
                  onTap: () {
                    Get.offNamedUntil('/signIn', (route) => false);
                  },
                  child: const Text(
                    'Sign In',
                    style: TextStyle(
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ),
              ],
            )
          ],
        ),
      )),
    );
  }
}
