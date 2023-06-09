import 'package:freezed_annotation/freezed_annotation.dart';

part 'send.freezed.dart';
part 'send.g.dart';

@freezed
class Send with _$Send {
  factory Send({
    required String room,
    required String msg,
  }) = _Send;

  factory Send.fromJson(Map<String, dynamic> json) => _$SendFromJson(json);
}