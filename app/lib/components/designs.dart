import 'package:flutter/material.dart';

class Designs {
  InputDecoration inputDecoration(String? labelText) {
    return InputDecoration(
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
    );
  }
}