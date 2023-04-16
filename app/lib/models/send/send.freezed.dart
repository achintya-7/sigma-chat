// coverage:ignore-file
// GENERATED CODE - DO NOT MODIFY BY HAND
// ignore_for_file: type=lint
// ignore_for_file: unused_element, deprecated_member_use, deprecated_member_use_from_same_package, use_function_type_syntax_for_parameters, unnecessary_const, avoid_init_to_null, invalid_override_different_default_values_named, prefer_expression_function_bodies, annotate_overrides, invalid_annotation_target, unnecessary_question_mark

part of 'send.dart';

// **************************************************************************
// FreezedGenerator
// **************************************************************************

T _$identity<T>(T value) => value;

final _privateConstructorUsedError = UnsupportedError(
    'It seems like you constructed your class using `MyClass._()`. This constructor is only meant to be used by freezed and you are not supposed to need it nor use it.\nPlease check the documentation here for more information: https://github.com/rrousselGit/freezed#custom-getters-and-methods');

Send _$SendFromJson(Map<String, dynamic> json) {
  return _Send.fromJson(json);
}

/// @nodoc
mixin _$Send {
  String get room => throw _privateConstructorUsedError;
  String get msg => throw _privateConstructorUsedError;

  Map<String, dynamic> toJson() => throw _privateConstructorUsedError;
  @JsonKey(ignore: true)
  $SendCopyWith<Send> get copyWith => throw _privateConstructorUsedError;
}

/// @nodoc
abstract class $SendCopyWith<$Res> {
  factory $SendCopyWith(Send value, $Res Function(Send) then) =
      _$SendCopyWithImpl<$Res, Send>;
  @useResult
  $Res call({String room, String msg});
}

/// @nodoc
class _$SendCopyWithImpl<$Res, $Val extends Send>
    implements $SendCopyWith<$Res> {
  _$SendCopyWithImpl(this._value, this._then);

  // ignore: unused_field
  final $Val _value;
  // ignore: unused_field
  final $Res Function($Val) _then;

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? room = null,
    Object? msg = null,
  }) {
    return _then(_value.copyWith(
      room: null == room
          ? _value.room
          : room // ignore: cast_nullable_to_non_nullable
              as String,
      msg: null == msg
          ? _value.msg
          : msg // ignore: cast_nullable_to_non_nullable
              as String,
    ) as $Val);
  }
}

/// @nodoc
abstract class _$$_SendCopyWith<$Res> implements $SendCopyWith<$Res> {
  factory _$$_SendCopyWith(_$_Send value, $Res Function(_$_Send) then) =
      __$$_SendCopyWithImpl<$Res>;
  @override
  @useResult
  $Res call({String room, String msg});
}

/// @nodoc
class __$$_SendCopyWithImpl<$Res> extends _$SendCopyWithImpl<$Res, _$_Send>
    implements _$$_SendCopyWith<$Res> {
  __$$_SendCopyWithImpl(_$_Send _value, $Res Function(_$_Send) _then)
      : super(_value, _then);

  @pragma('vm:prefer-inline')
  @override
  $Res call({
    Object? room = null,
    Object? msg = null,
  }) {
    return _then(_$_Send(
      room: null == room
          ? _value.room
          : room // ignore: cast_nullable_to_non_nullable
              as String,
      msg: null == msg
          ? _value.msg
          : msg // ignore: cast_nullable_to_non_nullable
              as String,
    ));
  }
}

/// @nodoc
@JsonSerializable()
class _$_Send implements _Send {
  _$_Send({required this.room, required this.msg});

  factory _$_Send.fromJson(Map<String, dynamic> json) => _$$_SendFromJson(json);

  @override
  final String room;
  @override
  final String msg;

  @override
  String toString() {
    return 'Send(room: $room, msg: $msg)';
  }

  @override
  bool operator ==(dynamic other) {
    return identical(this, other) ||
        (other.runtimeType == runtimeType &&
            other is _$_Send &&
            (identical(other.room, room) || other.room == room) &&
            (identical(other.msg, msg) || other.msg == msg));
  }

  @JsonKey(ignore: true)
  @override
  int get hashCode => Object.hash(runtimeType, room, msg);

  @JsonKey(ignore: true)
  @override
  @pragma('vm:prefer-inline')
  _$$_SendCopyWith<_$_Send> get copyWith =>
      __$$_SendCopyWithImpl<_$_Send>(this, _$identity);

  @override
  Map<String, dynamic> toJson() {
    return _$$_SendToJson(
      this,
    );
  }
}

abstract class _Send implements Send {
  factory _Send({required final String room, required final String msg}) =
      _$_Send;

  factory _Send.fromJson(Map<String, dynamic> json) = _$_Send.fromJson;

  @override
  String get room;
  @override
  String get msg;
  @override
  @JsonKey(ignore: true)
  _$$_SendCopyWith<_$_Send> get copyWith => throw _privateConstructorUsedError;
}
