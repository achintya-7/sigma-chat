import 'package:flutter/material.dart';

class MyInputField extends StatelessWidget {
  const MyInputField(
      {super.key,
      required this.labelText,
      required this.obscureText,
      required this.controller});

  final String labelText;
  final bool obscureText;
  final TextEditingController controller;

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 25),
      child: TextField(
        controller: controller,
        obscureText: obscureText,
        decoration: InputDecoration(
          enabledBorder: const OutlineInputBorder(
            borderSide: BorderSide(color: Colors.white),
            borderRadius: BorderRadius.all(Radius.circular(10.0)),
          ),
          focusedBorder: const OutlineInputBorder(
            borderSide: BorderSide(color: Colors.grey),
          ),
          labelText: labelText,
          fillColor: Colors.grey.shade200,
          filled: true,
        ),
      ),
    );
  }
}
