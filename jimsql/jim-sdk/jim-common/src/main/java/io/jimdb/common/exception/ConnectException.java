/*
 * Copyright 2019 The JIMDB Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */
package io.jimdb.common.exception;

/**
 * Connection Exception
 *
 */
public final class ConnectException extends BaseException {
  private ConnectException(ErrorCode code, String message, Throwable cause) {
    super(ErrorModule.RPC, code, message, cause);
  }

  public static ConnectException get(ErrorCode code, String... params) {
    return new ConnectException(code, message(code, params), null);
  }

  public static ConnectException get(ErrorCode code, Throwable cause, String... params) {
    return new ConnectException(code, message(code, params), cause);
  }
}
