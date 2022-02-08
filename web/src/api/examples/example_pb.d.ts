// package: 
// file: example.proto

import * as jspb from "google-protobuf";

export class ExampleRequest extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExampleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ExampleRequest): ExampleRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExampleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExampleRequest;
  static deserializeBinaryFromReader(message: ExampleRequest, reader: jspb.BinaryReader): ExampleRequest;
}

export namespace ExampleRequest {
  export type AsObject = {
    username: string,
  }
}

export class ExampleReply extends jspb.Message {
  getUsername(): string;
  setUsername(value: string): void;

  getCount(): number;
  setCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ExampleReply.AsObject;
  static toObject(includeInstance: boolean, msg: ExampleReply): ExampleReply.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ExampleReply, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ExampleReply;
  static deserializeBinaryFromReader(message: ExampleReply, reader: jspb.BinaryReader): ExampleReply;
}

export namespace ExampleReply {
  export type AsObject = {
    username: string,
    count: number,
  }
}

