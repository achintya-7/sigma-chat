import 'package:flutter/material.dart';

class Designs {
  InputDecoration inputDecoration(String? labelText) {
    return InputDecoration(
      enabledBorder: const OutlineInputBorder(
        borderSide: BorderSide(color: Colors.white),
        borderRadius: BorderRadius.all(Radius.circular(12.0)),
      ),
      focusedBorder: const OutlineInputBorder(
        borderSide: BorderSide(color: Colors.grey),
        borderRadius: BorderRadius.all(Radius.circular(12.0)),
      ),
      labelText: labelText,
      labelStyle: const TextStyle(color: Colors.grey),
      fillColor: Colors.grey.shade200,
      filled: true,
    );
  }
}
