package com.flipkart.wallet.exceptions;

public class InvalidUserException extends RuntimeException{
  private String msg;
  public InvalidUserException(String msg) {
    this.msg = msg;
  }
}
