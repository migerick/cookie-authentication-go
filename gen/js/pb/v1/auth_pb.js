// @generated by protoc-gen-es v1.6.0 with parameter "target=js+dts"
// @generated from file pb/v1/auth.proto (package pb.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message pb.v1.LoginRequest
 */
export const LoginRequest = proto3.makeMessageType(
  "pb.v1.LoginRequest",
  () => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ],
);

/**
 * @generated from message pb.v1.LogoutRequest
 */
export const LogoutRequest = proto3.makeMessageType(
  "pb.v1.LogoutRequest",
  [],
);

