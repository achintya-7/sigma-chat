import 'package:freezed_annotation/freezed_annotation.dart';

part 'chats.freezed.dart';
part 'chats.g.dart';

@freezed
class Chat with _$Chat {
  factory Chat({
    required String msg,
  }) = _Chat;

  factory Chat.fromJson(Map<String, dynamic> json) => _$ChatFromJson(json);
}
