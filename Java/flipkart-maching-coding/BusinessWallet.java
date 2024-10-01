package com.flipkart.wallet.model;


public class BusinessWallet extends Wallet {
  private String gst;

  public BusinessWallet(String accountHolder, double balance, String gst) {
      super("Business", accountHolder, balance);
      this.gst = gst;
  }

  public String getGst() {
      return gst;
  }
}
