package com.flipkart.wallet;

import com.flipkart.wallet.service.WalletService;

//@SpringBootApplication
public class WalletApplication {

  public static void main(String[] args) {
    // SpringApplication.run(WalletApplication.class, args);
    WalletService walletService = new WalletService();
    walletService.createWallet("Personal", "Harry", null, 100);
    walletService.createWallet("Personal", "Ron", null, 95);
    walletService.createWallet("Personal", "Hermione", null, 104);
    walletService.createWallet("Business", "Albus", "GST1234", 200);
    walletService.createWallet("Business", "Draco", "GST5678", 500);
    walletService.overview();
 //   walletService.transferMoney("Albus", "Albus", 30);
    walletService.transferMoney("Hermione", "Harry", 2);
    walletService.transferMoney("Albus", "Ron", 5);
    walletService.overview();
    walletService.balance("Harry");
  }

}
