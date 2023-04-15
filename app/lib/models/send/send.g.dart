// GENERATED CODE - DO NOT MODIFY BY HAND

part of 'send.dart';

// **************************************************************************
// JsonSerializableGenerator
// **************************************************************************

_$_Send _$$_SendFromJson(Map<String, dynamic> json) => _$_Send(
      room: json['room'] as String,
      chat: Chat.fromJson(json['chat'] as Map<String, dynamic>),
    );

Map<String, dynamic> _$$_SendToJson(_$_Send instance) => <String, dynamic>{
      'room': instance.room,
      'chat': instance.chat,
    };
