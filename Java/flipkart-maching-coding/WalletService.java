package com.flipkart.wallet.service;

import java.util.ArrayList;
import java.util.List;

import com.flipkart.wallet.exceptions.InsufficientBalanceException;
import com.flipkart.wallet.exceptions.InvalidUserException;
import com.flipkart.wallet.exceptions.WalletException;
import com.flipkart.wallet.model.BusinessWallet;
import com.flipkart.wallet.model.Wallet;

public class WalletService {
  private List<Wallet> wallets;
   private double transactionFee;

  public WalletService() {
    this.wallets = new ArrayList<>();
    this.transactionFee = 5.0;
  }

  public void createWallet(String accountType, String accountHolder, String gst, double amount) {
    if(amount < 0) {
      throw new WalletException("amount can not be negative");
    } 
    if("Personal".equals(accountType)) {
       if(  amount < 50)
         throw new WalletException("Can not create personal wallet with less than 50 amount");
       Wallet wallet = new Wallet(accountType, accountHolder, amount);
       wallets.add(wallet);
       System.out.println("Personal wallet created for " + accountHolder + " with a balance of F₹ " + amount);
    }else if ("Business".equals(accountType)) {
      BusinessWallet businessWallet = new BusinessWallet(accountHolder, amount, gst);
      wallets.add(businessWallet);
      System.out.println("Business wallet created for " + accountHolder + " with a balance of F₹ " + amount);
    }else  {
      System.out.println("Invalid account type");
    }
  }

  public void addMoney(String accountHolder, double amount) {
    for (Wallet wallet : wallets) {
      if (wallet.getAccountHolder().equals(accountHolder)) {
        wallet.setBalance(wallet.getBalance() + amount);
        System.out.println("F₹ " + amount + " added to " + accountHolder + "'s wallet");
        return;
      }
    }
    throw new WalletException("Wallet not found");
  }

  public void transferMoney(String sender, String receiver, double amount) {
    Wallet senderWallet = null;
    Wallet receiverWallet = null;

    for (Wallet wallet : wallets) {
      if (wallet.getAccountHolder().equals(sender)) {
        senderWallet = wallet;
      } else if (wallet.getAccountHolder().equals(receiver)) {
        receiverWallet = wallet;
      }
    }

    if (senderWallet == null || receiverWallet == null) {
      throw new InvalidUserException("Invalid sender or receiver");
    }

    if (senderWallet.getAccountType().equals("Business") && receiverWallet.getAccountType().equals("Personal")) {
      System.out.println("Business account cannot transfer money to a personal account");
      return;
    }
    if("Personal".equals(senderWallet.getAccountType()) &&"Personal".equals(receiverWallet.getAccountType()) && senderWallet.getBalance() < (amount + transactionFee) ) {
      throw new InsufficientBalanceException("Insufficient balance in the sender's account");
    } 
    if (senderWallet.getBalance() < amount) {
      throw new InsufficientBalanceException("Insufficient balance in the sender's account");
    }
    senderWallet.setBalance(senderWallet.getBalance() - amount);
    receiverWallet.setBalance(receiverWallet.getBalance() + amount);

    System.out.println("F₹ " + amount + " transferred from " + sender + "'s account to " + receiver + "'s account");

    if (senderWallet.getBalance() == receiverWallet.getBalance()) {
      senderWallet.setBalance(senderWallet.getBalance() + 10.0);
      receiverWallet.setBalance(receiverWallet.getBalance() + 10.0);
      System.out.println("Reward of F₹ 10 given to both as the  " + sender + " and " + receiver);
    }
  }

  public void overview() {
    for (Wallet wallet : wallets) {
      System.out.println(wallet.getAccountHolder() + "\t " + wallet.getBalance());
    }
  }

  public void balance(String accountHolder) {
    for (Wallet wallet : wallets) {
      if (wallet.getAccountHolder().equals(accountHolder)) {
        System.out.println("Balance for " + accountHolder + "\t" + wallet.getBalance());
        return;
      }
    }
    throw new WalletException("Wallet not found");
  }
}
