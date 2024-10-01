package com.flipkart.wallet.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@AllArgsConstructor
@NoArgsConstructor
public class Wallet {
  private String accountType;
  private String accountHolder;
  private String gst;
  private double balance;
  
  public Wallet(String accountType, String accountHolder, double balance) {
    this.accountType = accountType;
    this.accountHolder = accountHolder;
    this.balance = balance;
}
  
}
