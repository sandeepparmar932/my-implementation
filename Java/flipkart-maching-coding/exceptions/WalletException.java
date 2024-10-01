package com.flipkart.wallet.exceptions;

public class WalletException extends RuntimeException{
  private String msg;
  public WalletException(String string) {
  this.msg =msg;
  }

}
