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
package io.jimdb.sql.gen.model;

/**
 * @version V1.0
 */
public final class Field {
  private String name;
  private String type;
  private boolean notNull;
  private String defValue;
  private SignType sign;

  public Field() {
    this.notNull = false;
    this.sign = SignType.SIGNED;
  }

  public String getName() {
    return name;
  }

  public void setName(String name) {
    this.name = name;
  }

  public String getType() {
    return type;
  }

  public void setType(String type) {
    this.type = type;
  }

  public boolean isNotNull() {
    return notNull;
  }

  public void setNotNull(boolean notNull) {
    this.notNull = notNull;
  }

  public String getDefValue() {
    return defValue;
  }

  public void setDefValue(String defValue) {
    this.defValue = defValue;
  }

  public SignType getSign() {
    return sign;
  }

  public void setSign(SignType sign) {
    this.sign = sign;
  }

  /**
   * Field Sign Type Enum.
   */
  public enum SignType {
    SIGNED, UNSIGNED, ALL
  }
}
