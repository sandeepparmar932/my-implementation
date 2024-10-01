package com.flipkart.wallet.exceptions;

public class InsufficientBalanceException extends RuntimeException{
  
  private String msg;
  
  public InsufficientBalanceException(String msg) {
    this.msg = msg;
  }
}
